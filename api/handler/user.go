package handler

import (
	"app/model"
	"github.com/ilibs/gosql"
	"github.com/ilibs/very/core"
)

var UserList core.HandlerFunc = func(c *core.Context) core.Response {
	p := struct {
		Id int `json:"id"`
	}{}

	if err := c.ShouldBindJSON(p); err != nil {
		return c.Fail(201, err)
	}

	users := make([]*model.Users,0)
	err := gosql.Model(&users).Limit(10).All()
	if err != nil {
		c.Fail(500, err)
	}

	return c.Success(users)
}
