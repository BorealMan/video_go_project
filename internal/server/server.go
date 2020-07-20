package server

import (
	"github.com/gin-gonic/gin"
)

func Run() error {
	router := gin.Default()
	router.Run(":8080")
	return nil
}
