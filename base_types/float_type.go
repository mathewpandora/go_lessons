package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		a float32
		b float64
	)

	//NaN  - относится в флоат
	//Infinity () - относится к флоат

	zeroInt := -1
	res := 1 / zeroInt //тут будет паника при делении на ноль

	zeroFloat64 := 0.0
	res2 := 1.0 / zeroFloat64              //так будет инф но если прям написать res2 : 0.0 - ошибка компилятора
	fmt.Println(zeroFloat64 / zeroFloat64) //NaN
	fmt.Println(res2)
	fmt.Println(res)
	fmt.Println(a, b)

	nan := math.NaN()
	fmt.Println(nan)
	fmt.Println("проверка", math.IsNaN(nan))

	var inf_pos = math.Inf(1)
	inf_minus := math.Inf(-1)
	fmt.Println(inf_pos, inf_minus) //не понятно почему если первый аргумент - float32 ошибка
	fmt.Println("Infinity : ", math.IsInf(inf_minus, 0))

	var (
		i = 12
		f = 1.23
	)

	res_new := f + float64(i)
	fmt.Println("Answeer:", res_new)
	fmt.Println("res")
	var i2 = int(math.Round(res_new))

	fmt.Println(i2)
	//тест
	fmt.Println("max float64:", math.MaxFloat64)
	fmt.Println("max float32:", math.MaxFloat32)
	fmt.Println("max uint16:", math.MaxUint16)
	fmt.Println("max int16:", math.MaxInt16)
}
