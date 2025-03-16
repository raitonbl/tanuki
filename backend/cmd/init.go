package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "tanuki",
		Short: "Tanuki server command",
	}

	var serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start the Tanuki server",
		Run: func(cmd *cobra.Command, args []string) {
			config := loadConfig()
			fmt.Printf("Starting server with config: %+v\n", config)
		},
	}
}
