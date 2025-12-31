package middleware

import (
	"strings"

	"thuchanhgolang/pkg/jwt"
	"thuchanhgolang/pkg/response"

	"github.com/gin-gonic/gin"
)

func (mw *implMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")
		if tokenString == "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		payload, err := mw.jwtMgr.Verify(tokenString)
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		ctx := c.Request.Context()
		ctx = jwt.SetPayloadToContext(ctx, payload)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}

}
