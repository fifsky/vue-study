package handler

import (
	"github.com/goapt/gee"
)

var AdminUser gee.HandlerFunc = func(c *gee.Context) gee.Response {
	user := getLoginUser(c)
	return c.Success(user)
}
