package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
	"os"
)

const (
	EnvironmentVariablePrefix            = "TANUKI_SERVER_"
	ConfigurationFileEnvironmentVariable = EnvironmentVariablePrefix + "CONFIGURATION_FILE"
)

const (
	DefaultServerPort  = 8080
	DefaultServiceName = "tanuki"
	DefaultEnvironment = "development"
)

type Flags interface {
	Lookup(string) *pflag.Flag
}

func NewConfigurationFromFlags(flags Flags) (Config, error) {
	viperInstance := viper.New()
	viperInstance.SetEnvPrefix(EnvironmentVariablePrefix)
	viperInstance.AutomaticEnv()

	_ = viperInstance.BindPFlag("target", flags.Lookup("target"))
	_ = viperInstance.BindPFlag("database.username", flags.Lookup("database.username"))
	_ = viperInstance.BindPFlag("database.password", flags.Lookup("database.password"))
	_ = viperInstance.BindPFlag("database.name", flags.Lookup("database.name"))
	_ = viperInstance.BindPFlag("database.host", flags.Lookup("database.host"))
	_ = viperInstance.BindPFlag("database.port", flags.Lookup("database.port"))
	_ = viperInstance.BindPFlag("database.type", flags.Lookup("database.type"))
	_ = viperInstance.BindPFlag("database.options", flags.Lookup("database.options"))
	_ = viperInstance.BindPFlag("servers.registry.port", flags.Lookup("servers.registry.port"))
	_ = viperInstance.BindPFlag("servers.management.port", flags.Lookup("servers.management.port"))
	_ = viperInstance.BindPFlag("log-level", flags.Lookup("log-level"))

	configFilePath := os.Getenv(ConfigurationFileEnvironmentVariable)
	isConfigFileUserDefined := configFilePath != ""
	if !isConfigFileUserDefined {
		configFilePath = "local.yaml"
	}
	_, err := os.Stat(configFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return Config{}, err
		}
		if isConfigFileUserDefined {
			return Config{}, err
		}
	} else {
		viperInstance.SetConfigFile(configFilePath)
		if err = viperInstance.ReadInConfig(); err != nil {
			return Config{}, err
		}
	}
	var instance Config
	if err = viperInstance.Unmarshal(&instance); err != nil {
		return Config{}, err
	}
	if instance.LogLevel == "" {
		instance.LogLevel = InfoLogLevel
	}
	if !funk.Contains([]string{DebugLogLevel, InfoLogLevel}, instance.LogLevel) {
		return instance, fmt.Errorf("log-level must be either %s or %s", InfoLogLevel, DebugLogLevel)
	}
	if instance.Environment == "" {
		instance.Environment = DefaultEnvironment
	}
	if instance.Service == "" {
		instance.Service = DefaultServiceName
	}
	if instance.Targets == nil {
		instance.Targets = getDefaultTargets()
	}
	if instance.Servers.Registry.Port == nil {
		instance.Servers.Registry.Port = funk.PtrOf(DefaultServerPort).(*int)
	}

	assertTLSConfigurationOnServer := func(target Server, name string) error {
		if target.TLS == nil {
			return nil
		}
		if target.TLS.CertFile == "" {
			return fmt.Errorf("servers.%s.tls.cert must be set", name)
		}
		if _, prob := os.Stat(target.TLS.CertFile); prob != nil {
			return fmt.Errorf("servers.%s.tls.cert=%s couldn't be accessed due to: %v", name, target.TLS.CertFile, prob)
		}
		if target.TLS.KeyFile == "" {
			return fmt.Errorf("servers.%s.tls.key must be set", name)
		}
		if _, prob := os.Stat(target.TLS.KeyFile); prob != nil {
			return fmt.Errorf("servers.%s.tls.key=%s couldn't be accessed due to: %v", name, target.TLS.KeyFile, prob)
		}
		return nil
	}

	if err = assertTLSConfigurationOnServer(instance.Servers.Registry, "registry"); err != nil {
		return Config{}, err
	}

	if err = assertTLSConfigurationOnServer(instance.Servers.Management, "management"); err != nil {
		return Config{}, err
	}

	return instance, nil
}

func getDefaultTargets() []string {
	return []string{
		"https://registry.terraform.io/",
	}
}
