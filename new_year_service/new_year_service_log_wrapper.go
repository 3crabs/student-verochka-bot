package new_year_service

import "log"

type newYearServiceLogWrapper struct {
	service NewYearService
}

func NewNewYearServiceLogWrapper(service NewYearService) NewYearService {
	return &newYearServiceLogWrapper{service: service}
}

func (w *newYearServiceLogWrapper) GetRandomImageURL() string {
	log.Println("get random image url")
	return w.service.GetRandomImageURL()
}

func (w *newYearServiceLogWrapper) GetRandomMessage() NewYearMessage {
	log.Println("get random message")
	return w.service.GetRandomMessage()
}
