package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Your string: ")
	if !scanner.Scan() { //он отдает смог или не смог считать sticky_error
		if err := scanner.Err(); err != nil {
			log.Fatal()
		} else {
			log.Fatal()
		}
	}
	fmt.Println("Вы ввели:", scanner.Text())
}
