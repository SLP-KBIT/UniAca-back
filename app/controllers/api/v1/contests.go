package controllers

import (
	"github.com/revel/revel"
	//"github.com/SLP-KBIT/UniAca-back/app/controllers"
)

type ApiV1Contests struct {
	ApiV1Controller
}

func (c ApiV1Contests) New() revel.Result {
	r := Response{"new"}
	return c.RenderJson(r)
}
