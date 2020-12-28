package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
)

type Dew struct {
	name    string `json:"name,omitempty"`
	age     int    `json:"age,omitempty"`
	address string `json:"address,omitempty"`
}

var dews []Dew

func GetDew(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dews)
}

func CreateDew(w http.ResponseWriter, r *http.Request) {
	var dew Dew
	//json.NewDecoder(r.Body).Decode(&dew)
	json.Unmarshal(r.Body,&dew)
	fmt.Println(dew)
	dews = append(dews, dew)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dew)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/dews", GetDew).Methods("GET")
	router.HandleFunc("/dew", CreateDew).Methods("POST")
	log.Println("Listening on Port: 7070")
	log.Fatal(http.ListenAndServe(":7070", router))
}
