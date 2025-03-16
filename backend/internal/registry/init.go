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
	api := r.Group("/v1")
	routeContext := &DefaultContext{
		context: ctx,
		logger:  logger,
	}
	for _, f := range seq {
		f(routeContext, api)
	}
	return r.Run(fmt.Sprintf(":%d", *cfg.Servers.Registry.Port))
}
