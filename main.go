package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", HelloWorld)
	http.HandleFunc("/goodbye", GoodbyeWorld)
	http.ListenAndServe(":9000", http.DefaultServeMux)
}

func HelloWorld(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello")
}

func GoodbyeWorld(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(rw, "Goodbye ", name)
}
