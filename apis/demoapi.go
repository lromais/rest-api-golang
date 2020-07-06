package demoapi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fullName := vars["fullName"]
	fmt.Fprint(response, "Hello "+fullName)
}

func Sum(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	a := vars["a"]
	b := vars["b"]
	number1, _ := strconv.ParseInt(a, 10, 64)
	number2, _ := strconv.ParseInt(b, 10, 64)
	fmt.Fprint(response, number1+number2)
}
