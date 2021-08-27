package main

import "fmt"

func main() {

	value := deferAction()
	fmt.Println(value)
}

func deferAction() string {
	defer fmt.Println("close")
	return "hello world"
}
