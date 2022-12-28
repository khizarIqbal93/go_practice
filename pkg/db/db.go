package db

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Webpage struct {
	WebpageID    uint   `json:"-"`
	Webpage      string `json:"webpageUrl"`
	Description  string `json:"description"`
	LastModified time.Time
}

func GetLink() Webpage {
	dsn := "root:Passw0rd@/crawl"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	}

	var webpage Webpage
	db.First(&webpage)
	return webpage
}
