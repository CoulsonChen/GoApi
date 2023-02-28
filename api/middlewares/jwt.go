package middlewares

import (
	"net/http"
	"strings"

	"github.com/CoulsonChen/GoApi/pkg/services"
	"github.com/gin-gonic/gin"
)

type JwtMiddleware struct {
	service                services.IJwtService
	JwtSecretKey           string
	JwtTokenExpireDuration int
}

func JwtMiddlewareProvider(s services.IJwtService) *JwtMiddleware {
	return &JwtMiddleware{
		service:                s,
		JwtSecretKey:           "cc2e8a10-2512-4024-8580-84f06109630c",
		JwtTokenExpireDuration: 120,
	}
}

// JWTAuthMiddleware Middleware of JWT
func (j *JwtMiddleware) JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Get token from Header.Authorization field.
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Authorization is null in Header",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Format of Authorization is wrong",
			})
			c.Abort()
			return
		}
		// parts[0] is Bearer, parts is token.
		mc, err := j.service.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Invalid Token.",
			})
			c.Abort()
			return
		}
		// Store Account info into Context
		c.Set("account", mc.Account)
		// After that, we can get Account info from c.Get("account")
		c.Next()
	}
}
