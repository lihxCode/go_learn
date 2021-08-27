package main

import "fmt"

func main() {
	pace := "peace"
	fmt.Println(pace)
	fmt.Println("string can span multiple lines with the \n escape sequence")
	fmt.Println(`string can span multiple lines with the \n escape sequence`)

	runeType()
	character()
}

func runeType() {
	//rune 可以和int32互换使用
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Printf("%v %v %v %v \n", pi, alpha, omega, bang)
	fmt.Printf("%c %c %c %c \n", pi, alpha, omega, bang)
}

func character() {
	// var grade rune
	grade := 'A'
	fmt.Println(grade)
	fmt.Printf("%c", grade)
}
