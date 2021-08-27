package main

import "math/rand"

var era = "AD" ///作用于在main这个包内都是可见的，短声明不可以声明package变量

func main() {
	var count = 10
	for count < 10 {
		var num = rand.Intn(10) + 1
		println(num)
		count--
	}

	//短声明
	for num := 10; num > 0; num-- {
		println(num)
	}
	///在for 和swith if中没法使用var来声明变量

}
