package main

import (
	"fmt"
	"net/http"
)

type Handler struct {
	params map[string]interface{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if p := r.URL.Query().Get("new"); p != "" {
		h.params["new"] = p
	}

	if v := r.FormValue("new2"); v != "" {
		h.params["new2"] = v
	}

	_, _ = fmt.Fprintf(w, "params: %v. URL: %s", h.params, r.URL.String())
}

func main() {
	hh := &Handler{
		map[string]interface{}{"age": 25, "name": "gopher"},
	}
	http.Handle("/with-param", hh)

	_ = http.ListenAndServe(":8080", nil)
}
