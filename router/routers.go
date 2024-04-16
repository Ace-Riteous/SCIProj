package router

import (
	"SCIProj/api"
	"SCIProj/middleware"
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {
	UserGroup := r.Group("/user")
	{
		UserGroup.POST("/login", api.Login)
		UserGroup.POST("/register", api.Register)
	}
	PostGroup := r.Group("/post")
	{
		PostGroup.GET("/see_competitions", api.SeeCompetitions)
		PostGroup.POST("/add_competition", api.AddCompetition).Use(middleware.JWTAuthMiddelware())
	}
}
