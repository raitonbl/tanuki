package main

import (
	"github.com/raitonbl/tanuki/internal/cmd"
	"github.com/spf13/cobra"
	"log"
)

var (
	version string = "0.0.1"
)

func main() {
	app := &cobra.Command{
		Use:     "tanuki",
		Short:   "A command line service that serves a terraform registry proxy",
		Version: version,
	}
	app.AddCommand(cmd.NewCommands()...)
	if err := app.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
