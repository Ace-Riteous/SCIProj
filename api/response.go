package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type ResponseJson struct {
	Status int    `json:"-"`    //http状态码
	Code   int    `json:"code"` //服务标记码
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
	Total  int64  `json:"total,omitempty"`
}

func (m ResponseJson) IsEmpty() bool {
	return reflect.DeepEqual(m, ResponseJson{})
}

func httpResponse(ctx *gin.Context, status int, resp ResponseJson) {
	if resp.IsEmpty() {
		//结构为空
		ctx.AbortWithStatus(status)
	}

	ctx.AbortWithStatusJSON(status, resp)
}

// 构建状态码
func buildStatus(resp ResponseJson, defaultStatus int) int {
	if resp.Status == 0 {
		return defaultStatus
	}

	return resp.Status
}

// OK 用于请求成功的回应
func OK(ctx *gin.Context, resp ResponseJson) {
	//状态码在resp中status为空的时候默认statusOK
	ctx.AbortWithStatusJSON(buildStatus(resp, http.StatusOK), resp)
}

// Fail 用于请求失败的回应
func Fail(ctx *gin.Context, resp ResponseJson) {
	ctx.AbortWithStatusJSON(buildStatus(resp, http.StatusOK), resp)
}

// Unauthorized 用于未授权的回应
func Unauthorized(ctx *gin.Context, resp ResponseJson) {
	ctx.AbortWithStatusJSON(buildStatus(resp, http.StatusUnauthorized), resp)
}

// ServerFail 用于服务器级别故障
func ServerFail(ctx *gin.Context, resp ResponseJson) {
	ctx.AbortWithStatusJSON(buildStatus(resp, http.StatusInternalServerError), resp)
}
