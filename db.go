package main

import (
	"database/sql"
		"fmt"
		"net/http"
		"encoding/json"
		"io/ioutil"
		"log"
		"os"

		_"github.com/go-sql-driver/mysql"
		// "github.com/gorilla/mux"		
)

var contacts allcontacts
const dbname string = "Accounts"

func connectToDb(){
	fmt.Println("Connecting to mysql")

	db,err	=	sql.Open("mysql","root@tcp(127.0.0.1:3306)/"+dbname)

	if  err !=nil{
		panic( err.Error() )
	}

	defer db.Close()

}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Sprintln("<h3>This is the home page</h3>")
}

func allContacts(w http.ResponseWriter, r *http.Request){
	contacts = allcontacts{
		Contact{ Email:"abc@example.com",Password:"123",Status:"active" },
	} 
	json.NewEncoder(w).Encode( contacts )
}

func create(w http.ResponseWriter, r *http.Request){
	var contact Contact
	input, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal( err.Error() )
	}
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	// parse json  & puts data into contact var
	json.Unmarshal(input,&contact)
	contacts = append( contacts , contact )
	keyVal 	:= make(map[string]string)
	domain	:=	keyVal["domain"]
	emailorphone :=	keyVal["email_or_phone"]
	status 		:=	keyVal["status"]
	password 	:=keyVal["password"]


	stmt, err	:=	db.Prepare("insert into table1 (domain, email_or_phone, password, status) values (?,?,?,?)")
	if err != nil {
		panic( err.Error() )
	}

	_, err = stmt.Exec( domain,emailorphone,password,status )
  	if err != nil {
    	panic(err.Error())
	}
	  
	w.WriteHeader( http.StatusCreated )
	json.NewEncoder( w ).Encode( contact )
	log.SetOutput(file)
 

	defer stmt.Close()

	defer file.Close()


}