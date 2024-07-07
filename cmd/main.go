package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Nandainthegrass/Philosophica/dbase"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Quote struct {
	ID         string `json:"id"`
	Source     string `json:"source"`
	Philosophy string `json:"philosophy"`
	Quote      string `json:"quote"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env files")
	}
	res, err := http.Get("https://philosophy-quotes-api.glitch.me/quotes")
	if err != nil {
		log.Fatal("Error fetching quotes")
	}
	defer res.Body.Close()

	var quotes []Quote

	json.NewDecoder(res.Body).Decode(&quotes)

	//Now then we need to populate the database
	//let's open a connection first
	db, err := dbase.NewMYSQLStorage(mysql.Config{
		User:                 os.Getenv("USER"),
		Passwd:               os.Getenv("PASSWORD"),
		DBName:               os.Getenv("DBNAME"),
		Addr:                 os.Getenv("ADDR"),
		Net:                  "tcp",
		AllowNativePasswords: true,
	})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	dbase.InitStorage(db)
}
