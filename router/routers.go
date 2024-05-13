package router

import (
	"SCIProj/api"
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {
	UserGroup := r.Group("/user")
	{
		UserGroup.POST("/login", api.Login)
		UserGroup.POST("/register", api.Register)
		UserGroup.GET("/home", api.GetUserInfo)
	}
	PostGroup := r.Group("/post")
	{
		PostGroup.GET("/competition_all", api.GetCompetitionAll)
		PostGroup.POST("/competition_new", api.AddCompetition)
	}
	TeamGroup := r.Group("/team")
	{
		TeamGroup.POST("/team_new", api.NewTeam)
		TeamGroup.GET("/team_all", api.GetTeamAll)
		TeamGroup.GET("/team_is_not_full", api.GetTeamNotFull)
		TeamGroup.PUT("/join_team", api.JoinTeam)
	}
}
