package service

import (
	"fmt"

	"github.com/kazuki0924/go-gin/entity"
)

// VideoService interface
type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

// New VideoService
func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	fmt.Println(service.videos)
	return service.videos
}
