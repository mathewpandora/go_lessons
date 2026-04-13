package main

//замыкание - когда функция может замкнуть какое то значение
import (
	"fmt"
	"log"
)

func main() {

	func() {
		fmt.Println("Это анонимная функция")
	}()

	ruFn, err := helloFactory("ru")
	if err != nil {
		log.Fatalf("helloFactory error: %s", err.Error()) //err.Error() стринга с ошибкой
	}

	ruFn("Павел") //тперь  ruFn - это функция которая принимает стритнгу

	ruFn, err = helloFactory("en")
	if err != nil {
		log.Fatalf("helloFactory error: %s", err.Error()) //err.Error() стринга с ошибкой
	}

	ruFn("Pavel") //тперь  ruFn - это функция которая принимает стритнгу

	ruFn, err = helloFactory("efrn")
	if err != nil {
		log.Fatalf("helloFactory error: %s", err.Error()) //err.Error() стринга с ошибкой
	}

	ruFn("Pavel") //тперь  ruFn - это функция которая принимает стритнгу
}

func helloFactory(lang string) (func(name string), error) { //выозвращает функцию
	// helloFactory - создает функцию которая на рпавильном язывке здоровается

	var message string
	switch lang {
	case "ru":
		message = "Привет, %s\n"
	case "en":
		message = "Hello, %s\n"
	default:
		return nil, fmt.Errorf("Такого языка нет")
	}

	return func(name string) { //message - не удалится в helloFactory таку как она передается в func в этом суть замыкания
		fmt.Printf(message, name) //в message лежит спецификатор и name подставдяется в message
	}, nil //message не удалится потому чтот message потом используется в main поэтому переменная замкнуласть в ней и не дает ей удалиться
}

//Замыкание — это ситуация, когда функция использует переменные из внешнего окружения, даже после того, как эта область видимости завершила свое выполнение. Тем самым, используемые внутри функции переменные не будут удалены, пока не будет удалена сама функция.
