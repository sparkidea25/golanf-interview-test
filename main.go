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

const(
	DB = "data.bin"
)

type Dew struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

var dews []Dew

func PostDew(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Welcome to my Homme Route")
}

func GetDew(w http.ResponseWriter, r *http.Request) {
	var dews []Dew
	dews = make([]Dew, 0)
	db, err := os.Open(DB) 
	if err != nil {
		http.Error(w, "An error was encountered", http.StatusInternalServerError)
        return
	}
	defer db.Close()

	dbInfo, err := db.Stat()

	if(dbInfo.Size() > 0){
		data, err := ioutil.ReadAll(db)
		if err != nil {
			http.Error(w, "An error was encountered", http.StatusInternalServerError)
        	return
		}
		err = json.Unmarshal(data, &dews)
		if err != nil {
			http.Error(w, "An error was encountered", http.StatusInternalServerError)
			return
		}
	}

    
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dews)
}

func CreateDew(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var dew Dew
	err := decoder.Decode(&dew)
	if err != nil {
		return
	}


	dews = append(dews,dew)

    data, err := json.Marshal(dews)
    if err != nil {
		http.Error(w, "An error was encountered", http.StatusInternalServerError)
        return
	}
	
	err = ioutil.WriteFile(DB, data, 0644)
    if err != nil {
		http.Error(w, "An error was encountered", http.StatusInternalServerError)
        return
	}
	
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dew)
}



func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "7070"
	}
	log.Println("Listening on Port: "+port)
	return ":" + port
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", PostDew)
	router.HandleFunc("/dews", GetDew).Methods("GET")
	router.HandleFunc("/dew", CreateDew).Methods("POST")
	http.ListenAndServe(GetPort(), router)
}
