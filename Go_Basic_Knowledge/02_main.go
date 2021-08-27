package main

import (
	"strings"
)

func main() {
	var command = "onUpdate"
	var exit = strings.Contains(command, "Update")
	println(exit)
}
