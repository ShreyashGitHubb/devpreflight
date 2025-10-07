package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Checks ChecksConfig `mapstructure:"checks"`
	Env    EnvConfig    `mapstructure:"env"`
	Docker DockerConfig `mapstructure:"docker"`
	K8s    K8sConfig    `mapstructure:"k8s"`
	CI     CIConfig     `mapstructure:"ci"`
	Report ReportConfig `mapstructure:"report"`
}

type ChecksConfig struct {
	EnvParity     bool `mapstructure:"env_parity"`
	DockerfileLint bool `mapstructure:"dockerfile_lint"`
	K8sValidate   bool `mapstructure:"k8s_validate"`
	Observability bool `mapstructure:"observability"`
	FlakyTests    bool `mapstructure:"flaky_tests"`
}

type EnvConfig struct {
	RequiredKeys []string `mapstructure:"required_keys"`
}

type DockerConfig struct {
	ForbidLatest bool `mapstructure:"forbid_latest"`
}

type K8sConfig struct {
	SchemaValidate bool `mapstructure:"schema_validate"`
}

type CIConfig struct {
	TimeoutSeconds int `mapstructure:"timeout_seconds"`
}

type ReportConfig struct {
	Format string `mapstructure:"format"`
	Output string `mapstructure:"output"`
}

func LoadConfig() *Config {
	setDefaults()
	
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		// Return defaults if unmarshal fails
		return getDefaultConfig()
	}
	
	return &cfg
}

func setDefaults() {
	viper.SetDefault("checks.env_parity", true)
	viper.SetDefault("checks.dockerfile_lint", true)
	viper.SetDefault("checks.k8s_validate", true)
	viper.SetDefault("checks.observability", true)
	viper.SetDefault("checks.flaky_tests", false)
	viper.SetDefault("docker.forbid_latest", true)
	viper.SetDefault("k8s.schema_validate", true)
	viper.SetDefault("ci.timeout_seconds", 300)
	viper.SetDefault("report.format", "console")
}

func getDefaultConfig() *Config {
	return &Config{
		Checks: ChecksConfig{
			EnvParity:     true,
			DockerfileLint: true,
			K8sValidate:   true,
			Observability: true,
			FlakyTests:    false,
		},
		Docker: DockerConfig{
			ForbidLatest: true,
		},
		K8s: K8sConfig{
			SchemaValidate: true,
		},
		CI: CIConfig{
			TimeoutSeconds: 300,
		},
		Report: ReportConfig{
			Format: "console",
		},
	}
}
