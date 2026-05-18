package main

import (
	"fmt"
	"log"
)

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: $s", e.Code, e.Message)
}

func main() {
	if err := SomeFunc(); err != nil {
		log.Fatal("Ошибка: %w", err)
	}
}

func SomeFunc() error {
	return &MyError{
		Code:    500,
		Message: "Error",
	}
}
