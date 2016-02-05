package main

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/SLP-KBIT/UniAca-back/app/model"
)

func main() {
	dbSql, _ := sql.Open("postgres", "user=vitamin dbname=uniaca sslmode=disable")
	db, _ := gorm.Open("postgres", dbSql)
	db.CreateTable(&models.Question{})
}
