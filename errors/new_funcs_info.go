package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrFileNotExists = errors.New("FIle not exitrs ")

func main() {

	if err := run(); err != nil {
		//errors.As — достаёт из ошибки данные определённого вида, чтобы их использовать.
		//errors.Is — просто проверяет, есть ли в цепочке ошибок нужная (сравнивая с образцом).
		//errors.As - чаще всего используют с катосмными ошибками
		//Разница: As даёт доступ к содержимому ошибки, Is только отвечает «да/нет».
		//errors.As - извлекает ошибку определенного типаё
		//в errors.As - передаем указатель чтобы записать в него данные ошибки
		//errors.Is - разспаознает только если использовать спецификатор %w
		if errors.Is(err, ErrFileNotExists) { //Тут модет быть проблема енсли будем сранивать  ErrFileNotExists - потому что она уже другая ("DoWork: %w", err)
			log.Fatalf("FIle not exitrs ")
		} else {
			log.Fatalf("run: %w", err)
		}
	}
}

type Worker struct {
}

func New() *Worker {
	return &Worker{}
}

func (Worker) DoWork() error {
	return ErrFileNotExists
}

func run() error {
	w := New()
	if err := w.DoWork(); err != nil {
		return fmt.Errorf("DoWork: %w", err)
	}
	return nil
}

type FileNotExistsError struct {
	Path string
}

func (e FileNotExistsError) Error() string {
	return fmt.Sprintf("FileNotExistsError %s", e.path)
}

func NewFileNotExistsError(path string) *FileNotExistsError { //Правильно в конструкрторе отдавать указатель чтолы потомтнормально работакть с ерор эс
	return &FileNotExistsError{
		Path: path,
	}
} //Верно ответили, errors.Is проверяет, является ли ошибка определенным типом, сравнивая её с заданной ошибкой, в то время как errors.As позволяет извлечь значение ошибки определенного типа, если оно присутствует.
