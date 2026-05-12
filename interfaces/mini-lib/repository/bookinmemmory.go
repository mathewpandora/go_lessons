package repository

import (
	"errors"
	"mini-lib/model"
)

type BookInMemoryRepository struct {
	books map[int]model.Book
}

func NewBookInMemoryRepository() *BookInMemoryRepository {
	return &BookInMemoryRepository{
		books: make(map[int]model.Book),
	}
}

func (b BookInMemoryRepository) ByID(ID int) (model.Book, error) {
	book, exists := b.books[ID]
	if !exists {
		return model.Book{}, errors.New("This book does not exists")
	}
	return book, nil
}

func (b *BookInMemoryRepository) Add(book model.Book) error {

	if _, exexists := b.books[book.ID]; exexists {
		return errors.New("Book with this ID already exists")
	}
	b.books[book.ID] = book
	return nil
}

func (b *BookInMemoryRepository) Delete(ID int) error {

	if _, exexists := b.books[ID]; exexists {
		return errors.New("Book with this ID already exists")
	}

	delete(b.books, ID)
	return nil
}
