package main

import (
	"log"
	"mini-lib/repository"
	"mini-lib/worker"
)

func main() {

	repo := repository.NewBookInMemoryRepository()

	//dependency injection - внедрение зависимости - когда мы что то передаем в стуктуру что ей нужно

	//можно сделать более гибко передавать интерфейс
	wrkr, err := worker.New(repo)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	wrkr.CreateBook()

	wrkr.PrintBook(1)

}
