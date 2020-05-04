package main

import (
	"database/sql"
		"fmt"
		"net/http"
		"log"
		"encoding/json"
		_"github.com/go-sql-driver/mysql"		
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
	connectToDb()
	handleRequests()
}

func connectToDb(){
	fmt.Println("Connecting to mysql")

	db,err	:=	sql.Open("mysql","root@tcp(127.0.0.1:3306)/Accounts")

	if  err !=nil{
		panic( err.Error() )
	}

	defer db.Close()

}

func handleRequests(){
	fmt.Println("Server is listening to requests in 127.0.0.1:8080 ")
	http.HandleFunc("/",homePage)
	http.HandleFunc("/api/contacts",allContacts)
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

func insert(){
	insert, err	:=	db.Query("insert into table1 (domain,email_or_phone,password,status) 
						values ("facebook.com","01","1234",1)")

	if err !=nil {
		panic(err.Error())
	}

	defer insert.Close()

}