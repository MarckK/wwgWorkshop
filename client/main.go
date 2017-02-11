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

type Pets []Pet

func (a Pets) Len() int {
	return len(a)
}

func (a Pets) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Pets) Less(i, j int) bool {
	return a[i].Name > a[j].Name
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

	sort.Sort(pets)

	for _, pet := range pets {
		fmt.Printf("Name: %v Hobbies %v\n", pet.Name, pet.Hobbies)
	}
}
