package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"	
//	"github.com/mattn/go-sqlite3" // Import go-sqlite3 library	
	_ "modernc.org/sqlite"
	"log"
	"os"
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

func fillDatabase() {
	os.Remove("sqlite-database.db") // I delete the file to avoid duplicated records. 
                                    // SQLite is a file based database.

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close() // Defer Closing the database
	createTable(sqliteDatabase) // Create Database Tables

        // INSERT RECORDS
	insertStudent(sqliteDatabase, "0001", "Liana Kim", "Bachelor")
	insertStudent(sqliteDatabase, "0002", "Glen Rangel", "Bachelor")
	insertStudent(sqliteDatabase, "0003", "Martin Martins", "Master")
	insertStudent(sqliteDatabase, "0004", "Alayna Armitage", "PHD")
	insertStudent(sqliteDatabase, "0005", "Marni Benson", "Bachelor")
	insertStudent(sqliteDatabase, "0006", "Derrick Griffiths", "Master")
	insertStudent(sqliteDatabase, "0007", "Leigh Daly", "Bachelor")
	insertStudent(sqliteDatabase, "0008", "Marni Benson", "PHD")
	insertStudent(sqliteDatabase, "0009", "Klay Correa", "Bachelor")

        // DISPLAY INSERTED RECORDS
	displayStudents(sqliteDatabase)
}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE student (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT,
		"name" TEXT,
		"program" TEXT		
	  );` // SQL Statement for Create Table

	log.Println("Create student table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("student table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertStudent(db *sql.DB, code string, name string, program string) {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO student(code, name, program) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertStudentSQL) // Prepare statement. 
                                                   // This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(code, name, program)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayStudents(db *sql.DB) {
	row, err := db.Query("SELECT * FROM student ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var code string
		var name string
		var program string
		row.Scan(&id, &code, &name, &program)
		log.Println("Student: ", code, " ", name, " ", program)
	}
}

func main() {
	fillDatabase()
	return	
	http.HandleFunc("/", mttChassis)
	http.ListenAndServe(":8090", nil)
}