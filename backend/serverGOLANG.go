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

func nodeExists(node string) bool { // to me a node is a folder or a filepath
	_ , err := os.Stat(node)
	if err != nil { return false }
	if os.IsNotExist(err) {return false}
	return true;
}


func main() {
	//createProductsTableFromJson(MTT_JSON_NAME)
	//return
	http.HandleFunc("/", mttChassis)
	http.ListenAndServe(":8090", nil)
}