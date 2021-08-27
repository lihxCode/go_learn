package main

import "math/rand"

func main() {
	var command = "go east"
	if command == "go east" {
		println("your head futher up the mountain")
	} else if command == "hello" {
		println("22")
	} else {
		println("33")
	}
	year()
	oppsite()
	switchCase()
	forCondition()
	guess()
}

func year() {
	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 == 0)
	if leap {
		println("闰年")
	} else {
		println("非闰年")
	}
}

func oppsite() {
	var haveTrouch = true
	var littrouch = false
	if !haveTrouch && !littrouch {
		println("is trouch")
	}
}
func switchCase() {
	var a = 100
	switch a {
	case 100:
		println("100")
		fallthrough ///继续执行往下的
	case 200:
		println("200")
	case 300:
		println("400")
	}
}

func forCondition() {
	var count = 10
	for count > 0 {
		println(count)
		count--
	}
	println("liftoff")
}

func guess() {
	var count = 33
	var num = rand.Intn(100)
	for {
		if num == count {
			println("猜对了 %v", num)
			break
		} else if num > count {
			println("猜大了 %v", num)
		} else {
			println("猜小了 %v", num)
		}
		num = rand.Intn(100)
	}
}
