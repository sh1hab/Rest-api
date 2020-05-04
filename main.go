package main

import (
	"database/sql"
		"fmt"
		"net/http"
		"log"
	
		_"github.com/go-sql-driver/mysql"
		"github.com/gorilla/mux"		
)

// Contact is 
type Contact struct{
	Email string `json:"Email"`
	Password string `json:"Password"`
	Status string `json:"status"`
}

// Contacts is
type allcontacts []Contact

var db *sql.DB
var err error

func main(){
	fmt.Println("Program is started...")
	connectToDb()
	handleRequests()
}

func handleRequests(){
	fmt.Println("Server is listening to requests in 127.0.0.1:8080 ")

	router:=mux.NewRouter()
	router.Headers("Content-type","Application/json")
	router.HandleFunc("/",homePage).Methods("GET").Name("home")
	router.HandleFunc("/api/contacts",allContacts).Methods("GET").Name("contacts")
	router.HandleFunc("/api/contacts/create",create).Methods("POST").Name("create_contact")	
	// router.HandleFunc("/api/contacts/update",create).Methods("PUT").Name("update_contact")	
	// router.HandleFunc("/api/contacts/delete",create).Methods("DELETE").Name("delete_contact")	

	log.Fatal( http.ListenAndServe(":8080",router) )
}
