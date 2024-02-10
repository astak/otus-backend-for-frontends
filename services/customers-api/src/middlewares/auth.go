package middlewares

import (
	"strings"

	"github.com/Astak/otus-docker-basics-homework/web-service-gin/auth"
	common_auth "github.com/Astak/otus-docker-basics-homework/web-service-gin/common/auth"
	"github.com/gin-gonic/gin"
)

func Auth(jwtKey []byte) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		splitToken := strings.Split(tokenString, "Bearer ")
		tokenString = splitToken[1]
		claims, err := auth.ValidateToken(tokenString, jwtKey)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Set(common_auth.JwtKey, *claims)
		context.Next()
	}
}
