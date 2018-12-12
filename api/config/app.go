package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/fifsky/goconf"
	"github.com/ilibs/gosql"
	"github.com/ilibs/very/server"
	"github.com/verystar/golib/debug"
	"github.com/verystar/logger"
)

type common struct {
	Env         string `json:"env"`
	ConfigPath  string `json:"config_path"`
	StoragePath string `json:"storage_path"`
	Debug       string `json:"debug"`
	TokenSecret string `json:"token_secret"`
}

type app struct {
	Common    common                   `conf:"common"`
	Log       logger.Config            `conf:"log"`
	DB        map[string]*gosql.Config `conf:"database"`
	StartTime time.Time
}

var App = &app{
	StartTime: time.Now(),
}

func init() {
	server.ArgsInit()
	Load(server.ExtArgs)
}

func Load(args map[string]string) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "prod"
	}

	if args["env"] != "" {
		env = args["env"]
	}

	appPath := args["config"]
	if appPath == "" {
		_, file, _, _ := runtime.Caller(0)
		appPath = filepath.Dir(filepath.Dir(file))
	}

	App.Common.ConfigPath = filepath.Join(appPath, "config")

	conf, err := goconf.NewConfig(filepath.Join(App.Common.ConfigPath, env))
	if err != nil {
		logger.Fatalf("json config path error %s", err.Error())
	}

	//load config
	if err := conf.Load(App); err != nil {
		log.Fatal("Config Error:", err)
	}

	if !filepath.IsAbs(App.Common.StoragePath) {
		App.Common.StoragePath = filepath.Join(appPath, App.Common.StoragePath)
	}

	//debug model
	if args["debug"] != "" {
		App.Common.Debug = args["debug"]
	}

	//debug
	if App.Common.Debug == "on" {
		debug.Open("on", args["debug-tag"])
		debug.SavePath(filepath.Join(App.Common.StoragePath, "debug"))
		//log level
		App.Log.LogLevel = "debug"
		//log model
		App.Log.LogMode = "std"
	}

	if args["show-sql"] == "on" {
		for _, d := range App.DB {
			d.ShowSql = true
		}
	} else if args["show-sql"] == "off" {
		for _, d := range App.DB {
			d.ShowSql = false
		}
	}

	logger.Setting(func(c *logger.Config) {
		c.LogName = App.Log.LogName
		c.LogMode = App.Log.LogMode
		c.LogLevel = App.Log.LogLevel
		c.LogMaxFiles = App.Log.LogMaxFiles
		c.LogPath = filepath.Join(App.Common.StoragePath, "logs")
		c.LogSentryDSN = App.Log.LogSentryDSN
		c.LogSentryType = App.Log.LogSentryType
		c.LogDetail = App.Log.LogDetail
	})
}
