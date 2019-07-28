package test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func TestVideo() {
	http.HandleFunc("/", serveHttp)
	http.ListenAndServe(":8080", nil)
}

func serveHttp(w http.ResponseWriter, r *http.Request) {
	path, _ := os.Getwd()
	fmt.Println("current dir", path)
	video, err := os.Open("upload/1564297711.mp4")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("current show video:", video.Name())
	http.ServeContent(w, r, "test.mp4", time.Now(), video)
}
