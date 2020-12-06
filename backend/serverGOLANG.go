package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"database/sql"	
//	"github.com/mattn/go-sqlite3" // Import go-sqlite3 library	
	_ "modernc.org/sqlite"
	"log"
	"os"
	"strconv"
)

const MTT_DATABASE string ="MTT-sqlite-database.db"
const MTT_ACKNOWLEDGE string = "L'entreprise MTT a été informée, merci de votre intérêt"
const MTT_MAX_PRODUCTS int = 7
	

/* TO DO : improve error handling later on with this kind of treatment
// To put inside a function
if err := dec.Decode(&val); err != nil {
    if serr, ok := err.(*json.SyntaxError); ok {
        line, col := findLine(f, serr.Offset)
        return fmt.Errorf("%s:%d:%d: %v", f.Name(), line, col, err)
    }
    return err
}*/

type receivedFromMTTchassis struct {
	Nom string `json:"nom"`
	Prenom string `json:"prenom"`
	Telephone string `json:"telephone"`
	Mail string `json:"mail"`
	Produits [MTT_MAX_PRODUCTS] bool `json:"produits"` // TO DO : a slice instead of an array ?
	MessClient string `json:"messClient"`
}

// data coming from my vuejs client
//var data = {"nom" : this.nom, "prenom" : this.prenom, "telephone" : this.telephone, "mail" : this.mail}

type responseFromGOserver struct {
	Message string `json:"messageServer"`
}

func mttChassis(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body) // create json decoder ...
	var mttData receivedFromMTTchassis
	var reponseData responseFromGOserver
	
	decoder.Decode(&mttData) // ... and receive data from the vuejs client

	//fmt.Println(mttData)  KEEP it for debugging purpose

	sendMail(&mttData) // this struct may become bigger, so better to pass it by address
	newClientInDatabase(&mttData)

	reponseData.Message = MTT_ACKNOWLEDGE

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(reponseData); err != nil { panic(err) }
}

// +----------------------------------------------+
// | DATABASE BEGIN DATABASE BEGIN DATABASE BEGIN |
// +----------------------------------------------+

func nodeExists(node string) bool { // to me a node is a folder or a filepath
	_ , err := os.Stat(node)
	if err != nil { return false }
	if os.IsNotExist(err) {return false}
	return true;
}

func newClientInDatabase(newClient *receivedFromMTTchassis) {
	// os.Remove(MTT_DATABASE)	// Test purpose
	if (!nodeExists(MTT_DATABASE)) {
		log.Println("Création de la base de données...", MTT_DATABASE)
		file, err := os.Create(MTT_DATABASE)
		if err != nil {	log.Panic(err) }
		file.Close()
		log.Println("Base de données", MTT_DATABASE, "créée")
	}
	sqliteDatabase, err := sql.Open("sqlite", MTT_DATABASE) // Open my SQL base
	if err != nil {	log.Panic(err) }
	defer sqliteDatabase.Close()
	log.Println("Base de données", MTT_DATABASE, "ouverte")	

	createClientTable(sqliteDatabase) // will create table if it does not exist
	insertClient(sqliteDatabase, newClient) // will add or update client
}

func createClientTable(db *sql.DB) {
	/*	syntax for a primary key on 2 fields is the following :
	CREATE TABLE something (
			column1, 
			column2, 
			column3, 
			PRIMARY KEY (column1, column2)
		  );
	for one field with autoincrement, it is like : "integer_field" integer NOT NULL PRIMARY KEY AUTOINCREMENT ==> for a primary key based on a single field
		   */
	var createClientTableSQL string = `CREATE TABLE IF NOT EXISTS clients (
		"nom" TEXT,
		"prenom" TEXT,
		"telephone" TEXT,
		"mail" TEXT,
		PRIMARY KEY ("nom","mail")			
		);` // SQL Statement for Creating a clients table (if not existing)

	log.Println("Création ou ouverture d'une table des clients...")
	statement, err := db.Prepare(createClientTableSQL) // Prepare my SQL Statement
	if err != nil {	log.Panic(err) }
	defer statement.Close()				
	statement.Exec() // Execute my SQL Statement
	log.Println("Table des clients créée ou ouverte...")
	}

