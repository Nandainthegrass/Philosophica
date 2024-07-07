package dbase

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nandainthegrass/Philosophica/types"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func NewMYSQLStorage(config mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {

		return nil, err
	}
	return db, nil
}

func InitStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database")
	}
	fmt.Println("Database connected successfully")
}

func PopulateDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env files")
	}
	res, err := http.Get("https://philosophy-quotes-api.glitch.me/quotes")
	if err != nil {
		log.Fatal("Error fetching quotes")
	}
	defer res.Body.Close()

	var quotes []types.Quote

	json.NewDecoder(res.Body).Decode(&quotes)
	db, err := NewMYSQLStorage(mysql.Config{
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
	defer db.Close()

	InitStorage(db)
	stmt, err := db.Prepare("Insert into allquotes values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Error preparing statement")
	}
	defer stmt.Close()

	for _, quote := range quotes {
		_, err := stmt.Exec(quote.ID, quote.Source, quote.Philosophy, quote.Quote)
		if err != nil {
			log.Fatal("Error occured:", err)
		}
	}
}
