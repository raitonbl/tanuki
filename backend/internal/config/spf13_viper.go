package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
	"os"
	"strings"
)

const (
	EnvironmentVariablePrefix            = "TANUKI"
	ConfigurationFileEnvironmentVariable = EnvironmentVariablePrefix + "_CONFIGURATION_FILE"
)

const (
	DefaultServerPort     = 8443
	DefaultManagementPort = 8080
	DefaultServiceName    = "tanuki"
	DefaultEnvironment    = "development"
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
	if err = assertServerTLSConfiguration(&instance.Servers.Registry, "registry"); err != nil {
		return Config{}, err
	}
	if err = assertServerTLSConfiguration(&instance.Servers.Management, "management"); err != nil {
		return Config{}, err
	}
	if err = assertLogLevel(instance); err != nil {
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
	for key := range getSpf13FlagList() {
		_ = v.BindPFlag(key, flags.Lookup(key))
	}
}

func getSpf13FlagList() map[string]string {
	return map[string]string{
		// Database configuration
		"database.username": "Username for database authentication",
		"database.password": "Password for database authentication",
		"database.name":     "Name of the database",
		"database.host":     "Hostname or IP address of the database server",
		"database.port":     "Port on which the database is running",
		"database.type":     "Type of the database (e.g., postgres, mysql)",
		"database.options":  "Additional connection options for the database",
		// Server ports configuration
		"server.registry.port":       "Port for the registry server",
		"server.registry.tls.key":    "TLS key file for the registry server",
		"server.registry.tls.cert":   "TLS certificate file for the registry server",
		"server.management.port":     "Port for the management server",
		"server.management.tls.key":  "TLS key file for the management server",
		"server.management.tls.cert": "TLS certificate file for the management server",
		// Global configuration
		"log-level": "Logging level (e.g., debug, info, warn, error)",
		"target":    "Specifies the target registry where the requests are forwarded to",
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
	if config.Targets == nil || len(config.Targets) == 0 {
		config.Targets = getDefaultTargets()
	}
	if config.Servers.Registry.Port == nil {
		defaultPort := DefaultServerPort
		config.Servers.Registry.Port = &defaultPort
	}
	if config.Servers.Management.Port == nil {
		defaultPort := DefaultManagementPort
		config.Servers.Management.Port = &defaultPort
	}
}

func assertLogLevel(config Config) error {
	if funk.Contains([]string{DebugLogLevel, InfoLogLevel}, config.LogLevel) {
		return nil
	}
	return fmt.Errorf("log-level must be either %s or %s", InfoLogLevel, DebugLogLevel)
}

func assertServerTLSConfiguration(server *Server, name string) error {
	if server.TLS == nil || (server.TLS.KeyFile == "" && server.TLS.CertFile == "") {
		if server.TLS != nil {
			server.TLS = nil
		}
		return nil
	}
	requiredFiles := map[string]string{
		"cert": server.TLS.CertFile,
		"key":  server.TLS.KeyFile,
	}
	for field, path := range requiredFiles {
		if path == "" {
			return fmt.Errorf("server.%s.tls.%s must be set", name, field)
		}
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("server.%s.tls.%s=%s couldn't be accessed: %v", name, field, path, err)
		}
	}
	return nil
}

func getDefaultTargets() []string {
	return []string{
		"https://registry.terraform.io/",
	}
}
