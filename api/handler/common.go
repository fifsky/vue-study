package handler

import (
	"app/model"
	"github.com/ilibs/very/core"
)

func getLoginUser(c *core.Context) *model.Users {
	if u, ok := c.Get("userInfo"); ok {
		return u.(*model.Users)
	}
	return nil
}
