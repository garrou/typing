package middlewares

import (
	"net/http"
	"strings"
	"typing/utils"

	"github.com/gin-gonic/gin"
)

func AuthorizeJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.NewResponse("User not authenticated", nil))
			return
		}
		items := strings.Split(authHeader, " ")

		if len(items) != 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.NewResponse("User not authenticated", nil))
			return
		}
		token, err := utils.ValidateToken(items[1])

		if !token.Valid || err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewResponse("Invalid token", nil))
		}
		c.Set("userId", utils.ExtractUserId(token))
	}
}