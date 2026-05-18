package main

//Суть в том что если денлаем кастомные ошибки - нужно делать методы Error (чтобы было ошибкой) , Unwrap чтобы работало с Из
//
import (
	"errors"
	"fmt"
)

func main() {

}

type NetworkError struct {
	Err  error
	Code int
}

func NewNetworkError(msg string, code int) *NetworkError {
	return &NetworkError{
		Err:  errors.New(msg),
		Code: code,
	}
}

func NewNetworkErrorWithErr(err error, code int) *NetworkError {
	return &NetworkError{
		Err:  err,
		Code: code,
	}
}

func (e NetworkError) Error() string {
	return fmt.Sprintf("network error: %s", e.Err)
}

type WorkerS struct {
}

func NewWorker() *WorkerS {
	return &WorkerS{}
}

func (WorkerS) DoWork() error {
	return errors
}
