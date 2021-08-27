package main

import "fmt"

func main() {
	answer := 42
	fmt.Println(&answer)
	///无法获取&42之类字面量的值

	address := &answer
	fmt.Println(*address)

	///将*放在类型前面表示声明指针类型
	///将*放在变量前面表示解引用操作

	type person struct {
		name, superpower string
		age              int
	}
	timmy := &person{
		name: "Timothy",
		age:  10,
	}
	timmy.superpower = "flying"
	fmt.Println(timmy)
}
