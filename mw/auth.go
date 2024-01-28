package mw

import (
	"net/http"

	"github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		domain := "https://dev-gzm0pgbh.us.auth0.com/"

		config := auth0.NewConfiguration(
			auth0.NewJWKClient(auth0.JWKClientOptions{URI: domain + ".well-known/jwks.json"}, nil),
			[]string{"http://localhost:4000"},
			domain,
			jose.RS256,
		)

		validator := auth0.NewValidator(config, nil)

		token, err := validator.ValidateRequest(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("user", token)

		c.Next()
	}
}