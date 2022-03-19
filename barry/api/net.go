package api

import (
	"barry/config"
	"barry/model"
	"encoding/json"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
)

var db = config.Db()

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

func AddNet(c *gin.Context){
	appG := Gin{C: c}
	addInfo := model.Net{ID: 98, Command: "yi" }
	db.Create(&addInfo)
	addJson,_ := json.Marshal(addInfo)
	data := string(addJson)
	appG.Response(200,0,data)
}


func DelNet(c *gin.Context) {
	//appG := Gin{C: c}
	//db.Model()




}






