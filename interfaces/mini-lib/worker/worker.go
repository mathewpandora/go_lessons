package worker

import (
	"errors"
	"fmt"
	"mini-lib/model"
	"mini-lib/repository"
)

// Правило: если у структуры есть поля-указатели — всегда используйте *Worker для всех методов (даже читающих).

type repo interface {
	ByID(ID int) (model.Book, error)
	Add(book model.Book) error
	//Delete(ID int) error так как delete в воркере не используем поэтому его и не исползуе пишем самый минимум
}
type Worker struct {
	//repo *repository.BookInMemoryRepository
	repo repo
}

func New(repo *repository.BookInMemoryRepository) (*Worker, error) {
	if repo == nil {
		return nil, errors.New("repo is nil")
	}
	return &Worker{
		repo: repo,
	}, nil
}

func (w *Worker) CreateBook() error {

	b1, err := model.NewBook(1, "Book1", "Author1")

	if err != nil {
		return fmt.Errorf("Error while creating book %w", err)
	}

	if err := w.repo.Add(b1); err != nil {
		return fmt.Errorf("Error while adding book in repo: %w", err)
	}

	return nil
}

func (w *Worker) PrintBook(ID int) error {

	b, err := w.repo.ByID(ID)
	if err != nil {
		return fmt.Errorf("Repo can not find this book: %w", err)
	}
	fmt.Println(b)
	return nil
}
