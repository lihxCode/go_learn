package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hi", hiHanlder)
	mux.Handle("/welcome/goweb", welComeHanlder{Name: "Hi, Go Handle"})
	server := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func hiHanlder(w http.ResponseWriter, f *http.Request) {
	fmt.Fprintf(w, "欢迎页面")
}

type welComeHanlder struct {
	Name string
}

func (h welComeHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi, %s", h.Name)
}
