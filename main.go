package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

func PostDew(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to my Homme Route")
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No port founf " + port)
	}
	return ":" + port
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/home", PostDew)
	router.HandleFunc("/dews", GetDew).Methods("GET")
	router.HandleFunc("/dew", CreateDew).Methods("POST")
	log.Println("Listening on Port: 7070")
	http.ListenAndServe(GetPort(), router)
}
