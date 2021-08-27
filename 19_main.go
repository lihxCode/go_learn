package main

///普通写法
type report struct {
	sol       int
	tempature tempature
	location  location
}

///嵌入 可以直接调用对应的方法
///方法冲突多写本身的一个方法
type report1 struct {
	sol int
	tempature
	location
}

type tempature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func main() {

}
