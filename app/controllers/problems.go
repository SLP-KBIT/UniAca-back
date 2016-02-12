package controllers

import (
	"github.com/revel/revel"
	"github.com/SLP-KBIT/UniAca-back/app/models"
)

type Problems struct {
	*revel.Controller
}

func (c Problems) New() revel.Result {
	var err, msg string;

	if c.Session["err"] != "" {
		err = c.Session["err"]
		c.Session["err"] = ""
	}

	if c.Session["msg"] != "" {
		msg = c.Session["msg"]
		c.Session["msg"] = ""
	}

	return c.Render(err, msg)
}

func (c Problems) Create(content string, select1 string, select2 string, select3 string, select4 string) revel.Result {
	c.Validation.Required(content != "").Message("問題文が未入力です")
	c.Validation.Required(select1 != "").Message("正答/選択肢1が未入力です")
	c.Validation.Required(select2 != "").Message("選択肢2が未入力です")
	c.Validation.Required(select3 != "").Message("選択肢3が未入力です")
	c.Validation.Required(select4 != "").Message("選択肢4が未入力です")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Problems.New)
	}

	question := &models.Question{}
	question.Content = content
	question.Select1 = select1
	question.Select2 = select2
	question.Select3 = select3
	question.Select4 = select4
	question.Answer = question.Select1

	if err := DB.Create(question).Error; err != nil {
		c.Session["err"] =  "データの登録に失敗しました"
		return c.Redirect(Problems.New)
	}

	c.Session["msg"] = "登録が完了しました"
	return c.Redirect(Problems.New)
}
