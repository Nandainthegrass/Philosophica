package main

import (
	"log"
	"os"

	"github.com/Nandainthegrass/Philosophica/cmd/api"
	"github.com/Nandainthegrass/Philosophica/dbase"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	//Now then we need to populate the database
	//Uncomment the line if you want to populate the database
	//dbase.PopulateDB()
	godotenv.Load()
	db, err := dbase.NewMYSQLStorage(mysql.Config{
		User:                 os.Getenv("USER"),
		Passwd:               os.Getenv("PASSWORD"),
		DBName:               os.Getenv("DBNAME"),
		Addr:                 os.Getenv("ADDR"),
		Net:                  "tcp",
		AllowNativePasswords: true,
	})
	if err != nil {
		log.Fatal("Err connecting to database", err)
	}
	dbase.InitStorage(db)

	s := api.NewAPIServer(":5000", db)
	err = s.Run()
	if err != nil {
		log.Fatal("Error while running server")
	}

}
