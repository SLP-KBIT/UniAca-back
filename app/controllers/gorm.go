package controllers

import (
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"github.com/SLP-KBIT/UniAca-back/app/models"
	"log"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("postgres", "user=vitamin dbname=uniaca sslmode=disable")
	if err != nil {
		log.Panicf("Failed to connect to database: %n\n", err)
	}

	db.DB()
	db.AutoMigrate(&models.Question{})
	DB = &db
}

func dbInfoString() string {
	s, b := revel.Config.String("db.info")
	if !b {
		log.Panicf("database info not found")
	}

	return s
}
