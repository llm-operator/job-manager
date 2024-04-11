package config

import (
	"fmt"
	"os"
	"time"

	"github.com/llm-operator/job-manager/common/pkg/db"
	"gopkg.in/yaml.v3"
)

// S3Config is the S3 configuration.
type S3Config struct {
	EndpointURL string `yaml:"endpointUrl"`
	Bucket      string `yaml:"bucket"`
}

// ObjectStoreConfig is the object store configuration.
type ObjectStoreConfig struct {
	S3 S3Config `yaml:"s3"`
}

// Validate validates the object store configuration.
func (c *ObjectStoreConfig) Validate() error {
	if c.S3.EndpointURL == "" {
		return fmt.Errorf("s3 endpoint url must be set")
	}
	if c.S3.Bucket == "" {
		return fmt.Errorf("s3 bucket must be set")
	}
	return nil
}

// DebugConfig is the debug configuration.
type DebugConfig struct {
	KubeconfigPath string `yaml:"kubeconfigPath"`
	Standalone     bool   `yaml:"standalone"`
	SqlitePath     string `yaml:"sqlitePath"`
	UseFakeJob     bool   `yaml:"useFakeJob"`

	HuggingFaceAccessToken string `yaml:"huggingFaceAccessToken"`
}

// ModelStoreConfig is the model store configuration.
type ModelStoreConfig struct {
	Enable      bool   `yaml:"enable"`
	MountPath   string `yaml:"mountPath"`
	PVClaimName string `yaml:"pvClaimName"`
}

// KubernetesManagerConfig is the Kubernetes manager configuration.
type KubernetesManagerConfig struct {
	EnableLeaderElection bool   `yaml:"enableLeaderElection"`
	LeaderElectionID     string `yaml:"leaderElectionID"`

	MetricsBindAddress string `yaml:"metricsBindAddress"`
	HealthBindAddress  string `yaml:"healthBindAddress"`
	PprofBindAddress   string `yaml:"pprofBindAddress"`
}

// Config is the configuration.
type Config struct {
	JobPollingInterval time.Duration `yaml:"jobPollingInterval"`
	JobNamespace       string        `yaml:"jobNamespace"`

	ModelManagerServerAddr string `yaml:"modelManagerServerAddr"`

	// TODO(kenji): Remove this. This was created to share models between job-manager-dispatcher
	// and inference-manager-engine, but models should be stored in an object store instead.
	ModelStore ModelStoreConfig `yaml:"modelStore"`

	Database db.Config `yaml:"database"`

	ObjectStore ObjectStoreConfig `yaml:"objectStore"`

	Debug DebugConfig `yaml:"debug"`

	KubernetesManager KubernetesManagerConfig `yaml:"kubernetesManager"`
}

// Validate validates the configuration.
func (c *Config) Validate() error {
	if c.JobPollingInterval <= 0 {
		return fmt.Errorf("job polling interval must be greater than 0")
	}
	if c.JobNamespace == "" {
		return fmt.Errorf("job namespace must be set")
	}
	if c.Debug.Standalone {
		if c.Debug.SqlitePath == "" {
			return fmt.Errorf("sqlite path must be set")
		}
	} else {
		if c.ModelManagerServerAddr == "" {
			return fmt.Errorf("model manager address must be set")
		}

		if err := c.ObjectStore.Validate(); err != nil {
			return fmt.Errorf("object store: %s", err)
		}

		if err := c.Database.Validate(); err != nil {
			return fmt.Errorf("database: %s", err)
		}
	}
	return nil
}

// Parse parses the configuration file at the given path, returning a new
// Config struct.
func Parse(path string) (Config, error) {
	var config Config

	b, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("config: read: %s", err)
	}

	if err = yaml.Unmarshal(b, &config); err != nil {
		return config, fmt.Errorf("config: unmarshal: %s", err)
	}
	return config, nil
}
