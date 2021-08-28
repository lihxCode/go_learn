package main

import (
	"net/http"
)

type Refer struct {
	handlder http.Handler
	refer    string
}

func (rf *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		rf.handlder.ServeHTTP(w, r)
		w.Write([]byte("走了拦截器流程"))
	} else {
		w.WriteHeader(404)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is handler"))
}

func hello(w http.ResponseWriter, f *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	// server := &http.Server{
	// 	Addr: "0.0.0.0:80",
	// }
	// http.HandleFunc("/hello", hello)
	// server.ListenAndServe()

	referer := &Refer{
		handlder: http.HandlerFunc(myHandler),
		refer:    "123",
	}
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":80", referer)
}
