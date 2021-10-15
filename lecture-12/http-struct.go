package main

import (
	"fmt"
	"net/http"
)

type HalykHandler struct {
	Ipoteka bool
	Name    string

	isKnowGolang bool
}

func (h *HalykHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Ipoteka: %v, Name: %s, isKnowGolang: %v, URL: %s",
		h.Ipoteka, h.Name, h.isKnowGolang, r.URL.String())
}

func main() {
	hh := &HalykHandler{true, "Andrey", false}
	http.Handle("/andrey", hh)

	smartHh := &HalykHandler{false, "Student", true}
	http.Handle("/smart", smartHh)

	_ = http.ListenAndServe(":8080", nil)
}
