package handler

import (
	"github.com/ilibs/very/core"
)

var AdminUser core.HandlerFunc = func(c *core.Context) core.Response {
	user := getLoginUser(c)
	return c.Success(user)
}
