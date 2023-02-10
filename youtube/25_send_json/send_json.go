package main

import (
	"encoding/json"
	"net/http"
)

type Course struct {
	// Atributo en minusculas es privado
	// Atributo en capitalize es públic
	Title      string `json:"title"`
	VideoCount int    `json:"count"`
}

type Courses []Course

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// course := Course{"Curso de Go", 28}
		// json.NewEncoder(w).Encode(course)

		courses := Courses{
			{"Curso de Go", 28},
			{"Curso de Python", 31},
			{"Curso de Java", 25},
		}
		json.NewEncoder(w).Encode(courses)
	})
	http.ListenAndServe(":8000", nil)
}
