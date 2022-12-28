package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func AddLink() {
	db, err := sql.Open("mysql", "root:Passw0rd@/crawl")
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()
}

func PostLink() {
	db, err := sql.Open("mysql", "root:Passw0rd@/crawl")
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO webpages (webpage, description) VALUES('https://go.dev/blog/maps', 'hello')")
	if err != nil {
		log.Println(err.Error())
	}
	defer insert.Close()
}
