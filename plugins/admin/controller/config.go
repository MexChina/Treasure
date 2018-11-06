package controller

import "github.com/MexChina/Treasure/modules/config"

var Config config.Config

func SetConfig(cfg config.Config) {
	Config = cfg
}
