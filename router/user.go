package router

import (
	"SCIProj/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouters() {
	//注册路由
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		useApi := api.NewUserApi()
		rgPublicUser := rgPublic.Group("user")
		{
			rgPublicUser.POST("/login", useApi.Login)
			rgPublicUser.POST("/register", useApi.Register)
		}
		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.GET("/home", useApi.GetUserInfo)
		}
	})
}
