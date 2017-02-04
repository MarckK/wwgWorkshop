package main

import (
	"encoding/json"
	"net/http"

	"github.com/MarckK/wwg/animals"
)

var kittens []animals.Kitten

func main() {
	kittens = []animals.Kitten{
		animals.Kitten{
			Name: "Tigg",
			Hobbies: []string{
				"Playing with wool",
				"Eating",
			},
		},
	}

	http.HandleFunc("/list", ListKittens)

	http.ListenAndServe(":9000", http.DefaultServeMux)

}

func ListKittens(rw http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(kittens)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Write(data)
}
