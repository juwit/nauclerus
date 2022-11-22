package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juwit/nauclerus/service"
	"go.uber.org/zap"
)

func NewGreetMiddleware(greetService service.GreetService, logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		logger := logger.With("name", name)
		logger.Debugf("received a greeting request")
		greeting, err := greet(greetService, name)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": greeting})
	}
}

func greet(greetService service.GreetService, name string) (string, error) {
	if name == "" {
		return greetService()
	}
	return greetService(name)
}
