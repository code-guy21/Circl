package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome to my API")
}

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")

	log.Println("API is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}