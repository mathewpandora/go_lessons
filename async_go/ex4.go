package main

//Ебаенутая задача

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var GlMutex sync.Mutex
var GlCh chan int
var isUpdating bool

func main() {

	isUpdating = false

	http.HandleFunc("/weather", GetWeather)

	err := http.ListenAndServe("localhost:3333", nil)
	if err != nil {
		panic(err)
	}
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
	if isUpdating {
		degree := <-GlCh
		fmt.Fprintf(w, "temperatue: %d", degree)
	} else {
		GlCh = make(chan int)
		go WriterEx()
		degree := <-GlCh
		fmt.Fprintf(w, "temperatue: %d", degree)
	}

}

func WeatherApi() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}

func WriterEx() {

	GlMutex.Lock()
	if isUpdating {
		return
	} else {

		isUpdating = true

		go func() {

			degree := WeatherApi()
			GlCh <- degree
			GlMutex.Unlock()

		}()
	}

}

//Если одна горутина (например, первый запрос) вошла в защищённую секцию (вызвала Lock), то любая другая горутина, которая попытается вызвать Lock на том же мьютексе, приостановится и будет ждать, пока первая не вызовет Unlock.
