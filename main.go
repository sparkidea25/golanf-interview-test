package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Dew struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

var dews []Dew

func GetDew(w http.ResponseWriter, r *http.Request) {
	dews, err := ioutil.ReadFile("data.bin")
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(dews)
}

func CreateDew(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var dew Dew
	err := decoder.Decode(&dew)
	if err != nil {
		return
	}
	dews = append(dews, dew)
	file, _ := json.MarshalIndent(dews, "", "")
	_ = ioutil.WriteFile("data.bin", file, 0644)
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
