package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

///声明函数类型
type sensor func() kelvin

func main() {
	///函数赋给变量
	///函数作为参数
	///函数作为返回值
	sensor := fakeSensor
	println(sensor())

	sensor = realSensor
	println(sensor())

	f()

	k := func(message string) {
		fmt.Println("匿名函数" + message)
	}
	k("赋给了一个变量")

	func() {
		println("function anonymous")
	}() ///立即执行

	var num kelvin = 10
	sensorFunc := calibrate(fakeSensor, num)
	num++
	println(sensorFunc())
	println(sensorFunc())
	println(sensorFunc())
	println(sensorFunc())
	println(sensorFunc())
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return kelvin(rand.Intn(10))
}

func measureTempature(num int, sen sensor) {
	sen()
}

///匿名函数 没有名字的函数
var f = func() {
	fmt.Println("匿名函数 没有名字的函数")
}

///捕获变量 值传递
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}
