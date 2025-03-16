package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

const (
	EnvironmentVariablePrefix            = "TANUKI_SERVER_"
	ConfigurationFileEnvironmentVariable = EnvironmentVariablePrefix + "CONFIGURATION_FILE"
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

	configFilePath := viperInstance.GetString(ConfigurationFileEnvironmentVariable)
	isConfigFileDefined := configFilePath != ""
	if !isConfigFileDefined {
		configFilePath = "local.yaml"
	}
	_, err := os.Stat(configFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return Config{}, err
		}
		if isConfigFileDefined {
			return Config{}, err
		}
		viperInstance.SetConfigFile(configFilePath)
		if err = viperInstance.ReadInConfig(); err != nil {
			return Config{}, err
		}
	}
	var config Config
	if err = viperInstance.Unmarshal(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}
