package middlewares

import (
	"github.com/gin-gonic/gin"
	"server/pkg/common"
	"server/pkg/jwt"
	"server/pkg/models"
	"strings"
)

func AuthenticateRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			common.Respond(c, &common.APIResponse{
				Status:  common.APIStatus.Unauthorized,
				Message: "Missing token",
			})
			return
		}

		s := strings.Split(token, " ")

		tokenString := s[1]

		if tokenString == "" {
			common.Respond(c, &common.APIResponse{
				Status:  common.APIStatus.Unauthorized,
				Message: "Invalid token",
			})
			c.Abort()
			return
		}

		claims, err := jwt.VerifyToken(tokenString)
		if err != nil {
			common.Respond(c, &common.APIResponse{
				Status:    common.APIStatus.Unauthorized,
				Message:   "Invalid token",
				RootCause: err,
			})
			c.Abort()
			return
		}
		c.Set("currentUser", &models.User{CommonModel: models.CommonModel{ID: claims.UserID}})

		c.Next()
	}
}
