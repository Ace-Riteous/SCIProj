package core

import (
	"SCIProj/initialize"
	"SCIProj/router"
	"SCIProj/service"
)

func RunWindowsServer(Addr string) {
	go service.ServerHealthCheck()
	initialize.Redis()
	initialize.JWTANDMD()
	router.InitRouter()
}
