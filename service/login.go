package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-Admin/define"
	"go-Admin/helper"
	"go-Admin/models"
	"gorm.io/gorm"
	"net/http"
)

func LoginPassWord(c *gin.Context) {
	in := new(LoginPassWordRequest)
	err := c.ShouldBindJSON(in)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	//根据账号和密码查询用户
	sysUser, err := models.GetUserByNamePassWord(in.UserName, in.PassWord)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或者密码错误",
			})
			return
		}
	}

	// 生成token
	token, err := helper.GenerateToken(sysUser.ID, sysUser.UserName, define.TokenExpire)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	//刷新token
	refreshToken, err := helper.GenerateToken(sysUser.ID, sysUser.PassWord, define.ReFreshTokenExpire)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	data := &LoginPassWordReply{
		Token:        token,
		ReFreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"msg":      "登陆成功",
		"result":   data,
		"userInfo": sysUser,
	})
}
