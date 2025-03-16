package cmd

import (
	"github.com/raitonbl/tanuki/internal/config"
	"github.com/raitonbl/tanuki/internal/context"
	"github.com/raitonbl/tanuki/internal/registry"
	"github.com/spf13/cobra"
)

func NewServe() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serves both the service port and the management port",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.NewConfigurationFromFlags(cmd.Flags())
			if err != nil {
				return err
			}
			return serve(cfg)
		},
	}
	return cmd
}

func serve(cfg config.Config) error {
	ctx := context.NewContext(cfg)
	err := registry.ListenAndServe(ctx)
	if err != nil {
		ctx.Logger().Error(err.Error())
	}
	//TODO: Create management server
	return nil
}
