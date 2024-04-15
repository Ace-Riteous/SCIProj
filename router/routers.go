package router

import (
	"SCIProj/api"
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {

	r.GET("/see-competitions", api.SeeCompetitions)
	r.POST("/add-competition", api.AddCompetition)

}
