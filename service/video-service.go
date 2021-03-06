package service

import (
	"github.com/kazuki0924/go-gin/entity"
	"github.com/kazuki0924/go-gin/repository"
)

// VideoService interface
type VideoService interface {
	Save(entity.Video) entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
}

type videoService struct {
	// videos []entity.Video
	videoRepository repository.VideoRepository
}

// New VideoService
func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	// service.videos = append(service.videos, video)
	service.videoRepository.Save(video)
	return video
}

func (service *videoService) Update(video entity.Video) {
	service.videoRepository.Update(video)
}

func (service *videoService) Delete(video entity.Video) {
	service.videoRepository.Delete(video)
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()
}
