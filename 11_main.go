package main

type celsius float64 //声明新类型
///声明新类型极大提高可读性

type kelvin float64

func main() {
	const degree = 20
	var temperature celsius = degree
	temperature += 10

	//不同类型不能混用
	var warmUp float64 = 10
	temperature += celsius(warmUp)
	println(temperature)

	var k kelvin = 294.0
	var tmp = k.celsius()
	println(tmp)
}

///(k kelvin) 为接收者  通过方法添加行为
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
