package handler

import (
	"app/model"
	"github.com/ilibs/gosql"
	"github.com/ilibs/very/core"
)


var AdminUser core.HandlerFunc = func(c *core.Context) core.Response {
	user := getLoginUser(c)
	return c.Success(user)
}