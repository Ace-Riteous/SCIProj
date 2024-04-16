package router

import (
	"SCIProj/api"
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {
	r.GET("/see_competitions", api.SeeCompetitions)
	r.POST("/add_competition", api.AddCompetition)
	r.POST("/login", api.Login)
	r.POST("/register", api.Register)

}
