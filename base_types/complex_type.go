package main

//комплексные числа
import (
	"fmt"
)

func main() {
	var c64 complex64 = 1.0 + 2.3i  //64 бита
	var c32 complex128 = 1.0 + 2.3i //128 бита
	a := complex(6.5, 2)            //если передать первым агрументов float32 - бдует complex64 ; float64 - бдует complex128
	fmt.Println(c64, c32)
	c := a * c32
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(real(c), ":", imag(c))
}
