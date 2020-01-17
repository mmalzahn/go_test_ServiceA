package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/v1").Subrouter()

	api.HandleFunc("/user/{userID}/bills/{billID}", getIDs).Methods(http.MethodGet)
	r.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(":8080", r))
}
func getIDs(w http.ResponseWriter, r *http.Request) {

	pathParameter := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(pathParameter)
	w.Write(json)

	fmt.Println(json)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "muxer neu"}`))

	fmt.Println("{\"message\": \"muxer neu\"}")
}
