package model

import (
	"errors"
	"fmt"
)

type Book struct {
	ID     int
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("ID: %d, Title: %s, Author: %s", b.ID, b.Title, b.Author)
}

func NewBook(ID int, title, author string) (Book, error) { //Важно запомнить что структтура не может отдаваться как нил

	if title == "" {
		return Book{}, errors.New("Title could not be empty")
	} else if author == "" {
		return Book{}, errors.New("Author could not be empty")
	} else if ID < 0 {
		return Book{}, errors.New("ID could be more than 0")
	}

	NewBook := Book{
		ID:     ID,
		Title:  title,
		Author: author,
	}

	return NewBook, nil

}
