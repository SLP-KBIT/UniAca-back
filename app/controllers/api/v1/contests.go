package controllers

import (
	"github.com/revel/revel"
	"github.com/SLP-KBIT/UniAca-back/app/controllers"
	"github.com/SLP-KBIT/UniAca-back/app/models"
	"math/rand"
	"time"
)

type ApiV1Contests struct {
	ApiV1Controller
}

type Contest struct {
	Id string `json:"id"`
}

func (c ApiV1Contests) New() revel.Result {
	var ids []int

	contest := &models.Contest{}
	controllers.DB.Model(&models.Question{}).Pluck("id", &ids)
	contest.Question1 = randPop(&ids)
	contest.Question2 = randPop(&ids)
	contest.Question3 = randPop(&ids)
	contest.Question4 = randPop(&ids)
	contest.Question5 = randPop(&ids)
	contest.Question6 = randPop(&ids)
	contest.Question7 = randPop(&ids)
	contest.Question8 = randPop(&ids)
	contest.Question9 = randPop(&ids)
	contest.Question10 = randPop(&ids)
	controllers.DB.Create(&contest)

	attend := &models.Attend{}
	attend.ContestID = contest.ID
	controllers.DB.Create(&attend)

	return c.RenderJson(Response{attend.ID})
}

func randPop(sl *[]int) int {
	var tmp []int
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

func (c ApiV1Contests) Result(attendId int) revel.Result {
	var attend models.Attend
	controllers.DB.Where("id = ?", attendId).First(&attend)

	return c.RenderJson(Response{attend.GetResultScore()})
}
