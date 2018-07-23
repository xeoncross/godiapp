package main

import (
	_ "database/sql/driver"
	"log"

	"bitbucket.org/xeoncross/godiapp"
	"bitbucket.org/xeoncross/godiapp/config"
	"bitbucket.org/xeoncross/godiapp/http"
	"bitbucket.org/xeoncross/godiapp/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var err error
	config, err := config.Load("config.json")

	if err != nil {
		log.Fatal(err)
	}

	var db godiapp.UserService
	db, err = mysql.NewDB(config.MySQL.DSN)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server starting on", config.HTTP.Address)
	http.StartServer(config, db)
}
