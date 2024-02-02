package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Create_Table(db *sql.DB, name string) {

	//Note CREARTE TABLE/DATBASE "IF NOT EXITS" is not valid in PSQL
	//so there will be problems for migrations,creating db,tables etc.. stuff as only one time excution would be granted if you excute again this it will show errror like this is "Already exsits"
	//to avoide you can use pgX or sqlX library instead pure Psql library

	str := fmt.Sprintf("CREATE TABLE %s (username varchar(45) NOT NULL,password varchar(450) NOT NULL,enabled integer NOT NULL DEFAULT '1',PRIMARY KEY (username))", name)

	_, err := db.Exec(str)

	if err != nil {
		panic(err)
	} else {
		log.Println("Created a new table ", str)
	}

}

func main() {
	// connection string

	psqlconn := "postgres://postgres:55120@localhost/shubham_mishra_pvt_ltd?sslmode=disable"

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// to check wheather we connected to PSQL or not
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	Create_Table(db, "Details")

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
