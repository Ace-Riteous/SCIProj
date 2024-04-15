package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeCompetitions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "competition1",
	})
}

func AddCompetition(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "add competition success",
	})
}
