package controllers

import (
	"github.com/revel/revel"
	//"github.com/SLP-KBIT/UniAca-back/app/controllers"
)

type ApiV1Questions struct {
	ApiV1Controller
}

type Question struct {
	Id int `json:"id"`
	Text string `json:"text"`
	Choices [4]string `json:"choices"`
}


func (c ApiV1Questions) Get() revel.Result {
	r := Question{Id:1, Text:"ネットワーク上の機器を識別するために指定するネットワーク層における識別用の番号は何か"}
	r.Choices = [4]string{"IPアドレス", "MACアドレス", "メールアドレス", "郵便番号"}
	return c.RenderJson(r)
}

func (c ApiV1Questions) Answer(ans int) revel.Result {
	var r Response
	if ans == 0 {
		r = Response{"correct"}
	} else {
		r = Response{"wrong"}
	}
	return c.RenderJson(r)
}
