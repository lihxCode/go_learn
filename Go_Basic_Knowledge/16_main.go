package main

import "fmt"

///map 不会产生副本

func main() {
	tempature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	fmt.Println(tempature)
	earth := tempature["Earth"]
	fmt.Println(earth)

	tempature["Earth"] = 16
	fmt.Println(tempature)
 
	tempature["Venus"] = 20
	fmt.Println(tempature)

	moon := tempature["Moon"]
	fmt.Println(moon)

	if moon, ok := tempature["Moon"]; ok {
		fmt.Println(moon)
	} else {
		fmt.Println("no moon data")
	}

	///使用make预分配
	///第二参数为指定数量的key预分配空间
	tempature1 := make(map[float64]int, 8)
	fmt.Println(tempature1)

	tempature2 := []float64{-28.0, 32.0}
	frequency := make(map[float64]int)
	for _, t := range tempature2 {
		frequency[t]++
	}
	for key, value := range frequency {
		fmt.Printf("%+.2f occurs %d times\n", key, value)
	}
}
