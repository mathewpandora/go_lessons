package main

import "fmt"

func main() {
	//масивы - фиксированная по размеру коллекция элементов одного типа
	var nums [5]int //масив нумс 6 интовых элементов
	//по умолчанию подставлюятся zero-value значения
	var massive = [5]int{1, 2, 3, 4, 5}

	massive2 := [...]string{"dfv", "tfrvtv"} //сам поймет сколько нужно

	var nums2 = [3]string{2: "fvfv", 1: "fdvfv"} //на индексе 2 и 1 будет :

	fmt.Println(nums, massive, massive2, nums2)
}
