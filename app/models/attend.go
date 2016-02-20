package models

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

)

type Attend struct {
	ID       int `gorm:"primary_key" json:"id" binding:required`
	ContestID int `json:"contest_id" binding:required`
	Contest  Contest
	Score1   int `sql:"default:-1" json:"score1" binding:required`
	Score2   int `sql:"default:-1" json:"score2" binding:required`
	Score3   int `sql:"default:-1" json:"score3" binding:required`
	Score4   int `sql:"default:-1" json:"score4" binding:required`
	Score5   int `sql:"default:-1" json:"score5" binding:required`
	Score6   int `sql:"default:-1" json:"score6" binding:required`
	Score7   int `sql:"default:-1" json:"score7" binding:required`
	Score8   int `sql:"default:-1" json:"score8" binding:required`
	Score9   int `sql:"default:-1" json:"score9" binding:required`
	Score10  int `sql:"default:-1" json:"score10" binding:required`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (a Attend) GetResultScore() int {
	return a.Score1 + a.Score2 + a.Score3 + a.Score4 + a.Score5 + a.Score6 + a.Score7 + a.Score8 + a.Score9 + a.Score10
}

func (a *Attend) SetScore(num int, score int) {
	var column string
	db, _ := gorm.Open("postgres", "user=vitamin dbname=uniaca sslmode=disable")

	switch num {
	case 1 : column = "score1"
	case 2 : column = "score2"
	case 3 : column = "score3"
	case 4 : column = "score4"
	case 5 : column = "score5"
	case 6 : column = "score6"
	case 7 : column = "score7"
	case 8 : column = "score8"
	case 9 : column = "score9"
	case 10 : column = "score10"
	}

	db.Model(a).Update(column, score)
}
