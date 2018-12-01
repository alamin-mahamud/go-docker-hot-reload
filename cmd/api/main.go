package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola from the api!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("listening on 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
