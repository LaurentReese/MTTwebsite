package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type receivedFromMTTchassis struct {
	Nom string `json:"nom"`
	Prenom string `json:"prenom"`
	Telephone string `json:"telephone"`
	Mail string `json:"mail"`	
}

// data coming from my vuejs client
//var data = {"nom" : this.nom, "prenom" : this.prenom, "telephone" : this.telephone, "mail" : this.mail}

type responseFromGOserver struct {
	Message string `json:"messageServer"`
}

func mttChassis(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var mttData receivedFromMTTchassis

	var reponseData responseFromGOserver

	var response = "L'entreprise MTT a été informée, merci de votre intérêt"
	
	decoder.Decode(&mttData)

	fmt.Println(mttData)

	reponseData.Message = response	

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(reponseData); err != nil {	
        panic(err)
    }
}

func main() {
	http.HandleFunc("/", mttChassis)
	http.ListenAndServe(":8090", nil)
}