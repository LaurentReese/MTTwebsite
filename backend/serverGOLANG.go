package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
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
	sendMail(mttData)
	// TO DO : + later => fill in a database with the record "mttData"
	// Sqlite

	reponseData.Message = response	

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(reponseData); err != nil {	
        panic(err)
    }
}

func sendMail(info receivedFromMTTchassis) {
 // Sender data.
 from := "rene.lasurete@gmail.com" // old useless acount that I don't care about
 password := "d<5@M48c6UyDz]" // password is here in clear but I don't care as it's an old useless acount

 // Receiver email address.
 to := []string{
   "rene.lasurete@gmail.com",
 }

 // mail from myself to myself, just to create an entry to store a customer action ;-)

 // smtp server configuration.
 smtpHost := "smtp.gmail.com"
 smtpPort := "587"

 // Message.
 //message := []byte("Test d'un message sur le MTT serveur.")
 // TO DO ? : add date and time inside the message, although it is normally already in the mail itself
 message := []byte ("Message reçu de la part de " + info.Prenom + " " + info.Nom + "\r\n" +
                    "mail:" + info.Mail + "\r\n" +
                    "téléphone:" + info.Telephone + "\r\n" +
                    "==> intéressé par les chassis")
 
 // Authentication.
 auth := smtp.PlainAuth("", from, password, smtpHost)
 
 // Sending email.
 err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
 if err != nil {
   fmt.Println(err)
   return
 }
 fmt.Println("Email généré")
}


func main() {
//	var test receivedFromMTTchassis
//	sendMail(test)
//	return
	http.HandleFunc("/", mttChassis)
	http.ListenAndServe(":8090", nil)
}