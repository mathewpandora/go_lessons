package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//Утчека горутины возникает когда результьат работыв горутины (уже) никому не нужен
	//Но горутина продолджает работать

	//func my_func() {

	//	go func() {
	//		for {}|
	//		select {} //select - блокирующая операция она ждет запись или чтение из канала но так как никаких кейсв нет она просто заблокируется в рамках горутины и горутина останется висеть
	//	}()
	//}

	//Профилактика утечек горутин
	//Завершать работу горутины когда все от нее получил
	//Завершать работу горутины если в ней произошла непопровимая ошибка
	//Выход по контекстуц
	//if err != nil

	ticker := time.NewTicker(time.Second)

	config := 3

	go func() {
		defer ticker.Stop()
		var counter int
		for {
			select {
			case <-ticker.C:
				counter++
				fmt.Println(counter)
				if counter == config {
					fmt.Println("connected")
					//return
				}
			}
		}
	}()

	server := http.Server{
		Addr: "localhost:8000",
	}

	server.ListenAndServe()

	//Базовый случацй lesson3_2

}
