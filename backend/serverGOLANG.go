package main

import (
	"encoding/json"
//	"fmt"
	"net/http"
//	"net/smtp"
//	"database/sql"	
//	"github.com/mattn/go-sqlite3" // Import go-sqlite3 library	
//	_ "modernc.org/sqlite"
//	"log"
	"os"
//	"strconv"
//	"github.com/twinj/uuid"
//	"io/ioutil"	
)

const MTT_DATABASE string ="MTT-sqlite-database.db"
const MTT_ACKNOWLEDGE string = "L'entreprise MTT a été informée, merci de votre intérêt"
const MTT_NO_ROWS_IN_RESULT_SET string = "sql: no rows in result set"
const MTT_JSON_NAME string = "MTTchassis.json"
	
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
	Produits [] bool `json:"produits"` // a slice instead of an array !
	AddrTravaux string `json:"addrTravaux"`		
	MessClient string `json:"messClient"`
	DescProjet string `json:"descProjet"`		
	DateProjet string `json:"dateProjet"`
	Comment string `json:"comment"`
}


type receivedFromMTTchassisPassword struct {
	Password string `json:"password"`
}

type receivedFromMTTJson struct {
	Password string `json:"password"`
	Text string `json:"text"`
}

// data coming from my vuejs client
//var data = {"nom" : this.nom, "prenom" : this.prenom, "telephone" : this.telephone, "mail" : this.mail}

type responseFromGOserver struct {
	MessageServer string `json:"messageServer"`
}

// TO DO : there is a lot of code to factorize in mttChassis, mttDatabaseAction, mttJsonAction, ...
// But today I don't know which one I will keep, which one I will withdraw, etc
// So keep this work for later...

func checkPassword(pass string) (bool, string) {
	// Sometimes password string be found with a grep inside the binary
	// So I want to make it more difficult to find...
	// TO DO : of course change this password before last delivery and do not commit it !
	// TO DO : decrypt the password here, if it has been encrypted on the vuejs side
	if pass[0:1]=="L" && pass[1:2]=="a" && pass[2:3]=="u" && pass[3:4]=="r" && pass[4:5]=="e" && pass[5:6]=="n" && pass[6:7]=="t" {
		return true, "Mot de passe vérifié"
	}
	return false, "Mot de passe incorrect"
}

func mttChassis(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body) // create json decoder ...
	var mttData receivedFromMTTchassis
	var reponseData responseFromGOserver
	
	decoder.Decode(&mttData) // ... and receive data from the vuejs client
	if mttData.Comment != "" { return }
	//fmt.Println(mttData)  KEEP it for debugging purpose

	go sendMail(&mttData) // this struct may become bigger, so better to pass it by address.
	// And also : as sending the mail may take some time, do it in a separated thread...
	newClientInDatabase(&mttData) // ... while the database treatment (basically faster) is called in the current thread
	// N.B. no concurrency between mail sending and database updating, so no need to sync or whatever

	reponseData.MessageServer = MTT_ACKNOWLEDGE

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(reponseData); err != nil { panic(err) }
}

func mttDatabaseAction(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body) // create json decoder ...
	var mttDataPassword receivedFromMTTchassisPassword
	var reponseData responseFromGOserver
	
	decoder.Decode(&mttDataPassword) // ... and receive data from the vuejs client
	//fmt.Println(mttDataPassword.Password) // KEEP it for debugging purpose

	var verif bool
	verif, reponseData.MessageServer = checkPassword(mttDataPassword.Password)
	if (verif) {
		if !nodeExists(MTT_DATABASE) {
			// password is ok, but the database is not available
			reponseData.MessageServer = MTT_DATABASE + " n'existe pas sur le serveur !"
		} else {
			// password ok, database ok
			// Then send the database by mail : we could send it back to vuejs by http... but with vuejs it's forbidden to access the system file
			sendDatabaseByMail(MTT_DATABASE)
			reponseData.MessageServer = MTT_DATABASE + " vous a été envoyée par mail"
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(reponseData); err != nil { panic(err) }
}

func mttJsonAction(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body) // create json decoder ...
	var mttDataJson receivedFromMTTJson
	var	reponseData responseFromGOserver

	decoder.Decode(&mttDataJson) // ... and receive data from the vuejs client

	var verif bool
	verif, reponseData.MessageServer = checkPassword(mttDataJson.Password)
	if (verif) {
		jsonByteArray := [] byte(mttDataJson.Text)
		jsonOK := json.Valid(jsonByteArray)
		if !jsonOK {
			reponseData.MessageServer = "Erreur dans le fichier json transmis"
		} else {
			if createProductsTableFromJsonContent(jsonByteArray) {
				reponseData.MessageServer = "Succès : produits intégrés dans la base de données"
			} else {
				reponseData.MessageServer = "Erreur de syntaxe dans le fichier json transmis"			
			}
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(reponseData); err != nil { panic(err) }
}

func nodeExists(node string) bool { // to me a node is a folder or a filepath
	_ , err := os.Stat(node)
	if err != nil { return false }
	if os.IsNotExist(err) {return false}
	return true;
}


func main() {
//	createProductsTableFromJson(MTT_JSON_NAME)
//	return
	http.HandleFunc("/mttJsonAction", mttJsonAction)
	http.HandleFunc("/mttChassis", mttChassis)
	http.HandleFunc("/mttDatabaseAction", mttDatabaseAction)
	if os.Getenv("MTT_EXECUTION")=="PRODUCTION" {
		http.ListenAndServe(":80", nil)		// production
	} else {
		http.ListenAndServe(":8090", nil)	// local
	}

}