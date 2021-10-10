package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	First string `json:"first"`
}

func main() {

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := Person{
		First: "John",
	}

	p2 := Person{
		First: "Jane",
	}

	people := []Person{p1, p2}

	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Println("Encoded bad data!", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var people []Person
	err := json.NewDecoder(r.Body).Decode(&people)
	if err != nil {
		log.Println("Decode bad data!", err)
	}
	fmt.Println("People:", people)
}
