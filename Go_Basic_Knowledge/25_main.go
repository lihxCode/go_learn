package main

import (
	"fmt"
	"time"
)

///通道 channel
///可以在多个gotoutine之间传值
func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepGohper(i, c)
	}
	for i := 0; i < 5; i++ {
		gohperID := <-c
		fmt.Println("gohper", gohperID, "has finished sleeping")
	}
}

func sleepGohper(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, "snore...")
	c <- id
}
