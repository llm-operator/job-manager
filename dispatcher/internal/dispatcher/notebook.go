package dispatcher

import (
	"context"
	"fmt"

	v1 "github.com/llm-operator/job-manager/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	appsv1apply "k8s.io/client-go/applyconfigurations/apps/v1"
	corev1apply "k8s.io/client-go/applyconfigurations/core/v1"
	metav1apply "k8s.io/client-go/applyconfigurations/meta/v1"
	netv1apply "k8s.io/client-go/applyconfigurations/networking/v1"
	"k8s.io/utils/ptr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	managedAnnotationKey    = "llm-operator/managed"
	notebookIDAnnotationKey = "llm-operator/notebook-id"

	nbManagerName = "notebook-manager"
)

// NewNotebookManager creates a new NotebookManager
func NewNotebookManager(
	k8sClient client.Client,
	llmoBaseURL string,
	ingressClassName string,
	clusterID string,
) *NotebookManager {
	return &NotebookManager{
		k8sClient:        k8sClient,
		llmoBaseURL:      llmoBaseURL,
		ingressClassName: ingressClassName,
		clusterID:        clusterID,
	}
}

// NotebookManager is a struct that manages the notebook
type NotebookManager struct {
	k8sClient client.Client

	llmoBaseURL      string
	ingressClassName string

	clusterID string
}

func (n *NotebookManager) createNotebook(ctx context.Context, nb *v1.InternalNotebook) error {
	log := ctrl.LoggerFrom(ctx)
	log.Info("Creating a deployment for a notebook")

	name := nb.Notebook.Id
	labels := map[string]string{
		"app.kubernetes.io/name":       "llmo-notebook",
		"app.kubernetes.io/instance":   name,
		"app.kubernetes.io/created-by": nbManagerName,
	}

	var envs []*corev1apply.EnvVarApplyConfiguration
	for k, v := range nb.Notebook.Envs {
		envs = append(envs, corev1apply.EnvVar().WithName(k).WithValue(v))
	}
	envs = append(envs, corev1apply.EnvVar().WithName("OPENAI_BASE_URL").WithValue(n.llmoBaseURL))

	req := corev1.ResourceList{}
	limit := corev1.ResourceList{}
	if r := nb.Notebook.Resources; r != nil {
		if cpu := r.CpuMilicore; cpu != nil {
			if cpu.Requests > 0 {
				req[corev1.ResourceCPU] = *resource.NewMilliQuantity(int64(r.CpuMilicore.Requests), resource.DecimalSI)
			}
			if cpu.Limits > 0 {
				limit[corev1.ResourceCPU] = *resource.NewMilliQuantity(int64(r.CpuMilicore.Limits), resource.DecimalSI)
			}
		}
		if mem := r.MemoryMegabytes; mem != nil {
			if mem.Requests > 0 {
				req[corev1.ResourceMemory] = *resource.NewScaledQuantity(int64(mem.Requests), 6)
			}
			if mem.Limits > 0 {
				limit[corev1.ResourceMemory] = *resource.NewScaledQuantity(int64(mem.Limits), 6)
			}
		}
		if r.GpuCount > 0 {
			limit["nvidia.com/gpu"] = *resource.NewQuantity(int64(r.GpuCount), resource.DecimalSI)
		}
	}
	resources := corev1apply.ResourceRequirements()
	if len(req) > 0 {
		resources.WithRequests(req)
	}
	if len(limit) > 0 {
		resources.WithLimits(limit)
	}

	// TODO: set volume mounts and volumes for the notebook
	const (
		appPort  = 8888
		portName = "jupyter-web-ui"
	)
	var baseURL = fmt.Sprintf("/v1/sessions/%s/v1/services/notebooks/%s", n.clusterID, nb.Notebook.Id)

	deployConf := appsv1apply.
		Deployment(name, nb.Notebook.KubernetesNamespace).
		WithLabels(labels).
		WithAnnotations(map[string]string{
			managedAnnotationKey:    "true",
			notebookIDAnnotationKey: nb.Notebook.Id}).
		WithSpec(appsv1apply.DeploymentSpec().
			WithReplicas(1).
			WithSelector(metav1apply.LabelSelector().
				WithMatchLabels(labels)).
			WithTemplate(corev1apply.PodTemplateSpec().
				WithLabels(labels).
				WithSpec(corev1apply.PodSpec().
					WithContainers(corev1apply.Container().
						WithName("jupyterlab").
						WithImage(nb.Notebook.Image).
						WithImagePullPolicy(corev1.PullIfNotPresent).
						// TODO: rethink authentication method
						WithCommand("start-notebook.py").
						WithArgs(
							"--IdentityProvider.token=$(NOTEBOOK_TOKEN)",
							"--ServerApp.base_url="+baseURL,
							// This is needed when a user accesses the notebook
							// via Session Manager/Agent and internal ingress controller.
							// TODO(kenji): Tighten this.
							"--NotebookApp.allow_origin=*").
						WithPorts(corev1apply.ContainerPort().
							WithName(portName).
							WithContainerPort(appPort).
							WithProtocol(corev1.ProtocolTCP)).
						WithEnv(envs...).
						WithEnvFrom(corev1apply.EnvFromSource().
							WithSecretRef(corev1apply.SecretEnvSource().
								WithName(nb.Notebook.Id))).
						WithResources(resources)))))

	svcConf := corev1apply.Service(name, nb.Notebook.KubernetesNamespace).
		WithLabels(labels).
		WithAnnotations(map[string]string{
			managedAnnotationKey:    "true",
			notebookIDAnnotationKey: nb.Notebook.Id}).
		WithSpec(corev1apply.ServiceSpec().
			WithType(corev1.ServiceTypeClusterIP).
			WithSelector(labels).
			WithPorts(corev1apply.ServicePort().
				WithName(portName).
				WithPort(appPort).
				WithTargetPort(intstr.FromString(portName)).
				WithProtocol(corev1.ProtocolTCP)))

	ingConf := netv1apply.Ingress(name, nb.Notebook.KubernetesNamespace).
		WithLabels(labels).
		WithAnnotations(map[string]string{
			managedAnnotationKey:    "true",
			notebookIDAnnotationKey: nb.Notebook.Id}).
		WithSpec(netv1apply.IngressSpec().
			WithIngressClassName(n.ingressClassName).
			WithRules(netv1apply.IngressRule().
				WithHTTP(netv1apply.HTTPIngressRuleValue().
					WithPaths(netv1apply.HTTPIngressPath().
						WithPath(baseURL).
						WithPathType(netv1.PathTypePrefix).
						WithBackend(netv1apply.IngressBackend().
							WithService(netv1apply.IngressServiceBackend().
								WithName(name).
								WithPort(netv1apply.ServiceBackendPort().
									WithName(portName))))))))

	patchOpts := &client.PatchOptions{FieldManager: nbManagerName, Force: ptr.To(true)}
	deploy, err := n.applyObject(ctx, deployConf, patchOpts)
	if err != nil {
		return err
	}

	gvk := deploy.GetObjectKind().GroupVersionKind()
	ownerRef := metav1apply.OwnerReference().
		WithAPIVersion(gvk.GroupVersion().String()).
		WithKind(gvk.Kind).
		WithName(deploy.GetName()).
		WithUID(deploy.GetUID()).
		WithBlockOwnerDeletion(true).
		WithController(true)

	svcConf.WithOwnerReferences(ownerRef)
	ingConf.WithOwnerReferences(ownerRef)

	// Secret is pre-created by server, and dispatcher only set the owner reference here.
	// TODO(aya): garbage collect orphaned secrets
	secConf := corev1apply.Secret(nb.Notebook.Id, nb.Notebook.KubernetesNamespace).
		WithOwnerReferences(ownerRef)

	for _, obj := range []any{svcConf, ingConf, secConf} {
		if _, err := n.applyObject(ctx, obj, patchOpts); err != nil {
			return err
		}
	}
	return nil
}

