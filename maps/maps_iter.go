package main

import "fmt"

func main() {
	m := map[string]int{
		"one": 1,
		"two": 2, //неупорядоченые !!!
	}

	for k, v := range m {
		fmt.Printf("Key %s Value %d\n", k, v)
	}
}
