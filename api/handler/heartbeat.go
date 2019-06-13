package handler

import (
	"github.com/goapt/gee"
)

var Check gee.HandlerFunc = func(c *gee.Context) gee.Response {
	return c.String("ok")
}
