package initialize

import (
	"SCIProj/middleware"
	"SCIProj/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	Router.Use(middleware.CORSMiddleware())

	router.SetRouters(Router)

	return Router
}
