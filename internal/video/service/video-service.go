package service

import (
	"github.com/borealman/video_go_project/internal/video"
)

type VideoService interface {
	Save(video.Video) video.Video
	FindAll() []video.Video
}

type videoService struct {
	videos []video.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video video.Video) video.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []video.Video {
	return service.videos
}
