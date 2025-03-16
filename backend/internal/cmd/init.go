package cmd

import "github.com/spf13/cobra"

func NewCommands() []*cobra.Command {
	return []*cobra.Command{
		NewServe(),
	}
}
