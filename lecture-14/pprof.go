package main

import (
	"net/http"
	"regexp"
	"strings"

	_ "net/http/pprof"
)

var pattern = regexp.MustCompile("auctor")
var haystack = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras accumsan nisl et iaculis fringilla. Integer sapien orci, facilisis ut venenatis nec, suscipit at massa. Cras suscipit lectus non neque molestie, et imperdiet sem ultricies. Donec sit amet mattis nisi, efficitur posuere enim. Aliquam erat volutpat. Curabitur mattis nunc nisi, eu maximus dui facilisis in. Quisque vel tortor mauris. Praesent tellus sapien, vestibulum nec purus ut, luctus egestas odio. Ut ac ipsum non ipsum elementum pretium in id enim. Aenean eu augue fringilla, molestie orci et, tincidunt ipsum.
Nullam maximus odio vitae augue fermentum laoreet eget scelerisque ligula. Praesent pretium eu lacus in ornare. Maecenas fermentum id sapien non faucibus. Donec est tellus, auctor eu iaculis quis, accumsan vitae ligula. Fusce dolor nisl, pharetra eu facilisis non, hendrerit ac turpis. Pellentesque imperdiet aliquam quam in luctus. Curabitur ut orci sodales, faucibus nunc ac, maximus odio. Vivamus vitae nulla posuere, pellentesque quam posuere`

func FirstSubstring() {
	for {
		strings.Contains(haystack, "auctor")
	}
}

func FirstRegex() {
	for {
		_, _ = regexp.MatchString("auctor", haystack)
	}
}

func FirstPreRegex() {
	for {
		pattern.MatchString(haystack)
	}
}

func main() {
	go FirstSubstring()
	go FirstRegex()
	go FirstPreRegex()
	http.ListenAndServe("0.0.0.0:8181", nil)
}

/*
Бенчмарки

CPU
go tool pprof pprof_t http://localhost:8080/debug/pprof/profile?seconds=5


MEM
alloc_space — количество аллоцированных байт;
alloc_objects — количество аллоцированных объектов;

inuse_space — количество живых байт;
inuse_objects — количество живых объектов.


go tool pprof -inuse_space pprof_t http://localhost:8080/debug/pprof/heap

go tool pprof -alloc_space pprof_t http://localhost:8080/debug/pprof/heap
*/

/*
Трейс
запустить web приложение

curl http://localhost:8181/debug/pprof/trace?seconds=10 -o trace.out

go tool trace -http "127.0.0.1:8080" trace.out


бенч горутин
go tool pprof -http :8080 http://:8181/debug/pprof/goroutine
*/
