package dispatcher

import (
	"context"
	"time"

	"github.com/llm-operator/job-manager/common/pkg/store"
	"github.com/llm-operator/job-manager/dispatcher/internal/config"
	"github.com/llm-operator/job-manager/dispatcher/pkg/util"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	batchv1apply "k8s.io/client-go/applyconfigurations/batch/v1"
	corev1apply "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/utils/ptr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	managedJobAnnotationKey = "llm-operator/managed-pod"
	jobIDAnnotationKey      = "llm-operator/job-id"

	jobTTL = time.Hour * 24
)

// NewJobClient returns a new JobCreator.
func NewJobClient(
	k8sClient client.Client,
	namespace string,
	modelStoreConfig *config.ModelStoreConfig,
	useFakeJob bool,
	huggingFaceAccessToken string,
) *JobClient {
	return &JobClient{
		k8sClient:              k8sClient,
		namespace:              namespace,
		modelStoreConfig:       modelStoreConfig,
		useFakeJob:             useFakeJob,
		huggingFaceAccessToken: huggingFaceAccessToken,
	}
}

// JobClient operates a Kubernetes Job resource for a job.
type JobClient struct {
	k8sClient client.Client
	// TODO(kenji): Be able to specify the namespace per tenant.
	namespace              string
	modelStoreConfig       *config.ModelStoreConfig
	useFakeJob             bool
	huggingFaceAccessToken string
}

func (p *JobClient) createJob(ctx context.Context, jobData *store.Job) error {
	// TODO(kenji): Create a real fine-tuning job. See https://github.com/llm-operator/job-manager/tree/main/build/experiments/fine-tuning.
	log := ctrl.LoggerFrom(ctx)

	log.Info("Creating a k8s Job resource for a job")

	// TODO(kenji): Manage training files. Download them from the object store if needed.

	obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(batchv1apply.
		Job(util.GetK8sJobName(jobData.JobID), p.namespace).
		WithAnnotations(map[string]string{
			managedJobAnnotationKey: "true",
			jobIDAnnotationKey:      jobData.JobID}).
		WithSpec(p.jobSpec()))
	if err != nil {
		return err
	}
	patch := &unstructured.Unstructured{Object: obj}
	opts := &client.PatchOptions{FieldManager: "job-manager-dispatcher", Force: ptr.To(true)}
	return p.k8sClient.Patch(ctx, patch, client.Apply, opts)
}

func (p *JobClient) jobSpec() *batchv1apply.JobSpecApplyConfiguration {
	var image, cmd string
	var res *corev1apply.ResourceRequirementsApplyConfiguration
	if p.useFakeJob {
		image = "llm-operator/experiments-fake-job:latest"
		cmd = "mkdir /models/adapter; cp ./ggml-adapter-model.bin /models/adapter/ggml-adapter-model.bin"
	} else {
		image = "llm-operator/experiments-fine-tuning:latest"
		cmd = `
mkdir /models/adapter;
accelerate launch \
  --config_file=./single_gpu.yaml \
  --num_processes=1 \
  ./sft.py \
  --model_name=google/gemma-2b \
  --dataset_name=OpenAssistant/oasst_top1_2023-08-25 \
  --per_device_train_batch_size=2 \
  --gradient_accumulation_steps=1 \
  --max_steps=100 \
  --learning_rate=2e-4 \
  --save_steps=20_000 \
  --use_peft \
  --lora_r=16 \
  --lora_alpha=32 \
  --lora_target_modules q_proj k_proj v_proj o_proj \
  --load_in_4bit \
  --output_dir=./output &&
python ./convert-lora-to-ggml.py ./output &&
cp ./output/ggml-adapter-model.bin /models/adapter/
`
		res = corev1apply.ResourceRequirements().
			WithLimits(corev1.ResourceList{
				"nvidia.com/gpu": *resource.NewQuantity(1, resource.DecimalSI),
			})
	}

	container := corev1apply.Container().
		WithName("main").
		WithImage(image).
		WithImagePullPolicy(corev1.PullNever).
		WithCommand("/bin/bash", "-c", cmd).
		WithResources(res).
		WithEnv(corev1apply.EnvVar().
			WithName("HUGGING_FACE_HUB_TOKEN").
			WithValue(p.huggingFaceAccessToken))
	podSpec := corev1apply.PodSpec().
		WithContainers(container).
		WithRestartPolicy(corev1.RestartPolicyNever)
	jobSpec := batchv1apply.JobSpec().
		WithTTLSecondsAfterFinished(int32(jobTTL.Seconds())).
		WithBackoffLimit(3).
		WithTemplate(corev1apply.PodTemplateSpec().
			WithSpec(podSpec))

	if ms := p.modelStoreConfig; ms.Enable {
		const vname = "model-store"
		container.WithVolumeMounts(corev1apply.VolumeMount().
			WithName(vname).
			WithMountPath("/models"))
		podSpec.WithVolumes(corev1apply.Volume().
			WithName(vname).
			WithPersistentVolumeClaim(corev1apply.PersistentVolumeClaimVolumeSource().
				WithClaimName(ms.PVClaimName)))
	}
	return jobSpec
}

func isManagedJob(annotations map[string]string) bool {
	return annotations[managedJobAnnotationKey] == "true"
}