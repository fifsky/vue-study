package middleware

import (
	"strings"

	"app/config"
	"app/model"
	"encoding/base64"

	"github.com/goapt/gee"
	"github.com/goapt/golib/cryptoutil"
	"github.com/goapt/golib/hashing"
	"github.com/ilibs/gosql"
)

var AuthLogin gee.HandlerFunc = func(c *gee.Context) gee.Response {
	accessToken := c.Request.Header.Get("Access-Token")

	if accessToken == "" {
		return c.Fail(201, "Access Token不能为空")
	}

	bt, err := base64.StdEncoding.DecodeString(accessToken)

	if err != nil {
		return c.Fail(201, "Access Token错误"+err.Error())
	}

	cipherText, err := cryptoutil.AesDecrypt([]byte(config.App.Common.TokenSecret), bt)
	if err != nil {
		return c.Fail(201, "Access Token错误")
	}

	src := string(cipherText)
	v := strings.Split(src, ":")
	if len(v) != 2 || hashing.Md5(v[0]+config.App.Common.TokenSecret) != v[1] {
		return c.Fail(201, "Access Token不合法")
	}

	user := &model.Users{}
	err = gosql.Model(user).Where("id = ?", v[0]).Get()
	if err != nil {
		return c.Fail(201, "Access Token错误，用户不存在")
	}

	c.Set("userInfo", user)
	c.Next()
	return nil
}
