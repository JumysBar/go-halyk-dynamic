package main

import (
	"fmt"
	"net/http"
	"time"
)

// самый простой пример вашего личного web сервера
func main() {

	// функция котоаря помогает нам создать
	http.HandleFunc("/test",
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Second * 2)
			_, _ = w.Write([]byte("hi test"))
		})

	http.HandleFunc("/прикол/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header()
			_, _ = fmt.Fprintln(w, "hi UTF-8 из коробки:", r.URL.String())
		})

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("hi root"))
		})

	http.HandleFunc("/error",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			_, _ = w.Write([]byte("АШИБКА СЕРВЕРА"))
		})

	// запкст сервера в отдельной горутине
	go func() {
		_ = http.ListenAndServe(":8080", nil)
	}()

	// создание нового мультиплексера
	mux := http.NewServeMux()
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("hi root"))
		})

	mux.HandleFunc("/test",
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Second * 2)
			_, _ = w.Write([]byte("hi test"))
		})

	// собери свой WOK (сервер)
	anotherServer := http.Server{
		Addr:         ":8081",
		Handler:      mux,
		ReadTimeout:  time.Second * 1,
		WriteTimeout: time.Second * 1,
	}

	_ = anotherServer.ListenAndServe()
}
