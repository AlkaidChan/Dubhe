package jwt

import (
	"github.com/alkaidchen/Dubhe/internal/utils/errmsg"
	"github.com/alkaidchen/Dubhe/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// exclude api for login
		if c.Request.URL.Path == "/api/login" {
			c.Next()
		} else {
			token := c.Request.Header.Get("Authorization")
			if token == "" {
				c.AbortWithStatusJSON(http.StatusOK, errmsg.Fail(errmsg.ErrorTokenNotExist))
				return
			}
			claims, err := jwt.ParseToken(token)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusOK, errmsg.Fail(errmsg.ErrorTokenInvalid))
				return
			}
			if claims.VerifyExpiresAt(time.Now(), true) == false {
				c.AbortWithStatusJSON(http.StatusOK, errmsg.Fail(errmsg.ErrorTokenInvalid))
				return
			}
			c.Set("uid", claims.ID)
			c.Next()
		}
	}
}
