package cmd

import (
	"fmt"

	"app/model"

	"github.com/goapt/gee"
	"github.com/goapt/logger"
	"github.com/ilibs/gosql"
	"github.com/urfave/cli"
)

var TestCmd = cli.Command{
	Name:  "test",
	Usage: "test command eg: ./app test --id=7",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "id",
			Usage: "user id",
		},
	},
	Action: func(ctx *cli.Context) error {
		if !ctx.IsSet("id") {
			ctx.Set("id", "7")
		}

		user := &model.Users{}
		err := gosql.Model(user).Where("id = ?", ctx.Int("id")).Get()
		if err != nil {
			logger.Error("get user error")
			return err
		}

		fmt.Println(user)

		logger.Error("test", user)
		return nil
	},
}

func init() {
	gee.RegisterCommand(TestCmd)
}
