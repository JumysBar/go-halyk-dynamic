package main

import (
	"fmt"
	"net/http"
)

type PostHandler struct {
	params map[string]interface{}
}

func (h *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		_, _ = w.Write([]byte("Вы кто такие? я вас не звал!"))
		return
	}

	p := r.FormValue("password")
	if p == "" {
		http.Redirect(w, r, "/nope", http.StatusFound)
	}
	h.params["password"] = p

	_, _ = fmt.Fprintf(w, "params: %v. URL: %s", h.params, r.URL.String())
}

func main() {
	hh := &PostHandler{
		map[string]interface{}{},
	}
	http.Handle("/login", hh)

	http.HandleFunc("/nope",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(401)
			_, _ = w.Write([]byte("access denied"))
		})

	_ = http.ListenAndServe(":8080", nil)
}
