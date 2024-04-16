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
		PostGroup.GET("/see_competitions", api.GetCompetition)
		PostGroup.POST("/add_competition", api.AddCompetition)
	}
}
