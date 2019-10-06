package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email`
}

var people []Person

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(people)
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var p Person
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	people = append(people, p)
	json.NewEncoder(w).Encode(people)
}

func Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	params := mux.Vars(r)
	for _, p := range people {
		if p.ID == params["id"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	params := mux.Vars(r)

	for i, p := range people {
		if p.ID == params["id"] {

			var p1 Person
			if err := json.NewDecoder(r.Body).Decode(&p1); err != nil {
				json.NewEncoder(w).Encode(err)
				return
			}
			people[i] = p1
			json.NewEncoder(w).Encode(p1)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {

	people = []Person{
		{"1", "Sangeet Kumar", "sangeet.kumar@gmail.com"},
		{"2", "John Doe", "john.doe@example.com"},
	}

	router := mux.NewRouter()

	router.HandleFunc("/people", List).Methods("GET")
	router.HandleFunc("/people", Create).Methods("POST")
	router.HandleFunc("/people/{id}", Read).Methods("GET")
	router.HandleFunc("/people/{id}", Update).Methods("POST")
	router.HandleFunc("/people/{id}", Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
