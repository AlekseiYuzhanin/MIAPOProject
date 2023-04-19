package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:id`
	Isbn string `json:isbn`
	Title string `json:title`
	Director *Director `json:director`
}

type Director struct{
	Firstname string `json:firstname`
	Lastname string `json:lastname`
}

var movies []Movie

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/movie",getMovies).Methods("GET")
	r.HandleFunc("/movies{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Server started at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}
