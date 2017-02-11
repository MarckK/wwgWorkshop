package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type Pet struct {
	Name    string
	Hobbies []string
}

// var pets []Pet
// var pets Pets
//the above 2 lines mean same thing as below line

type Pets []Pet

func (a Pets) Len() int {
	return len(a)
}

func (a Pets) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Pets) Less(i, j int) bool {
	return a[i].Name > a[j].Name
	// return len(a[i].Hobbies) < len(a[j].Hobbies)
}

func main() {
	resp, err := http.Get("http://localhost:9000/list")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	pets := Pets{}
	err = json.Unmarshal(data, &pets)

	// log.Println(string(data))
	// log.Println(pets)

	sort.Sort(pets)

	for _, pet := range pets {
		// log.Println(pet)
		fmt.Printf("Name: %v Hobbies %v\n", pet.Name, pet.Hobbies)
	}
}
