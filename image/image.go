package image

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

// Структура всех картинок
type Images struct {
	Images []struct {
		Url string `json:"url"`
	} `json:"images"`
}

type imageService struct {
	images []string
}

func (is *imageService) GetImage() string {
	index := rand.Intn(len(is.images))
	return is.images[index]
}

func NewImageService() (*imageService, error) {
	images, err := readImagesFromConfig()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &imageService{images: images}, nil
}

func readImagesFromConfig() ([]string, error) {
	jsonFile, err := os.Open("config/newYearImages.json")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	images := &Images{}
	if err := json.Unmarshal(byteValue, images); err != nil {
		log.Println(err)
		return nil, err
	}
	var urls []string
	for _, image := range images.Images {
		urls = append(urls, image.Url)
	}
	return urls, nil
}
