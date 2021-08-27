package main

import "fmt"

func main() {
	var planets [8]string
	planets[0] = "Sun"
	planets[1] = "Earth"
	planets[2] = "Moon"

	earth := planets[1]
	println(earth)
	println(len(planets))
	println(planets[3] == "") ///没赋值的，默认为类型的零值

	// i := 8
	// planets[i] = "pluto"
	// pluto := planets[i]
	// println(pluto)

	anyelements()
	ergodic()
}

func anyelements() {
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
	fmt.Println(planet)
}

func ergodic() {
	drawfs := [5]string{"Ceres", "Pluto", "Hamuea", "Makemake", "Eirs"}
	for i := 0; i < len(drawfs); i++ {
		drawf := drawfs[i]
		println(drawf)
	}

	for i, drawf := range drawfs {
		println(i, drawf)
	}

	dummyDrawfs := drawfs //值拷贝
	drawfs[0] = "Dummy"
	fmt.Println(dummyDrawfs)
	fmt.Println(drawfs)
}
