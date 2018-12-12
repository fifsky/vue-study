package handler

import (
	"app/model"
	"github.com/gin-gonic/gin"
	"github.com/ilibs/gosql"
	"github.com/ilibs/very/core"
	"github.com/verystar/golib/pagination"
	"github.com/verystar/logger"
)

var AdminMoodList core.HandlerFunc = func(c *core.Context) core.Response {
	p := &struct {
		Page int `json:"page"`
	}{}
	if err := c.ShouldBindJSON(p); err != nil {
		return c.Fail(201, "参数错误:"+err.Error())
	}

	num := 10
	moods, err := model.MoodGetList(p.Page, num)
	if err != nil {
		return c.Fail(202, err)
	}

	total, err := gosql.Model(&model.Moods{}).Count()
	if err != nil {
		return c.Fail(203, err)
	}

	pager := pagination.New(int(total), num, p.Page, 3)

	return c.Success(gin.H{
		"list": moods,
		"page": gin.H{
			"total":   pager.Total(),
			"num":     pager.PagingNum(),
			"current": pager.Current(),
		},
	})
}

var AdminMoodPost core.HandlerFunc = func(c *core.Context) core.Response {
	moods := &model.Moods{}
	if err := c.ShouldBindJSON(moods); err != nil {
		return c.Fail(201, "参数错误:"+err.Error())
	}

	moods.UserId = getLoginUser(c).Id

	if moods.Content == "" {
		return c.Fail(201, "内容不能为空")
	}

	id, err := gosql.Model(moods).Create();
	if err != nil {
		return c.Fail(201, "发表失败")
	}

	userMood := &model.UserMoods{}
	err = gosql.Model(userMood).Where("id = ?", id).Get()
	if err != nil {
		return c.Fail(201, "发表失败")
	}

	return c.Success(userMood)
}

var AdminMoodDelete core.HandlerFunc = func(c *core.Context) core.Response {
	mood := &model.Moods{}
	if err := c.ShouldBindJSON(mood); err != nil {
		return c.Fail(201, "参数错误:"+err.Error())
	}

	if _, err := gosql.Model(mood).Delete(); err != nil {
		logger.Error(err)
		return c.Fail(201, "删除失败")
	}
	return c.Success(nil)
}
