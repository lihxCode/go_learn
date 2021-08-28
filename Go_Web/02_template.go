package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func handleTemplate(w http.ResponseWriter, f *http.Request) {
	t, err := template.ParseFiles("./views/02_template.tmpl")
	if err != nil {
		fmt.Println("template parsefile failed, error:", err)
		return
	}
	// name := "Go Learn"
	user := UserInfo{
		Name:   "Jimmy",
		Gender: "Male",
		Age:    19,
	}
	t.Execute(w, user)
}

func main() {
	http.HandleFunc("/", handleTemplate)
	http.ListenAndServe(":80", nil)
}
