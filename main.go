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
			Like: 0,
		},
		&animals.Kitten{
			Name: "Pearl",
			Hobbies: []string{
				"Hunting mice",
				"Climbing trees",
			},
			Like: 0,
		},
		&animals.Dog{
			Name: "Beau",
			Hobbies: []string{
				"Barking",
				"Eating",
			},
			Like: 0,
		},
	}

	http.HandleFunc("/list", ListPets)

	http.HandleFunc("/like", LikedPet)

	http.ListenAndServe(":9000", http.DefaultServeMux)

}

func ListPets(rw http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(pets)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Write(data)
}

func LikedPet(rw http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	for _, pet := range pets {
		if pet.GetName() == name {
			pet.IncrementLikeCounter()
		}
	}

	data, err := json.Marshal(pets)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Write(data)
}
