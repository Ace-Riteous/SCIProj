package router

import (
	"SCIProj/api"
	"github.com/gin-gonic/gin"
)

func InitTeamRouters() {
	//注册路由
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		teamApi := api.NewTeamApi()
		rgPublicTeam := rgPublic.Group("team")
		{
			rgPublicTeam.GET("/team_all", teamApi.GetTeamAll)
			rgPublicTeam.GET("/team_is_not_full", teamApi.GetTeamNotFull)
		}
		rgAuthTeam := rgAuth.Group("user")
		{
			rgAuthTeam.POST("/team_new", teamApi.NewTeam)
			rgAuthTeam.PUT("/join_team", teamApi.JoinTeam)
		}
	})
}
