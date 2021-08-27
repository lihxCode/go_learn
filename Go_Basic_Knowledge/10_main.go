package main

///函数
func main() {
	maxNum := maxNum(10, 20)
	println(maxNum)
	minNum := MinNum(10, 20)
	println(minNum)
}

///语法 func 函数名（参数1 参数类型，参数2 参数类型）返回值类型 {函数体}
func maxNum(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

///大写开头，可以被导出，在其他包中使用
func MinNum(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
