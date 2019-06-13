package handler

import (
	"app/model"
	"github.com/goapt/gee"
)

func getLoginUser(c *gee.Context) *model.Users {
	if u, ok := c.Get("userInfo"); ok {
		return u.(*model.Users)
	}
	return nil
}
