package main

import (
		"fmt"
		"net/http"
		"log"
		"encoding/json"		
)

// Contact is 
type Contact struct{
	Email string `json:"Email"`
	Password string `json:"Password"`
	Status string `json:"status"`
}

// Contacts is
type Contacts []Contact

func main(){
	fmt.Println("Program is started...")
	handleRequests()
}

func handleRequests(){
	fmt.Println("Server is listening to requests in 127.0.0.1:8080 ")
	http.HandleFunc("/",homePage)
	http.HandleFunc("/contacts",allContacts)
	log.Fatal(http.ListenAndServe(":8080",nil) )
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Sprintln("<h3>This is the home page</h3>")
}

func allContacts(w http.ResponseWriter, r *http.Request){
	contacts := Contacts{
		Contact{ Email:"abc@example.com",Password:"123",Status:"active" },
	} 
	json.NewEncoder(w).Encode( contacts )
}