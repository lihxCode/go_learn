package main

import (
	"fmt"
	"strings"
)

var t interface {
	talk() string
}

type talker interface {
	talk() string
}

type martain struct{}

func (m martain) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

///接口
func main() {
	t = martain{} ///结构类型
	fmt.Println(t.talk())

	t = laser(3) ///基础类型
	fmt.Println(t.talk())

	shout(martain{})
	shout(laser(2))
}
