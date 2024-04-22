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
		PostGroup.GET("/see_competition", api.GetCompetitionAll)
		PostGroup.POST("/add_competition", api.AddCompetition)
	}
	TeamGroup := r.Group("/team")
	{
		TeamGroup.POST("/new_team", api.NewTeam)
		TeamGroup.POST("/get_team", api.GetTeamAll)
		TeamGroup.GET("/get_team_is_not_full", api.GetTeamNotFull)
		TeamGroup.PUT("/join_team", api.JoinTeam)
	}
}
