package server

import (
	"github.com/gin-gonic/gin"
)

func Run() error {
	router := gin.Default()

	//	Routes
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.Run(":8080")
	return nil
}
