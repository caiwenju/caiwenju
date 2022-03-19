package api

import (
	"barry/config"
	_ "barry/config"
	"encoding/json"
	"fmt"
	_ "github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
)


type ReportTables interface {
	GetFilterInfo(ctx *gin.Context) *gin.Context
	GetTableInfo(ctx *gin.Context) *gin.Context
}

type ReportTable struct {
}

type Params struct {
	XAxis          string `form:"x_axis" json:"x_axis" uri:"x_axis" xml:"x_axis" binding:"required"`
	Project        string `form:"project" json:"project" uri:"project" xml:"project" binding:"required"`
	StartTime      string `form:"start_time" json:"start_time" uri:"start_time" xml:"start_time"`
	EndTime        string `form:"end_time" json:"end_time" uri:"end_time" xml:"end_time"`
	OtherCondition string `form:"other_condition" json:"other_condition" uri:"other_condition" xml:"other_condition"`
}

func (r *ReportTable) MergeCondition(condition []interface{}) {
	for inx, val := range condition {

	}
}

func (r *ReportTable) GetFilterInfo(c *gin.Context) {
	// 获取筛选条件信息
	appG := Gin{C: c}
	rspDate := map[string]interface{}{
		"FilterFiledLst": config.FilterFiledLst,
		"FilterLst":      config.FilterLst,
	}
	addJson, _ := json.Marshal(rspDate)
	data := string(addJson)
	appG.Response(200, 0, data)
}

func (r ReportTable) GetTableInfo(c *gin.Context) {
	// 获取筛选结果
	var filterMap Params
	appG := Gin{C: c}
	var filter map[string]interface{}
	if err := c.ShouldBindJSON(&filterMap); err != nil {
		appG.Response(500, 1, filterMap)
	} else {
		if filterMap.StartTime != "" {
			filter["start_time"] = filterMap.StartTime
		}
		if filterMap.EndTime != "" {
			filter["end_time"] = filterMap.EndTime
		}

		// 整理合并筛选条件

	}

}
