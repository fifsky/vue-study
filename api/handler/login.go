package handler

import (
	"app/config"
	"app/model"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goapt/golib/cryptoutil"
	"github.com/goapt/golib/hashing"

	"github.com/goapt/gee"
	"github.com/ilibs/gosql"
)

var Login gee.HandlerFunc = func(c *gee.Context) gee.Response {
	p := &struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}{}

	if err := c.ShouldBindJSON(p); err != nil {
		return c.Fail(201, err)
	}

	if p.UserName == "" || p.Password == "" {
		return c.Fail(201, "用户名密码不能为空")
	}

	user := &model.Users{Name: p.UserName, Password: hashing.Md5(p.Password)}
	err := gosql.Model(user).Get()
	if err != nil {
		return c.Fail(202, "用户名或密码错误")
	}

	src := fmt.Sprintf("%d:%s", user.Id, hashing.Md5(fmt.Sprintf("%d%s", user.Id, config.App.Common.TokenSecret)))
	cipherText, err := cryptoutil.AesEncrypt([]byte(config.App.Common.TokenSecret), []byte(src))
	if err != nil {
		return c.Fail(201, "Access Token加密错误"+err.Error())
	}

	return c.Success(gin.H{
		"access_token": base64.StdEncoding.EncodeToString(cipherText),
		"user":         user,
	})
}
