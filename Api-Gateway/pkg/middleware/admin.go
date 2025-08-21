package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/helper"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/response"
)

// AdminAuthMiddleware is a middleware for validating admin tokens.
func AdminAuthMiddleware() gin.HandlerFunc {
	return func (c *gin.Context)  {
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			resp := response.ClientResponse(http.StatusUnauthorized, "No auth header provided", nil, nil)
			c.JSON(http.StatusUnauthorized, resp)
			c.Abort()
			return 
		}
		tokenString := helper.GetTokenFromHeader(tokenHeader)
		claims, err := helper.ValidateToken(tokenString)
		if err != nil {
			resp := response.ClientResponse(http.StatusUnauthorized, "Invalid Token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, resp)
			c.Abort()
			return
		}

		if claims.Role != "admin" {
			resp := response.ClientResponse(http.StatusUnauthorized, "Insufficient permissions", nil, "Admin role required")
			c.JSON(http.StatusUnauthorized, resp)
			c.Abort()
			return
		}

		c.Set("user_role", claims.Role)
		c.Set("tokenClaims", claims)
		c.Next()
	}
}