func (n *NotebookManager) applyObject(ctx context.Context, applyConfig any, opts ...client.PatchOption) (client.Object, error) {
	uobj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(applyConfig)
	if err != nil {
		return nil, err
	}
	obj := &unstructured.Unstructured{Object: uobj}
	if err := n.k8sClient.Patch(ctx, obj, client.Apply, opts...); err != nil {
		return nil, err
	}
	return obj, nil
}

func (n *NotebookManager) stopNotebook(ctx context.Context, nb *v1.InternalNotebook) error {
	log := ctrl.LoggerFrom(ctx)
	log.Info("Stopping a deployment for a notebook")

	var deploy appsv1.Deployment
	if err := n.k8sClient.Get(ctx, types.NamespacedName{
		Name:      nb.Notebook.Id,
		Namespace: nb.Notebook.KubernetesNamespace,
	}, &deploy); err != nil {
		return err
	}

	scale := &autoscalingv1.Scale{Spec: autoscalingv1.ScaleSpec{Replicas: 0}}
	return n.k8sClient.SubResource("scale").Update(ctx, &deploy, client.WithSubResourceBody(scale))
}

func (n *NotebookManager) deleteNotebook(ctx context.Context, nb *v1.InternalNotebook) error {
	log := ctrl.LoggerFrom(ctx)
	log.Info("Deleting a deployment for a notebook")

	var deploy appsv1.Deployment
	if err := n.k8sClient.Get(ctx, types.NamespacedName{
		Name:      nb.Notebook.Id,
		Namespace: nb.Notebook.KubernetesNamespace,
	}, &deploy); err != nil {
		if apierrors.IsNotFound(err) {
			log.V(4).Info("Deployment not found")
			return nil
		}
		return err
	}

	if err := n.k8sClient.Delete(ctx, &deploy); err == nil {
		if apierrors.IsNotFound(err) {
			log.V(4).Info("Deployment not found")
			return nil
		}
		return err
	}
	return nil
}
