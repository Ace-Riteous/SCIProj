package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
}

func Success(c *gin.Context, data interface{}) {
	var result Result
	result.Code = http.StatusOK
	result.Error = ""
	result.Data = data
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, result)
}

func Error(c *gin.Context, err error) {
	var result Result
	result.Code = -999
	result.Error = err.Error()
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, result)
}
