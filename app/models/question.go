package models

import (
	"time"
)

type Question struct {
	ID       int    `gorm:"primary_key" json:"id" binding:required`
	Content  string `sql:"size:255" json:"content" binding:required`
	Select1  string `sql:"size:255" json:"select1" binding:required`
	Select2  string `sql:"size:255" json:"select2" binding:required`
	Select3  string `sql:"size:255" json:"select3" binding:required`
	Select4  string `sql:"size:255" json:"select4" binding:required`
	Answer   string `sql:"size:255" json:"answer" binding:required`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
