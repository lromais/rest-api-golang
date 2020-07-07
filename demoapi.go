package demoapi

import (
	"fmt"
	"net/http"
	"strconv"
	"log"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/hello/{fullName}", Hello).Methods("GET")
	a.Router.HandleFunc("/sum/{a}/{b}", Sum).Methods("GET")
}


func Hello(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fullName := vars["fullName"]
	fmt.Fprint(response, "Ola "+fullName)
}

func Sum(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	a := vars["a"]
	b := vars["b"]
	number1, _ := strconv.ParseInt(a, 10, 64)
	number2, _ := strconv.ParseInt(b, 10, 64)
	fmt.Fprint(response, number1+number2)
}
