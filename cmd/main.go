package main

import (
	demoapi "github.com/lromais/rest-api-golang"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/demo/hello/{fullName}", demoapi.Hello).Methods("GET")
	router.HandleFunc("/api/demo/sum/{a}/{b}", demoapi.Sum).Methods("GET")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}
}
