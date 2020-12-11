package main

import (
//	"encoding/json"
	"fmt"
//	"net/http"
	"net/smtp"
	"strconv"
	"github.com/jordan-wright/email"
)

func sendMail(info *receivedFromMTTchassis) {
	// Sender data.
	from := "rene.lasurete@gmail.com" // old useless account that I don't care about
	password := "d<5@M48c6UyDz]" // password is here in clear but I don't care as it's an old useless account

	// Receiver email address.
	to := []string{	"rene.lasurete@gmail.com" }

	// mail from myself to myself, just to create an entry to store a customer action ;-)

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// TO DO ? : add date and time inside the message, although it is normally already in the mail itself
	messageString := "Message reçu de la part de " + info.Prenom + " " + info.Nom + "\r\n" +
					"mail:" + info.Mail + "\r\n" +
					"téléphone:" + info.Telephone + "\r\n" +
					"==> intéressé(e) par les produits :" + "\r\n"
	for i:=0;i<len(info.Produits);i++ {
		if (info.Produits[i]) {		
			// Hummm, here, later there will be a correspondance table or array to point from relative product number to absolute product number.
			// Typically, something like Corresp[i] will give a number which will be the absolute product number corresponding to the local product number being i.
			// FOR THE MOMENT I assume that this relative product number is an absolute product number
			// The Corresp array will be filled in either using a json file or a database table
			messageString += strconv.Itoa(i+1) // + 1 because product number doesn't start at 0
			messageString += "\r\n"
		}
	}

	messageString += "\r\n"
	if info.AddrTravaux != "" { messageString += "Adresse travaux :" + "\r\n" + info.AddrTravaux + "\r\n\r\n"}
	
	if info.MessClient != "" { messageString += "Message Client :" + "\r\n" + info.MessClient + "\r\n"}
	//fmt.Println(messageString)
	message := [] byte (messageString)	// fmt.Println(message) would give only numbers

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

func sendMailDatabase(dbName string) {
	// See https://github.com/jordan-wright/email/blob/master/README.md			
	e := email.NewEmail()
	e.From = "rene.lasurete@gmail.com"
	password := "d<5@M48c6UyDz]" // password is here in clear but I don't care as it's an old useless account	
	e.To = []string{"rene.lasurete@gmail.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "MTT Database"
	e.Text = []byte("Veuillez trouver la database MTT en pièce attachée.")
	//e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	e.AttachFile(dbName)		
	err := e.Send(  "smtp.gmail.com:587",
			 smtp.PlainAuth("",
			 e.From,
			 password,
			 "smtp.gmail.com"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email généré contenant la database", dbName)
	e = nil
}

