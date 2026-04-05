package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
	fmt.Println(math.E)
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Sqrt(16))
	fmt.Println(math.Abs(-10.5))
	fmt.Println(math.Sin(math.Pi / 2))
	fmt.Println(math.Min(1, 2))
	//округление
	i := 0.1
	x := i + 0.2
	fmt.Println(x)             //0.30000000000000004 изза памяти
	fmt.Println(math.Round(x)) // по правилам матетмаикаи округлеям
	fmt.Println(math.Floor(x)) //  вниз
	fmt.Println(math.Ceil(x))  // верх
	fmt.Println(math.Trunc(x)) //округлить и остапвить флоат
	//до какого то знака запятой (2)
	num := 3.14159265 * 100
	fmt.Println(math.Round(num) / 100)
}
