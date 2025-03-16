package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	EnvironmentVariablePrefix            = "TANUKI"
	ConfigurationFileEnvironmentVariable = EnvironmentVariablePrefix + "_CONFIGURATION_FILE"
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
	viperInstance := createViperInstance()
	bindFlagsToViper(viperInstance, flags)
	err := readViperConfigurationFile(viperInstance)
	if err != nil {
		return Config{}, err
	}
	fmt.Println(
		"-------------------",
	)

	for k, v := range viperInstance.AllKeys() {
		fmt.Println(k, v)
	}
	fmt.Println(
		"-------------------",
	)
	var instance Config
	if err = viperInstance.Unmarshal(&instance); err != nil {
		return Config{}, err
	}

	setConfigurationDefaults(&instance)

	if err = validateServersTLSConfiguration(instance.Servers.Registry, "registry"); err != nil {
		return Config{}, err
	}
	if err = validateServersTLSConfiguration(instance.Servers.Management, "management"); err != nil {
		return Config{}, err
	}

	return instance, nil
}

func createViperInstance() *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(EnvironmentVariablePrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	return v
}

func bindFlagsToViper(v *viper.Viper, flags Flags) {
	flagMappings := []string{
		"target",
		"database.username", "database.password", "database.name", "database.host", "database.port", "database.type", "database.options",
		"servers.registry.port", "servers.registry.tls.key", "servers.registry.tls.cert",
		"servers.management.port", "servers.management.tls.key", "servers.management.tls.cert",
		"log-level",
	}

	for _, key := range flagMappings {
		_ = v.BindPFlag(key, flags.Lookup(key))
	}
}

func readViperConfigurationFile(v *viper.Viper) error {
	configFilePath := os.Getenv(ConfigurationFileEnvironmentVariable)
	isUserDefined := configFilePath != ""
	if configFilePath == "" {
		configFilePath = "local.yaml"
	}
	if _, err := os.Stat(configFilePath); err != nil {
		if os.IsNotExist(err) && !isUserDefined {
			return nil // No config file is fine unless explicitly specified
		}
		return err
	}
	v.SetConfigFile(configFilePath)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func setConfigurationDefaults(config *Config) {
	if config.LogLevel == "" {
		config.LogLevel = InfoLogLevel
	}
	if config.Environment == "" {
		config.Environment = DefaultEnvironment
	}
	if config.Service == "" {
		config.Service = DefaultServiceName
	}
	if config.Targets == nil {
		config.Targets = getDefaultTargets()
	}
	if config.Servers.Registry.Port == nil {
		defaultPort := DefaultServerPort
		config.Servers.Registry.Port = &defaultPort
	}

	validLogLevels := map[string]bool{DebugLogLevel: true, InfoLogLevel: true}
	if !validLogLevels[config.LogLevel] {
		panic(fmt.Errorf("log-level must be either %s or %s", InfoLogLevel, DebugLogLevel))
	}
}

func validateServersTLSConfiguration(server Server, name string) error {
	if server.TLS == nil {
		return nil
	}
	requiredFiles := map[string]string{
		"cert": server.TLS.CertFile,
		"key":  server.TLS.KeyFile,
	}
	for field, path := range requiredFiles {
		if path == "" {
			return fmt.Errorf("servers.%s.tls.%s must be set", name, field)
		}
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("servers.%s.tls.%s=%s couldn't be accessed: %v", name, field, path, err)
		}
	}
	return nil
}

func getDefaultTargets() []string {
	return []string{
		"https://registry.terraform.io/",
	}
}
