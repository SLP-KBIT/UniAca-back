package models

import (
	"time"
)

type Contest struct {
	ID          int `gorm:"primary_key" json:"id" binding:required`
	Question1   int `json:"question1" binding:required`
	Question2   int `json:"question2" binding:required`
	Question3   int `json:"question3" binding:required`
	Question4   int `json:"question4" binding:required`
	Question5   int `json:"question5" binding:required`
	Question6   int `json:"question6" binding:required`
	Question7   int `json:"question7" binding:required`
	Question8   int `json:"question8" binding:required`
	Question9   int `json:"question9" binding:required`
	Question10  int `json:"question10" binding:required`
	Attends []Attend

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (c Contest) GetQuestionId(i int) int {
	var n int

	switch i {
	case 1 : n = c.Question1
	case 2 : n = c.Question2
	case 3 : n = c.Question3
	case 4 : n = c.Question4
	case 5 : n = c.Question5
	case 6 : n = c.Question6
	case 7 : n = c.Question7
	case 8 : n = c.Question8
	case 9 : n = c.Question9
	case 10 : n = c.Question10
	default : n = -1
	}

	return n
}
