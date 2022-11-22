package http

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewRouter(greetMiddleware gin.HandlerFunc, debugMode bool, logger *zap.SugaredLogger) (router *gin.Engine) {
	if !debugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router = gin.New()
	router.Use(ginzap.Ginzap(logger.Desugar(), time.RFC3339, true), ginzap.RecoveryWithZap(logger.Desugar(), false))
	router.GET("/", greetMiddleware)
	router.GET("/:name", greetMiddleware)
	return
}
