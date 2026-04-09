package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("Введите пароль:")
	var pass string
	fmt.Scan(&pass)

	isSecured := scurePassword(pass)

	switch isSecured {
	case true:
		fmt.Println("Защищен")
	case false:
		fmt.Println("не защиен")
	}
}

func scurePassword(pass string) (isSecure bool) {
	return utf8.RuneCountInString(pass) >= 6 &&
		pass != strings.ToLower(pass) &&
		pass != strings.ToUpper(pass)
}

func scurePassword1(pass string) (isSecure bool) {
	if utf8.RuneCountInString(pass) < 6 {
		return false
	}
	if pass == strings.ToLower(pass) {
		return false
	}
	if pass == strings.ToLower(pass) {
		return false
	}
	return true
}
