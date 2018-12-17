package cmd

import (
	"app/router"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ilibs/very/command"
	"github.com/urfave/cli"
	"github.com/verystar/golib/osutil"
)

var HTTPCmd = cli.Command{
	Name:  "http",
	Usage: "http command eg: ./app http --addr=:8080",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Usage: "http listen ip:port",
		},
	},
	Action: func(ctx *cli.Context) error {
		if !ctx.IsSet("addr") {
			ctx.Set("addr", ":8081")
		}

		serv := gin.Default()
		srv := &http.Server{
			Addr:    ctx.String("addr"),
			Handler: serv,
		}

		//router
		router.Route(serv)

		osutil.RegisterShutDown(func(sig os.Signal) {
			ctxw, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			srv.Close()
			if err := srv.Shutdown(ctxw); err != nil {
				log.Fatal("HTTP Server Shutdown:", err)
			}
			log.Println("HTTP Server exiting")
		})

		log.Printf("HTTP listen: %s\n", ctx.String("addr"))

		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP listen: %s\n", err)
		}

		return nil
	},
}

func init() {
	command.Register(HTTPCmd)
}
