package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Dew struct {
	name    string `json:"id,omitempty"`
	age     int    `json:"id,omitempty"`
	address string `json:"id,omitempty"`
}

var dews []Dew

func GetDew(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dews)
}

func CreateDew(w http.ResponseWriter, r *http.Request) {
	var dew Dew
	_ = json.NewDecoder(r.Body).Decode(&dew)
	dews = append(dews, dew)
	json.NewEncoder(w).Encode(dews)

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/dews", GetDew).Methods("GET")
	router.HandleFunc("/dew", CreateDew).Methods("POST")
	log.Println("Listening on Port: 7070")
	log.Fatal(http.ListenAndServe(":7070", router))
}
