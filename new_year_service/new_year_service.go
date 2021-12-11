package new_year_service

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
)

type newYearConfig struct {
	Images   []string         `json:"images"`
	Messages []NewYearMessage `json:"messages"`
}

type NewYearMessage struct {
	Header string `json:"header"`
	Text   string `json:"text"`
}

type newYearService struct {
	images   []string
	messages []NewYearMessage
}

func NewNewYearService() (NewYearService, error) {
	images, messages, err := readImagesFromConfig()
	if err != nil {
		return nil, err
	}
	return &newYearService{images: images, messages: messages}, nil
}

func (is *newYearService) GetRandomImageURL() string {
	return is.images[rand.Intn(len(is.images))]
}

func (is *newYearService) GetRandomMessage() NewYearMessage {
	return is.messages[rand.Intn(len(is.messages))]
}

func readImagesFromConfig() ([]string, []NewYearMessage, error) {
	file, err := os.Open("config/new_year_config.json")
	if err != nil {
		return nil, nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, nil, err
	}
	newYearConfig := &newYearConfig{}
	if err := json.Unmarshal(bytes, newYearConfig); err != nil {
		return nil, nil, err
	}
	return newYearConfig.Images, newYearConfig.Messages, nil
}
