package main

import (
	"fmt"
	"net/http"
	"testing_go/controller"

	"github.com/gorilla/mux"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/pokemon/{id}", controller.GetPokemon).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error found")
	}
}
