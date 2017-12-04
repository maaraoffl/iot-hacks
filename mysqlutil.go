package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strconv"
	"time"
)

func InsertState(ipAddr string, deviceName string, binaryState int) {

	log.Println(time.Now())

	dbInstance := os.Getenv("MYSQL_DB_INSTANCE")
	dbPort := 3306
	database := "IoT"
	dbUser := os.Getenv("MYSQL_DB_USER")
	dbPassword := os.Getenv("MYSQL_DB_PASSWORD")

	// fmt.Printf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbInstance, strconv.Itoa(dbPort), database)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbInstance, strconv.Itoa(dbPort), database))
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
	defer db.Close()
	stmtIns, err := db.Prepare("INSERT INTO StateChangeInfo VALUES(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(ipAddr, deviceName, time.Now(), binaryState)

}
