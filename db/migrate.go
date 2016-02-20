package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/SLP-KBIT/UniAca-back/app/models"
)

func main() {
	db, _ := gorm.Open("postgres", "user=vitamin dbname=uniaca sslmode=disable")
	db.CreateTable(&models.Question{})
	db.CreateTable(&models.Contest{})
	db.CreateTable(&models.Attend{})
}
