package registry

import (
	"github.com/raitonbl/tanuki/internal/config"
	"github.com/raitonbl/tanuki/internal/context"
	"go.uber.org/zap"
)

type DefaultContext struct {
	logger  *zap.Logger
	context context.Context
}

func (instance *DefaultContext) Logger() *zap.Logger {
	return instance.logger
}

func (instance *DefaultContext) Configuration() config.Config {
	return instance.context.Configuration()
}
