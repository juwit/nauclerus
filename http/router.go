package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() (router *gin.Engine) {
	router = gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})
	return
}
