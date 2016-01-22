package controllers

import (
	"github.com/revel/revel"
	"net/http"
	"github.com/SLP-KBIT/UniAca-back/app/utils"
)

type ApiV1Controller struct {
	*revel.Controller
}

type ErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Results interface{} `json:"results"`
}

func (c *ApiV1Controller) BindParams(s interface{}) error {
	return utils.JsonDecode(c.Request.Body, s)
}

func (c *ApiV1Controller) HandleBadRequestError(s string) revel.Result {
	c.Response.Status = http.StatusBadRequest
	r := ErrorResponse{ c.Response.Status, s }
	return c.RenderJson(r)
}

func (c *ApiV1Controller) HandleInernalServerError(s string) revel.Result {
	c.Response.Status = http.StatusInternalServerError
	r := ErrorResponse{ c.Response.Status, s }
	return c.RenderJson(r)
}
