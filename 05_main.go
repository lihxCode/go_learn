package main

import (
	"fmt"
	"math"
)

func main() {
	days := 365.2425
	println(days)
	var answer float64 = 42
	println(answer)

	var pi32 float32 = math.Pi
	println(pi32)
	var pi64 = math.Pi
	println(pi64)

	var num = 0.3333333
	fmt.Printf("%05.2f", num)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank == 0.3)
	fmt.Println(math.Abs(piggyBank-0.3) < 0.0001)
}
