package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type kundenStammdaten struct {
	ID       string
	Vorname  string
	Nachname string
	Stasse   string
	Ort      string
	PLZ      int
}

/*
Hauptfunktion
*/

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/v1").Subrouter()

	api.HandleFunc("/user/{userID}", getIDs).Methods(http.MethodGet)
	r.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(":8080", r))
}

/*
gezielte Abfrage von Daten
*/

func getIDs(w http.ResponseWriter, r *http.Request) {

	pathParameter := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json, _ := json.Marshal(getData(pathParameter["userID"]))
	w.Write(json)

	fmt.Println(json)
}

/*
Defaultlistener
*/
func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "muxer neu"}`))

	fmt.Println("{\"message\": \"muxer neu\"}")
}

/*
Daten
*/
func getData(dataID string) kundenStammdaten {
	var userdata map[string]kundenStammdaten
	userdata = make(map[string]kundenStammdaten)
	userdata["matthias"] = kundenStammdaten{ID: "4325", Vorname: "Matthias", Nachname: "Malzahn", Stasse: "Musterstraße 1", Ort: "Irgendwo", PLZ: 12345}
	userdata["iris"] = kundenStammdaten{ID: "3245", Vorname: "Iris", Nachname: "Malzahn", Stasse: "Musterstraße 1", Ort: "Irgendwo", PLZ: 12345}
	userdata["hans"] = kundenStammdaten{ID: "46577", Vorname: "Hans", Nachname: "Mustermann", Stasse: "Mustergasse 42", Ort: "Irgendwoanders", PLZ: 54321}
	userdata["josef"] = kundenStammdaten{ID: "61233", Vorname: "Josef", Nachname: "Hinterhuber", Stasse: "Hauptstraße 20", Ort: "Musterstadt", PLZ: 73523}
	userdata["martin"] = kundenStammdaten{ID: "23332", Vorname: "Martin", Nachname: "Müller", Stasse: "Musterstraße 1", Ort: "berlin", PLZ: 29821}
	userdata["susanne"] = kundenStammdaten{ID: "76856", Vorname: "Susanne", Nachname: "Meier", Stasse: "Musterstraße 1", Ort: "Hamburg", PLZ: 38746}
	return userdata[dataID]
}
