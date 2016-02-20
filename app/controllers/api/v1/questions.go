package controllers

import (
	"github.com/revel/revel"
	"github.com/SLP-KBIT/UniAca-back/app/controllers"
	"github.com/SLP-KBIT/UniAca-back/app/models"
	"math/rand"
	"time"
)

type ApiV1Questions struct {
	ApiV1Controller
}

type Quiz struct {
	Id int `json:"id"`
	Text string `json:"text"`
	Choices [4]string `json:"choices"`
}


func (c ApiV1Questions) Get(attendId int, num int) revel.Result {
	var attend models.Attend
	var contest models.Contest
	var question models.Question

	controllers.DB.Where("id = ?", attendId).First(&attend)
	controllers.DB.Model(&attend).Related(&contest)
	n := contest.GetQuestionId(num)

	if err := controllers.DB.First(&question, n).Error; err != nil {
		return c.RenderJson(ErrorResponse{ Code: 404, Message: "Invalid id" })
	}
	q := Quiz{Id: question.ID, Text: question.Content}
	choices := []string{question.Select1, question.Select2, question.Select3, question.Select4}
	q.Choices[0] = randPopChoices(&choices)
	q.Choices[1] = randPopChoices(&choices)
	q.Choices[2] = randPopChoices(&choices)
	q.Choices[3] = randPopChoices(&choices)
	return c.RenderJson(q)
}

func (c ApiV1Questions) Answer(attendId int, num int, answer string) revel.Result {
	var attend models.Attend
	var contest models.Contest
	var question models.Question

	controllers.DB.Where("id = ?", attendId).First(&attend)
	controllers.DB.Model(&attend).Related(&contest)
	n := contest.GetQuestionId(num)

	if err := controllers.DB.First(&question, n).Error; err != nil {
		return c.RenderJson(ErrorResponse{ Code: 404, Message: "Invalid id" })
	}

	var r Response
	if answer == question.Answer {
		attend.SetScore(num, 10)
		r = Response{"correct"}
	} else {
		attend.SetScore(num, 0)
		r = Response{"wrong"}
	}
	return c.RenderJson(r)
}

func (c ApiV1Questions) Create() revel.Result {
	question := &models.Question{}

	if err := c.BindParams(question); err != nil {
		return c.RenderJson(ErrorResponse{ Code: 400, Message: "Bad Request" })
	}

	if err := controllers.DB.Create(question).Error; err != nil {
		return c.RenderJson(ErrorResponse{ Code: 500, Message: "Internal Server Error" })
	}
	return c.RenderJson(Response{"create"})
}

func randPopChoices(sl *[]string) string {
	var tmp []string
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(*sl))
	a := (*sl)[i]

	for k := 0; k < len(*sl); k++ {
		if ( k == i ) { continue }
		tmp = append(tmp, (*sl)[k])
	}
	*sl = tmp

	return a
}
