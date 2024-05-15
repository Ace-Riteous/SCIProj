package initialize

import (
	"SCIProj/global"
)

func JWTANDMD() {
	secret := global.VP.GetString("jwt.signing-key")
	global.JWTKey = []byte(secret)
}
