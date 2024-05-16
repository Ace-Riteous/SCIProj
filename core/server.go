package core

import (
	"SCIProj/initialize"
	"SCIProj/router"
	"SCIProj/service"
	"fmt"
	"time"
)

func RunWindowsServer() {
	go service.ServerHealthCheck()
	initialize.Redis()
	initialize.JWTANDMD()
	fmt.Println(time.Now().Unix())
	router.InitRouter()
}
