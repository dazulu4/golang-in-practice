package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Respuesta en root")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hola mundo")
	})
	http.ListenAndServe(":8000", nil)
}
