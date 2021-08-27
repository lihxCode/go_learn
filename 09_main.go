package main

import (
	"strconv"
	"strings"
)

func main() {
	//类型转换
	age := 41
	marsAge := float64(age)
	println(marsAge)

	earthDays := 365.2425
	println(int(earthDays))

	var pi rune = 960
	println(string(pi))

	countdown := 10
	str := "minutes" + strconv.Itoa(countdown) + "end"
	println(str)

	name := "LiSi"
	name = strings.ToLower(name) ///ToUpper
	println(name)

}
