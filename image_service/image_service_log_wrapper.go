package image_service

import "log"

type imageServiceLogWrapper struct {
	service ImageService
}

func NewImageServiceLogWrapper(service ImageService) ImageService {
	return &imageServiceLogWrapper{service: service}
}

func (w *imageServiceLogWrapper) GetRandomImageURL() string {
	log.Println("get random image url")
	return w.service.GetRandomImageURL()
}
