package router

import (
	"barry/api"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apis := r.Group("barry/api")
	{
		//获取标签列表
		apis.GET("/get_net_info", api.AddNet)
	}

	return r
}
