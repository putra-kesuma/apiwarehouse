package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

//try env conf with godotenv
func ConnectDb()*sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	DB_HOST:=os.Getenv("DB_HOST")
	DB_PORT:=os.Getenv("DB_PORT")
	DB_USER:=os.Getenv("DB_USER")
	DB_PASS:=os.Getenv("DB_PASS")
	DB_SCHEMA:=os.Getenv("DB_SCHEMA")
	DRIVER_DB:=os.Getenv("DRIVER_DB")
	dataSource:=fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",DB_USER,DB_PASS,DB_HOST,DB_PORT,DB_SCHEMA)
	db, err := sql.Open(DRIVER_DB, dataSource)

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
