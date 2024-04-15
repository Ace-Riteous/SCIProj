package model

import (
	"encoding/json"
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
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, resultJson)

}

func Error(c *gin.Context, err error) {
	var result Result
	result.Code = -999
	result.Error = err.Error()
	resultJson, _ := json.Marshal(result)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, resultJson)
}