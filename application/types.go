package application

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Application struct {
	Logger    *zap.SugaredLogger
	DebugMode bool
	Router    *gin.Engine
}
