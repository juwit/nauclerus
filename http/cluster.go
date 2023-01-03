package http

import (
	"github.com/gin-gonic/gin"
	"github.com/juwit/nauclerus/service"
	"go.uber.org/zap"
	"net/http"
)

func NewClusterMiddleware(cluster *service.Cluster, logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Debugf("listing namespaces")
		namespaces := cluster.ListNamespaces()
		c.JSON(http.StatusOK, gin.H{
			"namespaces": namespaces,
		})
	}
}
