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
	router := fasthttprouter.New()
	eng := engine.Default()
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
		DOMAIN: "localhost",
		PREFIX: "admin",
		INDEX:  "/",
		THEME: "",
		TITLE: "Treasure",
		LANGUAGE: "cn",
		STORE: config.Store{
			PATH:   "./uploads",
			PREFIX: "uploads",
		},
	}

	logger.SetLogger(`{"Console": {"level": "DEBG,TRAC","color":true}}`)
	if err := eng.AddConfig(cfg).AddPlugins(admin.NewAdmin()).Use(router); err != nil {
		panic(err)
	}
	logger.Fatal(fasthttp.ListenAndServe(":8897", router.Handler))
}
