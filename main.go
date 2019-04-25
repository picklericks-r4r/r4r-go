package main

import (
	"fmt"
	"net/http"

	"./controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/health",
		controllers.Health).Methods("GET")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Print(err)
	}
}
