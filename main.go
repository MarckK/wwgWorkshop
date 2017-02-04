package main

import (
	"encoding/json"
	"net/http"

	"github.com/MarckK/wwg/animals"
)

var pets []animals.Pet

func main() {
	pets = []animals.Pet{
		&animals.Kitten{
			Name: "Tigg",
			Hobbies: []string{
				"Playing with wool",
				"Eating",
			},
		},
		&animals.Kitten{
			Name: "Pearl",
			Hobbies: []string{
				"Hunting mice",
				"Climbing trees",
			},
		},
		&animals.Dog{
			Name: "Beau",
			Hobbies: []string{
				"Barking",
				"Eating",
			},
		},
	}

	http.HandleFunc("/list", ListKittens)

	http.ListenAndServe(":9000", http.DefaultServeMux)

}

func ListKittens(rw http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(pets)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Write(data)
}
