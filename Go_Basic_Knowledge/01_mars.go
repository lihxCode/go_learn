package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("my weight on the surface of Mars is ")
	fmt.Print(140 * 0.3783)
	fmt.Print("lbs, and I would be ")
	fmt.Print(41 * 365 / 867)
	fmt.Print(" years old")
	fmt.Println("")
	fmt.Println("my weight on the surface of Mars is ", 140*0.3783, "lbs, and I would be ", 41*365/867, " years old")

	printf()
	variable()
	calculator()
	random()
}

func printf() {
	fmt.Printf("hello %v", 123)
}

func variable() {
	const v1 = 120
	var v2 = 100
	println(v1)
	println(v2)

	var (
		distance = 1000
		speed    = 100
	)
	println(distance)
	println(speed)

	distance = 2000
	println(distance)
}

func calculator() {
	var age = 10
	age = age + 1
	age++
	println(age)
}

func random() {
	var num = rand.Intn(10) + 1
	println(num)
	num = rand.Intn(10) + 1
	println(num)
}
