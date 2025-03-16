package registry

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/raitonbl/tanuki/internal/context"
	"go.uber.org/zap"
	"time"
)

func ListenAndServe(ctx context.Context) error {
	cfg := ctx.Configuration()
	logger := ctx.Logger().With(
		zap.String("component", "registry"),
		zap.Int("port", *cfg.Servers.Registry.Port),
	)
	r := gin.Default()
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	seq := []func(context.Context, *gin.RouterGroup){
		setProvidersRoute,
	}
	api := r.Group("/registry.terraform.io")
	routeContext := &DefaultContext{
		context: ctx,
		logger:  logger,
	}
	for _, f := range seq {
		f(routeContext, api)
	}
	portDefinition := fmt.Sprintf("0.0.0.0:%d", *cfg.Servers.Registry.Port)
	if cfg.Servers.Registry.TLS != nil {
		return r.RunTLS(portDefinition, cfg.Servers.Registry.TLS.CertFile, cfg.Servers.Registry.TLS.KeyFile)
	}
	return r.Run(portDefinition)
}
