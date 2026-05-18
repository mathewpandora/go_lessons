package main

import "fmt"

func main() {
	count := 0
	for {
		fmt.Println(count)
		if count >= 3 {
			break
		} else if count == 2 {
			continue
		}
	}

	//return и цикл - если в цикле return то на этом моменте функция отдаст значение 
}
