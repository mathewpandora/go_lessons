package main

func main() {

}

func isPalindrome(mass []int) bool {

	for i := 0; i < len(mass); i++ {
		if mass[i] != mass[len(mass)-i] {
			return false
		}
	}
	return true
}