func insertClient(db *sql.DB, newClient *receivedFromMTTchassis) { // nom string, prenom string, telephone string, mail string) {
	var count int
	log.Println("Ajout d'un nouveau client...")

	var testClientSQL string = `SELECT COUNT(*) FROM clients WHERE nom = ? AND mail = ?`
	// 1) Test if record with this primary key is already present in the database
	statement, err := db.Prepare(testClientSQL) // Prepare statement.
												// should avoid SQL injections
	if err != nil {	log.Panic(err) }
	defer statement.Close()															
	err = statement.QueryRow(newClient.Nom, newClient.Mail).Scan(&count)
	if err != nil {	log.Panic(err) }

	if count == 1 { // It is already in the database, so just update
		// 2) Update the existing record
		log.Println("Client déjà existant...")
		log.Println("Mise à jour du client existant...")		
		var insertClientSQL string = `UPDATE clients SET prenom = ?, telephone = ? WHERE nom = ? AND mail = ?` // important : PRIMARY KEY = ("nom","mail")
		statement, err = db.Prepare(insertClientSQL) // Prepare statement.
		// should avoid SQL injections
		if err != nil { log.Panic(err) }
		_, err = statement.Exec(newClient.Prenom, newClient.Telephone, newClient.Nom, newClient.Mail) // proper code should be (*newClient).Prenom, (*newClient).Telephone ...
		if err != nil { log.Panic(err) }
		log.Println("Client existant mis à jour...")
		return // DONE: update of the client record
		// TO DO : update with new products that the client is interested in
	}

	// 3) Insert the record if not already present in the database
	// Here I'll add a UUID after the mail field, to uniquely identify the potential customer in the table <Interesting_Products>
	// And something like Corresp[i] will give an absolute product number corresponding to the local product number being i.
	// FOR THE MOMENT I assume that this relative product number is an absolute product number
	// The Corresp array will be filled in either using a json file or a database table

	var insertClientSQL string = `INSERT INTO clients(nom, prenom, telephone, mail) VALUES (?, ?, ?, ?)`
	// N.B. It would have liked to perform a WHERE NOT EXISTS (SELECT * FROM clients WHERE nom = ? AND mail = ? )
	// But (after many trials) it seems it does not work with sqlite (and/or GOLANG ?). Never mind, to make it work I've done the steps 1) and 2) just above
	statement, err = db.Prepare(insertClientSQL) // Prepare statement.
													// should avoid SQL injections
	if err != nil { panic(err) }
	_, err = statement.Exec(newClient.Nom, newClient.Prenom, newClient.Telephone, newClient.Mail) //, newClient.Nom, newClient.Mail) // (*newClient).Nom, (*newClient).Prenom, (*newClient).Telephone, (*newClient).Mail, (*newClient).Nom, (*newClient).Mail)
	if err != nil { log.Panic(err) }
	log.Println("Nouveau client ajouté...")
}
	
// TO DO : to get the clients (from the vuejs side), improve the following function
func displayClients(db *sql.DB) {
	row, err := db.Query("SELECT * FROM clients ORDER BY nom")
	if err != nil {	log.Panic(err) }
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var nom string
		var prenom string
		var telephone string
		var mail string	
		var produits [MTT_MAX_PRODUCTS] int
		row.Scan(&nom, &prenom, &telephone, &mail, &produits)
		log.Println("Client:", nom, prenom, telephone, mail, produits)
	}
}

// +----------------------------------------+
// | DATABASE END DATABASE END DATABASE END |
// +----------------------------------------+

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

	for i:=0;i<MTT_MAX_PRODUCTS;i++ {
		if (info.Produits[i]) {		
			// Hummm, here, later there will be a correspondance table or array to point from relative product number to absolute product number.
			// Typically, something like Corresp[i] will give a number which will be the absolute product number corresponding to the local product number being i.
			// FOR THE MOMENT I assume that this relative product number is an absolute product number
			// The Corresp array will be filled in either using a json file or a database table
			messageString += strconv.Itoa(i+1) // + 1 because product number doesn't start at 0
			messageString += "\r\n"
		}
	}

	if info.MessClient != "" { messageString += "Message Client :" + "\r\n" + info.MessClient }
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


func main() {
	http.HandleFunc("/", mttChassis)
	http.ListenAndServe(":8090", nil)
}