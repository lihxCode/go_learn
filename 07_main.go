package main

import (
	"fmt"
	"math/big"
)

func main() {
	//很大的数字
	//超过10的18次方使用big
	var distance = 24e18
	fmt.Println(distance)
	lightSpeed := big.NewInt(2999)
	secondPerDay := big.NewInt(86499)
	println(lightSpeed, secondPerDay)

	num := new(big.Int)
	num.SetString("1000000000000000000000000000", 10)
	println(num)
}
