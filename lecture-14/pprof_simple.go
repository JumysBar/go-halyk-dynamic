package main

import (
	"net/http"
	"time"

	_ "net/http/pprof"
)

func CPUHogger() {
	var acc uint64
	t := time.Tick(2 * time.Second)
	for {
		select {
		case <-t:
			time.Sleep(50 * time.Millisecond)
		default:
			acc++
		}
	}
}

func main() {
	go CPUHogger()
	go CPUHogger()
	http.ListenAndServe("127.0.0.1:8181", nil)
}
