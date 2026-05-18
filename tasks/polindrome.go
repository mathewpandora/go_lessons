package main

import "fmt"

func main() {
	mass := [6]int{1,2,3,3,2,1}
	fmt.Println(isPalindrome(mass))
}

func isPalindrome(mass [10]int) bool {

	for i := 0; i < len(mass); i++ {
		if mass[i] != mass[len(mass)-i] {
			return false
		}
	}
	return true
}
