package controllers

import (
	"github.com/revel/revel"
	"github.com/SLP-KBIT/UniAca-back/app/controllers"
	"github.com/SLP-KBIT/UniAca-back/app/models"
)

type ApiV1Questions struct {
	ApiV1Controller
}

type Quiz struct {
	Id int `json:"id"`
	Text string `json:"text"`
	Choices [4]string `json:"choices"`
}


func (c ApiV1Questions) Get(num int) revel.Result {
	question := &models.Question{}
	if err := controllers.DB.First(&question, num).Error; err != nil {
		return c.RenderJson(ErrorResponse{ Code: 404, Message: "Invalid id" })
	}
	q := Quiz{Id: question.ID, Text: question.Content}
	q.Choices = [4]string{question.Select1, question.Select2, question.Select3, question.Select4}
	return c.RenderJson(q)
}

func (c ApiV1Questions) Answer(num int, answer string) revel.Result {
	question := &models.Question{}
	if err := controllers.DB.First(&question, num).Error; err != nil {
		return c.RenderJson(ErrorResponse{ Code: 404, Message: "Invalid id" })
	}

	var r Response
	if answer == question.Answer {
		r = Response{"correct"}
	} else {
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
