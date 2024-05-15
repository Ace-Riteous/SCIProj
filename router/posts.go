package router

import (
	"SCIProj/api"
	"github.com/gin-gonic/gin"
)

func InitPostRouters() {
	//注册路由
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		postApi := api.NewPostApi()
		rgPublicPost := rgPublic.Group("post")
		{
			rgPublicPost.GET("/competition_all", postApi.GetCompetitionAll)
		}
		rgAuthPost := rgAuth.Group("post")
		{
			rgAuthPost.POST("/competition_new", postApi.AddCompetition)
		}
	})
}
