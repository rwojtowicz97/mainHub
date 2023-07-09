package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/os", func(c *gin.Context) {
		c.String(200, time.Now().Format(time.RFC850))
	})
	return router
}
