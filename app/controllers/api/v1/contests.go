package controllers

import (
	"github.com/revel/revel"
	"../../../controllers"
)

type ApiV1Contests struct {
	ApiV1Controller
}

func (c ApiV1Contests) New() revel.Result {
	r := Reponse{"new"}
	return RenderJson(r)
}
