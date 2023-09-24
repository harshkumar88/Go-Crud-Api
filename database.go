package main

import (    
    "fmt"
   "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

func databaseConnection(){
  
    db, err := sql.Open("mysql", "pikachu:clgIgqqFVKvFHvBw@tcp(10.86.144.3:3306)/testDB")
    if err != nil {
        fmt.Println("error validating sql.optn arguments")
        panic(err.Error())
    }

    dB = db

    err = db.Ping()
    if err != nil {
        fmt.Println("error verifying connection with db.Ping")
        panic(err.Error())
    }

    fmt.Println("Successfully connected to Database")

    _, err = db.Exec("CREATE TABLE IF NOT EXISTS todosTable (id VARCHAR(255), item VARCHAR(255), completed BOOLEAN)")
    if err != nil {
		
        fmt.Println("Table already exist")
        return
    }

    


}