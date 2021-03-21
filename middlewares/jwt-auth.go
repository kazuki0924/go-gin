package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kazuki0924/go-gin/service"
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println("Claims[Name]", claims["name"])
			fmt.Println("Claims[Admin]", claims["admin"])
			fmt.Println("Claims[Issuer]", claims["iss"])
			fmt.Println("Claims[IssuedAt]", claims["iat"])
			fmt.Println("Claims[ExpiresAt]", claims["exp"])
		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
