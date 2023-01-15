package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// http.ServeFile(w, r, "index.html")

		fmt.Println(r.URL.Path)
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.ListenAndServe(":8000", nil)
}
