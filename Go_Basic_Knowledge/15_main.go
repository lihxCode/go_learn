package main

import "fmt"

///切片 append函数 添加元素
///限制新增大小[0:4:4] 最后一个为新增容量大小
///make函数对slice进行预分配
func main() {
	drawfs := []string{"Jimmy", "Tony", "Keins"}
	drawfs = append(drawfs, "Kx")
	drawfs = append(drawfs, "Kx", "FD", "wonder")
	drawfs1 := append(drawfs, "Kx", "FD", "wonder")
	drawfs2 := append(drawfs, "456")
	drawfs1[0] = "111"
	// drawfs2[0] = "222"
	fmt.Println(drawfs)
	fmt.Println(drawfs1)
	fmt.Println(drawfs2)

	all := sum(10, 20, 30, 40, 50)
	println(all)
	sliceInt := []int{10, 20, 30, 40, 50}
	all = sum(sliceInt...)
	println(all)
}

///可变参数函数
////类型其实为切片
func sum(num ...int) int {
	var sum int
	for i := range num {
		sum += num[i]
	}
	return sum
}
