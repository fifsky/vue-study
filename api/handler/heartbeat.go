package handler

import (
	"github.com/ilibs/very/core"
)

var Check core.HandlerFunc = func(c *core.Context) core.Response {
	return c.String("ok")
}