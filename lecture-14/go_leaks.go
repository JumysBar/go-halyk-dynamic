package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	_ "net/http/pprof"
)

func goLeak(c chan error) {
	time.AfterFunc(time.Second, func() {
		err := errors.New("время вышло")
		c <- err
		fmt.Println("TADA")
	})
}

func main() {

	go func() {
		http.ListenAndServe("127.0.0.1:8181", nil)
	}()

	var results chan error = nil

	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Millisecond * 10)
		go goLeak(results)
	}

	fmt.Println("из канала", <-results)
}
