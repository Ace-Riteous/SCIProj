package middleware

import (
	"SCIProj/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuthMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		skipList := []string{
			"/user/login",
			"/user/register",
			"/user/home",
			"/post/see_competition",
			"/post/add_competition",
			"/team/get_team_is_not_full",
			"/team/get_team",
			"/team/new_team",
		}
		for _, item := range skipList {
			if item == c.FullPath() {
				c.Next()
				return
			}
		}
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("sciproj"), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if claims, ok := token.Claims.(*utils.Claims); ok && token.Valid {
			c.Set("uid", claims.Uid)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
	}
}
