package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const OUTPUT_PATH = "public/download"

func main() {
	url := "https://picsum.photos/200/300"

	err := os.MkdirAll(OUTPUT_PATH, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	destination, err := os.Create(fmt.Sprintf("%s/image_%d.%s", OUTPUT_PATH, time.Now().Unix(), "jpg"))
	if err != nil {
		log.Println(err)
	}

	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	_, err = io.Copy(destination, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Download succeed")
}
