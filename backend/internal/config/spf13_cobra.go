package config

import "github.com/spf13/pflag"

func BindSpf13CobraFlags(flag *pflag.FlagSet) {
	for key, desc := range getSpf13FlagList() {
		switch key {
		case "database.port", "server.registry.port", "server.management.port":
			defaultsTo := -1
			if key == "server.registry.port" {
				defaultsTo = DefaultServerPort
			} else if key == "server.management.port" {
				defaultsTo = DefaultManagementPort
			}
			flag.Int(key, defaultsTo, desc)
		case "log-level":
			flag.String(key, InfoLogLevel, desc)
		case "target":
			flag.StringSlice(key, nil, desc)
		default:
			defaultsTo := ""
			flag.String(key, defaultsTo, desc)
		}
	}
}
