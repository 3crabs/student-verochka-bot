package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

// Структура всех картинок
type Images struct {
	Images []Image `json:"images"`
}

// Структура картинки
type Image struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Type string `json:"type"`
}

// Возвращает байты картинки, сохранённой по передаваемому пути на сервере
func File(path string) []byte {
	photoBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return photoBytes
}

// Возвращает случайную ссылку на картинку из images.json
func RandomFileFromConfig() string {
	jsonFile, err := os.Open("config/images.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened images.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var images Images
	json.Unmarshal(byteValue, &images)

	countImages := len(images.Images)
	fileNumber := rand.Intn(countImages)
	if countImages < 1 {
		fmt.Println("Images.json is empty")
	}
	return images.Images[fileNumber].Url
}
