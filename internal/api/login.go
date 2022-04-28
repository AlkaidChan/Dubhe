package api

import (
	"github.com/alkaidchen/Dubhe/internal/dto"
	"github.com/alkaidchen/Dubhe/internal/utils/errmsg"
	"github.com/alkaidchen/Dubhe/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	var data dto.LoginForm
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println("decode json data failed, ", err.Error())
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}
	if data.Username != "admin" || data.Password != "admin" {
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorPasswordWrong))
		return
	}
	user := struct {
		ID uint `json:"id"`
	}{
		ID: 1,
	}
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		log.Println("generate token failed, ", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	c.JSON(http.StatusOK, errmsg.Success(token))
}

func Check(c *gin.Context) {
	c.JSON(http.StatusOK, errmsg.Success("ok"))
}
