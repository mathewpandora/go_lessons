package main

import (
	"fmt"
	"slices"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	slice = slice[:len(slice)-1] //удалить послдениц
	slice = slice[1:]            //удалить первый

	indexToRemove := 2
	slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
	fmt.Printf("1 cap of slice:%v is %v\n", slice, cap(slice))
	slice = slices.Clip(slice) //делает капасити = лен s[:len(s):len(s)] len(s) - капасити
	fmt.Printf("2 cap of slice:%v is %v\n", slice, cap(slice))
	//Память не обнулится просто будент новыц капасити а памяти он будет занимать столько же
	//если хотим с памятью тот надо через мейк сделать новый и в него скоприровать через копи

}
