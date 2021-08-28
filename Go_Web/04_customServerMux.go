package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHanlder)
	mux.HandleFunc("/hi", hiHanlder)
	mux.HandleFunc("/hi/web", webHanlder)

	server := &http.Server{
		Addr:         ":80",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	if error := server.ListenAndServe(); error != nil {
		log.Fatal(error)
	}
}

func hi(w http.ResponseWriter, f *http.Request) {
	fmt.Fprintf(w, "Hi Web")
}

func indexHanlder(w http.ResponseWriter, f *http.Request) {
	fmt.Fprintf(w, "首页 处理器为indexHanlder")
}

func hiHanlder(w http.ResponseWriter, f *http.Request) {
	fmt.Fprintf(w, "欢迎页面 处理器为hiHanlder")
}

func webHanlder(w http.ResponseWriter, f *http.Request) {
	fmt.Fprintf(w, "web页面 处理器为webHanlder")
}
