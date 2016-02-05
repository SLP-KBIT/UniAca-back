package models

import (
	"time"
)

type Question struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Content  string `sql:"size:255" json:"content"`
	Select1  string `sql:"size:255" json:"select1"`
	Select2  string `sql:"size:255" json:"select2"`
	Select3  string `sql:"size:255" json:"select3"`
	Select4  string `sql:"size:255" json:"select4"`
	Answer   string `sql:"size:255" json:"answer"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
