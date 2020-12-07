package main

import (
	"encoding/json"
//	"fmt"
//	"net/http"
//	"net/smtp"
	"database/sql"	
//	"github.com/mattn/go-sqlite3" // Import go-sqlite3 library	
	_ "modernc.org/sqlite" // DO NOT UNCOMMENT
	"log"
	"os"
//	"strconv"
	"github.com/twinj/uuid"
	"io/ioutil"	
)

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
		"messClient" TEXT,
		"uuid" TEXT,
		PRIMARY KEY ("nom","mail")			
		);` // SQL Statement for Creating a clients table (if not existing)

	log.Println("Création ou ouverture d'une table des clients...")
	statement, err := db.Prepare(createClientTableSQL) // Prepare my SQL Statement
	if err != nil {	log.Panic(err) }
	defer statement.Close()				
	_, err = statement.Exec() // Execute my SQL Statement
	if err != nil { log.Panic(err) }	
	log.Println("Table des clients créée ou ouverte...")
}

func insertClient(db *sql.DB, newClient *receivedFromMTTchassis) {
	var unique_id string = ""
	var existingRecord bool = false
	log.Println("Ajout d'un nouveau client...")

	var testClientSQL string = `SELECT uuid FROM clients WHERE nom = ? AND mail = ?`
	// 1) Test if a record with this primary key is already present in the database
	statement, err := db.Prepare(testClientSQL) // Prepare statement.
												// should avoid SQL injections
	if err != nil {	log.Panic(err) }
	defer statement.Close()															
	err = statement.QueryRow(newClient.Nom, newClient.Mail).Scan(&unique_id)
	// TO DO : I'm not very happy with the test below : it relies on an error string which may change if the database driver changes : ==> improve that
	if err!=nil && err.Error() == MTT_NO_ROWS_IN_RESULT_SET {
		// There is no record for this client : so it is a new client, then we create its UUID
		unique_id = uuid.NewV4().String()
		err = nil
		existingRecord = false
	}  else { 
		// there is already a record for this client, and its UUID is unique_id
		existingRecord = true
	}

	if existingRecord { // It is already in the database, so just update
		// 2) Update the existing record
		log.Println("Client déjà existant...")
		log.Println("Mise à jour du client existant...")		
		var insertClientSQL string = `UPDATE clients SET prenom = ?, telephone = ?, messClient = ? WHERE nom = ? AND mail = ?` // important : PRIMARY KEY = ("nom","mail")
		statement, err = db.Prepare(insertClientSQL) // Prepare statement.
		// should avoid SQL injections
		if err != nil { log.Panic(err) }
		_, err = statement.Exec(newClient.Prenom, newClient.Telephone, newClient.MessClient, newClient.Nom, newClient.Mail) // proper code should be (*newClient).Prenom, (*newClient).Telephone ...
		if err != nil { log.Panic(err) }
		log.Println("Client existant mis à jour...")
		// DONE: update of the client record
	} else {
		// 3) Insert the record if not already present in the database
		// Here I'll add a UUID after the mail field, to uniquely identify the potential customer in the table <Interesting_Products>
		// And something like Corresp[i] will give an absolute product number corresponding to the local product number being i.
		// FOR THE MOMENT I assume that this relative product number is an absolute product number
		// The Corresp array will be filled in either using a json file or a database table

		var insertClientSQL string = `INSERT INTO clients(nom, prenom, telephone, mail, messClient, uuid) VALUES (?, ?, ?, ?, ?, ?)`
		// N.B. It would have liked to perform a WHERE NOT EXISTS (SELECT * FROM clients WHERE nom = ? AND mail = ? )
		// But (after many trials) it seems it does not work with sqlite (and/or GOLANG ?). Never mind, to make it work I've done the steps 1) and 2) just above
		statement, err = db.Prepare(insertClientSQL) // Prepare statement.
														// should avoid SQL injections
		if err != nil { panic(err) }
		_, err = statement.Exec(newClient.Nom, newClient.Prenom, newClient.Telephone, newClient.Mail, newClient.MessClient, unique_id) //, newClient.Nom, newClient.Mail) // (*newClient).Nom, (*newClient).Prenom, (*newClient).Telephone, (*newClient).Mail, (*newClient).Nom, (*newClient).Mail)
		if err != nil { log.Panic(err) }
		log.Println("Nouveau client ajouté...")
	}
	createOrUpdateInterestingProducts(db, unique_id, newClient.Produits) // update with new products that the client is interested in
}

func createOrUpdateInterestingProducts(db *sql.DB, unique_id string, products [] bool) {
	createInterestingProductsTable(db)
	UpdateInterestingProducts(db, unique_id, products)
	log.Println("-->", unique_id, "<--")	
	log.Println("-->", products , "<--")
}

func createInterestingProductsTable(db *sql.DB) {
	var createInterestingProductsTableSQL string = `CREATE TABLE IF NOT EXISTS InterestingProducts (
		"uuid" TEXT,
		"productNum" SMALLINT
		);` // SQL Statement to create a table of interesting products (if not existing)
	// NB : SMALLINT can go up to 32767 : far enough
	log.Println("Création ou ouverture d'une table des produits intéressants...")
	statement, err := db.Prepare(createInterestingProductsTableSQL) // Prepare my SQL Statement
	if err != nil {	log.Panic(err) }
	defer statement.Close()				
	_, err = statement.Exec() // Execute my SQL Statement
	if err != nil { log.Panic(err) }		
	log.Println("Table des produits intéressants créée...")
}

func UpdateInterestingProducts(db *sql.DB, unique_id string, produits [] bool) {
	// 1) Remove old records
	var deleteOldInterestingProductsSQL string = `DELETE from InterestingProducts WHERE uuid = ?` // important : PRIMARY KEY = uuid
	statement, err := db.Prepare(deleteOldInterestingProductsSQL) // Prepare statement.
	// should avoid SQL injections
	if err != nil { log.Panic(err) }
	defer statement.Close()	
	_, err = statement.Exec(unique_id)
	if err != nil { log.Panic(err) }
	// 2) Add new records
	var updateInterestingProductsSQL string = `INSERT INTO InterestingProducts(uuid, productNum) VALUES (?, ?)`
	for i:=0;i<len(produits);i++ {
		if (produits[i]) {
			statement, err = db.Prepare(updateInterestingProductsSQL) // Prepare statement.					
			if err != nil { panic(err) }			
			_, err = statement.Exec(unique_id, i+1) // First product starts at 1, not 0
			if err != nil { panic(err) }			
			// Hummm, here, later there will be a correspondance table or array to point from relative product number to absolute product number.
			// Typically, something like Corresp[i] will give a number which will be the absolute product number corresponding to the local product number being i.
			// FOR THE MOMENT I assume that this relative product number is an absolute product number
			// The Corresp array will be filled in either using a json file or a database table
		}
	}
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
		var produits [] bool
		row.Scan(&nom, &prenom, &telephone, &mail, &produits)
		log.Println("Client:", nom, prenom, telephone, mail, produits)
	}
}

func createProductsTableFromJson(myjson string) {
	// First : read the json content and fill in a structure
	type Product struct {
		ProductID			string	`json:"productID"`
		ProductName			string	`json:"productName"`
		ProductDescription	string	`json:"productDescription"`
		ProductLink			string	`json:"productLink"`
		ProductPrice		int 	`json:"productPrice"`
		ProductDelay		string	`json:"productDelay"`
		ProductActive		bool	`json:"productActive`
		ProductDateAdded	string	`json:"productDateAdded"`
	}

	type Products struct {
		JsonName		string `json:"jsonName"`
		Version			string `json:"version"`
		CreationDate	string `json:"creationData"`
		LastUpdate		string `json:"lastUpdate"`
		LastUpdater		string `json:"lastUpdater"`
		Products []Product     `json:"products"`
	}

	// Open our jsonFile
	jsonFile, err := os.Open(myjson)
	// if we os.Open returns an error then handle it
	if err != nil { panic(err) }
	// defer the closing of our jsonFile so that we can parse
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var products Products

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'products' which we defined above
	json.Unmarshal(byteValue, &products)

	// 2) Write the structure content inside a table of my database ? (not mandatory)
	// return // because maybe the json is enough to handle all that
	// 2.1) Create database if not existing
	if (!nodeExists(MTT_DATABASE)) {
		file, err := os.Create(MTT_DATABASE)
		if err != nil {	log.Panic(err) }
		file.Close()
	}
	sqliteDatabase, err := sql.Open("sqlite", MTT_DATABASE) // Open my SQL base
	if err != nil {	panic(err) }
	defer sqliteDatabase.Close()
	// 2.2) create products table if not existing
	var createProductsTableSQL string = `CREATE TABLE IF NOT EXISTS Products (
		"productID" TEXT,
		"productName" TEXT,
		"productDescription" TEXT,
		"productLink" TEXT,
		"productPrice" SMALLINT UNSIGNED,
		"productDelay" TEXT,
		"productActive" BIT,
		"productDateAdded" TEXT
		);` // SQL Statement to create a table of products (if not existing)
	// NB : SMALLINT UNSIGNED can go up to 65535 : used for a price
	// NB : BIT can be 0 (false) or 1 (true)
	statement, err := sqliteDatabase.Prepare(createProductsTableSQL) // Prepare my SQL Statement
	if err != nil {	panic(err) }
	defer statement.Close()				
	_, err = statement.Exec() // Execute my SQL Statement
	if err != nil { panic(err) }		
	// 2.3) Remove old records (if any)
	var deleteOldProductsSQL string = `DELETE from Products`
	statement, err = sqliteDatabase.Prepare(deleteOldProductsSQL) // Prepare statement.
	if err != nil { log.Panic(err) }
	_, err = statement.Exec()
	if err != nil { log.Panic(err) }
	// 2.4) Now insert the products in the database
	var insertProductsSQL string = `INSERT INTO Products(
		productID,
		productName,
		productDescription,
		productLink,
		productPrice,
		productDelay,
		productActive,
		productDateAdded)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	statement, err = sqliteDatabase.Prepare(insertProductsSQL) // Prepare statement.
	if err != nil { log.Panic(err) }
	var b uint8
	for i:=0; i<len(products.Products); i++ {
		if (products.Products[i].ProductActive) {b=1} else {b=0}
		_, err = statement.Exec(products.Products[i].ProductID,
								products.Products[i].ProductName,
								products.Products[i].ProductDescription,
								products.Products[i].ProductLink,
								products.Products[i].ProductPrice,
								products.Products[i].ProductDelay,
								b, // products.Products[i].ProductActive
								products.Products[i].ProductDateAdded)
		if err != nil { log.Panic(err) }
	}
}
