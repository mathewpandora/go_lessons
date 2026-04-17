package main

import "fmt"

func main() {
	PrintReplaced("Кукушка!")
}

func PrintReplaced(str string) {
	var res string
	for _, sym := range str {
		if string(sym) == "у" {
			res += "а"
		} else {
			res += string(sym)
		}
	}
	fmt.Println(res)
}
