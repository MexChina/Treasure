package main

import (
	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
	"github.com/MexChina/Treasure/engine"
	"github.com/MexChina/Treasure/modules/config"
	"github.com/MexChina/Treasure/plugins/admin"
	"github.com/MexChina/Treasure/modules/logger"
)

func main() {
	cfg := config.Config{
		DATABASE: []config.Database{
			{
				HOST:         "192.168.1.201",
				PORT:         "3306",
				USER:         "devuser",
				PWD:          "devuser",
				NAME:         "a_admin",
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       "mysql",
			},
		},
		DOMAIN:   "localhost",
		PREFIX:   "admin",
		INDEX:    "/",
		THEME:    "",
		TITLE:    "Treasure",
		LANGUAGE: "cn",
		STORE: config.Store{
			PATH:   "./uploads",
			PREFIX: "uploads",
		},
		LOGGER: logger.LogConfig{
			Console:&logger.ConsoleLogger{
				Level:"DEBG",
				Colorful:true,
			},
		},
		VERSION:"0.0.1",
		ASSETS:"src/github.com/MexChina/Treasure/template/adminlte/resource",
	}
	router := fasthttprouter.New()
	eng := engine.Default()
	if err := eng.AddConfig(cfg).AddPlugins(admin.NewAdmin()).AddLogger(cfg.LOGGER).Use(router); err != nil {
		logger.Painc(err)
	}
	logger.Fatal(fasthttp.ListenAndServe(":8897", router.Handler))
}
