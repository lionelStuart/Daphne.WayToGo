package main

import (
	"Live/controller"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr:        ":9000",
		Handler:     &controller.MyHandler{},
		ReadTimeout: 10 * time.Second,
	}

	server.ListenAndServe()
	//test.TestVideo()
}
