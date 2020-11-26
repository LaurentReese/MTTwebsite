package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*type struct_received struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}*/

type struct_received struct {
	nom string `json:"nom"`
	prenom string `json:"prenom"`
	telephone string `json:"telephone"`
	mail string `json:"mail"`
}

type numsResponseData struct {
	Add float64 `json:"add"`
	Mul float64 `json:"mul"`
	Sub float64 `json:"sub"`
	Div float64 `json:"div"`
}

func process(numsdata struct_received) (numsResponseData) {
	
	var numsres numsResponseData
	/*numsres.Add = numsdata.Num1 + numsdata.Num2
	numsres.Mul = numsdata.Num1 * numsdata.Num2
	numsres.Sub = numsdata.Num1 - numsdata.Num2
	numsres.Div = numsdata.Num1 / numsdata.Num2*/

	return numsres
}

func mttChassis(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var numsData struct_received
	var numsResData numsResponseData
	
	decoder.Decode(&numsData)
	fmt.Println(numsData)
	/*
	fmt.Println("======================")		
	fmt.Println("==>",numsData.nom,"<==")	
	fmt.Println("==>",numsData.prenom,"<==")		
	fmt.Println("==>",numsData.telephone,"<==")			
	fmt.Println("==>",numsData.mail,"<==")			
	fmt.Println("===================")*/

	//numsResData = process(numsData)

	// fmt.Println(numsResData)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(numsResData); err != nil {
        panic(err)
    }
}

func main() {
	http.HandleFunc("/", mttChassis)
	http.ListenAndServe(":8090", nil)
}