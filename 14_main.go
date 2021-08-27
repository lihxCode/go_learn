package main

import (
	"fmt"
	"sort"
)

///切片 就是指向数组的窗口
///忽略起始索引，表示从数组的起始位置进行切分
///忽略掉结束索引，表示以数组的长度作为结束索引
///切片切字符串，是按字节来算，而不是字符
///函数以slice作为参数而不是数组
func main() {
	planet := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Urans",
		"Neptune",
	}
	slice1 := planet[0:4] ///包括0但是不包括4
	slice2 := planet[4:6]
	slice3 := planet[6:8]
	slice4 := planet[6:]
	slice5 := planet[:4]
	slice6 := planet[:]
	fmt.Println(slice1, slice2, slice3, slice4, slice5, slice6)

	///声明slice的两种方式
	slicePlanet1 := planet[:]                    ///切分数组
	slicePlanet2 := []string{"a", "b", "c", "d"} ///直接声明 类似于数组
	fmt.Println(slicePlanet1, slicePlanet2)

	sort.StringSlice(slicePlanet1).Sort()
	fmt.Println(slicePlanet1)
}
