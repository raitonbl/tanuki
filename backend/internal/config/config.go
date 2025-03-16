package config

const (
	InfoLogLevel  = "INFO"
	DebugLogLevel = "DEBUG"
)

type Config struct {
	Targets     []string `mapstructure:"target"`
	Servers     Servers  `mapstructure:"servers"`
	Database    Database `mapstructure:"database"`
	LogLevel    string   `mapstructure:"log-level"`
	Environment string   `mapstructure:"environment"`
	Service     string   `mapstructure:"service"`
	Solution    string   `mapstructure:"part-of"`
}

type Servers struct {
	Registry   Server `mapstructure:"registry"`
	Management Server `mapstructure:"management"`
}

type Server struct {
	Port *int       `mapstructure:"port"`
	TLS  *ServerTLS `mapstructure:"tls"`
}

type ServerTLS struct {
	KeyFile  string `mapstructure:"key"`
	CertFile string `mapstructure:"cert"`
}

type Database struct {
	Username         string `mapstructure:"username"`
	Password         string `mapstructure:"password"`
	Name             string `mapstructure:"name"`
	Host             string `mapstructure:"host"`
	Port             int    `mapstructure:"port"`
	ManagementSystem string `mapstructure:"type"`
	Options          string `mapstructure:"options"`
}
