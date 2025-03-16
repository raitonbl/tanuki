package context

import (
	"github.com/raitonbl/tanuki/internal/config"
	"go.uber.org/zap"
)

type Context interface {
	Logger() *zap.Logger
	Configuration() config.Config
}
