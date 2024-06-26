package main

import (
	"context"
	"crypto/tls"
	"log"
	"log/slog"

	"github.com/go-logr/logr"
	cv1 "github.com/llm-operator/cluster-manager/api/v1"
	fv1 "github.com/llm-operator/file-manager/api/v1"
	v1 "github.com/llm-operator/job-manager/api/v1"
	"github.com/llm-operator/job-manager/dispatcher/internal/config"
	"github.com/llm-operator/job-manager/dispatcher/internal/dispatcher"
	"github.com/llm-operator/job-manager/dispatcher/internal/s3"
	mv1 "github.com/llm-operator/model-manager/api/v1"
	"github.com/llm-operator/rbac-manager/pkg/auth"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	ctrl "sigs.k8s.io/controller-runtime"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

const flagConfig = "config"

var setupLog = ctrl.Log.WithName("setup")

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := cmd.Flags().GetString(flagConfig)
		if err != nil {
			return err
		}

		c, err := config.Parse(path)
		if err != nil {
			return err
		}

		if err := c.Validate(); err != nil {
			return err
		}

		if err := run(cmd.Context(), &c); err != nil {
			return err
		}
		return nil
	},
}

func run(ctx context.Context, c *config.Config) error {
	ctrl.SetLogger(logr.FromSlogHandler(slog.Default().Handler()))

	restConfig, err := newRestConfig(c.Debug.KubeconfigPath)
	if err != nil {
		return err
	}
	mgr, err := ctrl.NewManager(restConfig, ctrl.Options{
		LeaderElection:   c.KubernetesManager.EnableLeaderElection,
		LeaderElectionID: c.KubernetesManager.LeaderElectionID,
		Metrics: metricsserver.Options{
			BindAddress: c.KubernetesManager.MetricsBindAddress,
		},
		HealthProbeBindAddress: c.KubernetesManager.HealthBindAddress,
		PprofBindAddress:       c.KubernetesManager.PprofBindAddress,
	})
	if err != nil {
		return err
	}

	jc := dispatcher.NewJobClient(
		mgr.GetClient(),
		c.Job,
		c.KueueIntegration,
	)

	preProcessor, postProcessor, err := newProcessors(c)
	if err != nil {
		return err
	}

	clusterID, err := getClusterID(ctx, c)
	if err != nil {
		return err
	}

	nb := dispatcher.NewNotebookManager(mgr.GetClient(), c.Notebook.LLMOperatorBaseURL, c.Notebook.IngressClassName, clusterID)

	conn, err := grpc.NewClient(c.JobManagerServerWorkerServiceAddr, grpcOption(c))
	if err != nil {
		return err
	}
	ftClient := v1.NewFineTuningWorkerServiceClient(conn)
	wsClient := v1.NewWorkspaceWorkerServiceClient(conn)

	if err := dispatcher.New(ftClient, wsClient, jc, preProcessor, nb, c.PollingInterval).
		SetupWithManager(mgr); err != nil {
		return err
	}

	if err := dispatcher.NewLifecycleManager(ftClient, mgr.GetClient(), postProcessor).
		SetupWithManager(mgr); err != nil {
		return err
	}
	return mgr.Start(ctx)
}

func getClusterID(ctx context.Context, c *config.Config) (string, error) {
	conn, err := grpc.NewClient(c.ClusterManagerServerWorkerServiceAddr, grpcOption(c))
	if err != nil {
		return "", err
	}
	clClient := cv1.NewClustersWorkerServiceClient(conn)

	ctx = auth.AppendWorkerAuthorization(ctx)
	cluster, err := clClient.GetSelfCluster(ctx, &cv1.GetSelfClusterRequest{})
	if err != nil {
		return "", err
	}
	log.Printf("Obtained the cluster ID: %q\n", cluster.Id)
	return cluster.Id, nil
}

func newProcessors(c *config.Config) (dispatcher.PreProcessorI, dispatcher.PostProcessorI, error) {
	if c.Debug.Standalone {
		return &dispatcher.NoopPreProcessor{}, &dispatcher.NoopPostProcessor{}, nil
	}

	option := grpcOption(c)
	conn, err := grpc.NewClient(c.FileManagerServerWorkerServiceAddr, option)
	if err != nil {
		return nil, nil, err
	}
	fclient := fv1.NewFilesWorkerServiceClient(conn)

	conn, err = grpc.NewClient(c.ModelManagerServerWorkerServiceAddr, option)
	if err != nil {
		return nil, nil, err
	}
	mclient := mv1.NewModelsWorkerServiceClient(conn)
	s3Client := s3.NewClient(c.ObjectStore.S3)

	preProcessor := dispatcher.NewPreProcessor(fclient, mclient, s3Client)
	postProcessor := dispatcher.NewPostProcessor(mclient)
	return preProcessor, postProcessor, nil
}

func newRestConfig(kubeconfigPath string) (*rest.Config, error) {
	if kubeconfigPath != "" {
		setupLog.Info("Using kubeconfig at", "path", kubeconfigPath)
		return clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	}
	return rest.InClusterConfig()
}

func init() {
	runCmd.Flags().StringP(flagConfig, "c", "", "Configuration file path")
	_ = runCmd.MarkFlagRequired(flagConfig)
}

func grpcOption(c *config.Config) grpc.DialOption {
	if c.Worker.TLS.Enable {
		return grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{}))
	}
	return grpc.WithTransportCredentials(insecure.NewCredentials())
}
