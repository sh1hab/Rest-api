package main

import (
	"database/sql"
		"fmt"
		"net/http"
		"encoding/json"
		"io/ioutil"
		"log"
		"runtime/debug"

		_"github.com/go-sql-driver/mysql"
)

var contacts allcontacts
const dbname string = "Accounts"

func connectToDb(){
	fmt.Println("Connecting to mysql")

	db,err	=	sql.Open("mysql","root@tcp(127.0.0.1:3306)/Accounts")
	if r := recover(); r != nil {
		fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
	}

	// defer db.Close()

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
	// file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// parse json  & puts data into contact var
	contacts 		= 	append( contacts , contact )
	keyVal 			:=	make(map[string]string)
	json.Unmarshal( input, &keyVal )
	domain			:=	keyVal["domain"]
	emailorphone 	:=	keyVal["email_or_phone"]
	status 			:=	keyVal["status"]
	password 		:=	keyVal["password"]

	stmt, err	:=	db.Prepare("INSERT INTO table1 (domain, email_or_phone, password, status) VALUES (?,?,?,?)")
	if r := recover(); r != nil {
		fmt.Println( r )
		fmt.Println("Stacktrace from panic: \n" + string( debug.Stack() ) )
	}

	defer stmt.Close()

	_, err = stmt.Exec( domain,emailorphone,password,status )
	if r := recover(); r != nil {
		fmt.Println( r )
		fmt.Println("Stacktrace from panic: \n" + string( debug.Stack() ) )
	}

	w.WriteHeader( http.StatusCreated )
	json.NewEncoder( w ).Encode( contact )
	// log.SetOutput(file)
	// defer file.Close()

}