package main

import (
		"fmt"
		"net/http"
		"log"
		
)

type Contact struct{
	email string `json:"Email"`
	password string `json:Password`
	status string `json:status`
}

var Contacts []Contact

func main(){
	handleRequests()
}

func handleRequests(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/contacts",allContacts)
	log.Fatal(http.ListenAndServe(":8080",nil) )
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Sprintln("<h3>This is the home page</h3>")
}

func allContacts(w http.ResponseWriter, r *http.Request){
	contacts := Contacts{
		Contact{ "email":"abc@example.com","password":"123","status":"active" },
	} 
	json.NewEncoder(w).Encode( contacts )
}