package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"meeting/pkg/util"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		token := c.GetHeader("M-Token")
		msg := "success"
		code := 1
		if token == "" {
			code = 0
			msg = "token不能为空"
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				code = 0
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					msg = "token已过期"
				default:
					msg = "token验证失败"
				}
			}
		}

		if code != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 0,
				"msg":  msg,
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
