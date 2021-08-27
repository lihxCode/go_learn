package main

import "fmt"

type person struct {
	age  int
	name string
}

func upAge(p *person) {
	p.age++
}

type Teacher struct {
	p    person
	book string
}

func main() {
	p := person{10, "jimmy"}
	fmt.Println(p.age)
	upAge(&p)
	fmt.Println(p.age)

	t := Teacher{person{10, "Tony"}, "Math"}
	fmt.Println(t)
	upAge(&t.p)
	fmt.Println(t)

	///通过指针修改数组

	var board [8][8]rune
	reset(&board)
	fmt.Printf("%c", board[0][0])
}

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
}
