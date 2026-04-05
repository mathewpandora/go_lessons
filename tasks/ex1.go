package main

import (
	"fmt"
	"strconv"
)

func main() {
	quantityStr, priceStr := "100", "5"
	quantityInt, _ := strconv.Atoi(quantityStr)
	priceFloat, _ := strconv.ParseFloat(priceStr, 64)
	quantityFloat := float64(quantityInt)
	//fmt.Println(quantityFloat, reflect.TypeOf(quantityFloat), priceFloat, reflect.TypeOf(quantityFloat))
	totalPrice := quantityFloat * priceFloat
	totalPrice = totalPrice * 100
	fmt.Printf("%.2f", totalPrice) //  %-начало формата . - разделитель 2 - колво знака после запятой f - float
}
