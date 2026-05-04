package main

import "fmt"

func main() {

	// := struct{}{} //пустая структура и весит ноль

	slice := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(notUnique(slice))
}

func notUnique(s []int) bool {
	counter := make(map[int]int) //тут можно было использовать пустую структурк и проверять была ли она ужде в мапе
	for _, v := range s {
		counter[v]++
		if counter[v] > 1 {
			return true
		}
	}
	return false
}
