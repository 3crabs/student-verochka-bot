package image_service

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
)

type newYearConfig struct {
	Images []string `json:"images"`
}

type imageService struct {
	images []string
}

func NewImageService() (ImageService, error) {
	images, err := readImagesFromConfig()
	if err != nil {
		return nil, err
	}
	return &imageService{images: images}, nil
}

func (is *imageService) GetRandomImageURL() string {
	return is.images[rand.Intn(len(is.images))]
}

func readImagesFromConfig() ([]string, error) {
	file, err := os.Open("config/new_year_config.json")
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	newYearConfig := &newYearConfig{}
	if err := json.Unmarshal(bytes, newYearConfig); err != nil {
		return nil, err
	}
	return newYearConfig.Images, nil
}
