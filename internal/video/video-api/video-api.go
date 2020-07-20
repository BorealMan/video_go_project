package videoapi

import (
	"github.com/borealman/internal/video"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindALL() []video.Video
	Save(c *gin.Context)
}

func Routes()
