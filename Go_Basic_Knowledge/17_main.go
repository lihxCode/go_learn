package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	lat  float64 `json:"lattitue"`
	long float64 `json:"longitude"`
}

type dis struct {
	lat  float64
	long float64
}

func main() {
	var curiosity struct {
		lat  float64
		long float64
	}
	curiosity.lat = -4.5858
	curiosity.long = 137.4884
	fmt.Println(curiosity.lat)
	fmt.Println(curiosity.long)

	var spirit location
	spirit.lat = -16.83
	spirit.long = 17.313

	var opportunity location
	opportunity.lat = -39.31
	opportunity.long = 18.9301

	opportunity1 := location{lat: -1.13, long: 37.11}
	fmt.Println(opportunity1)

	///struct复制
	opportunity2 := opportunity1
	opportunity2.lat = 10000
	fmt.Println(opportunity2.lat)
	fmt.Println(opportunity1.lat)

	byte, err := json.Marshal(opportunity1) //字段名需要为大写才能导出，小写不能导出
	existOnError(err)
	fmt.Println(string(byte))
}

func existOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func distance(lat1, long1, lat2, long2 float64) float64 {
	return 0.0
}
