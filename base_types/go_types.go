package main

import (
	"fmt"
	"reflect"
)

func main() { //почему main подчеркнут

	v := 453
	fmt.Println(reflect.TypeOf(v), v)

	/*
		числа:

		int
		int8
		int16
		int32
		int64

		uint - только положительные
		uint8
		uint16
		uint32
		uint64

		float
		float32
		float64

		complex64
		complex128



		bool true false

		Составные:

		Array
		Slice
		Map
		Sctruct
		string - неизменяемые

		Function

		interface

		pointer

	*/
}
