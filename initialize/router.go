package initialize

import (
	"SCIProj/middleware"
	"SCIProj/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	Router.Use(middleware.CORSMiddleware(), middleware.ZapLogger(), middleware.JWTAuthMiddelware())

	router.SetRouters(Router)

	return Router
}
