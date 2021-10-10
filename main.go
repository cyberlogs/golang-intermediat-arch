package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	First string `json:"first"`
}

func main() {

	http.HandleFunc("/encode", foo)
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
