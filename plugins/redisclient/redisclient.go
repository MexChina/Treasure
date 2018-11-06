package redisclient

import (
	"github.com/MexChina/Treasure/context"
	"github.com/MexChina/Treasure/modules/config"
	"github.com/MexChina/Treasure/plugins"
)

type RedisClient struct {
	app *context.App
}

var Plug = new(RedisClient)

var Config config.Config

func (redis *RedisClient) InitPlugin() {
	cfg := config.Get()

	Config = cfg
	Config.PREFIX = "/" + Config.PREFIX

	Plug.app = InitRouter(Config.PREFIX)

}

func NewRedisClient() *RedisClient {
	return Plug
}

func (redis *RedisClient) GetRequest() []context.Path {
	return redis.app.Requests
}

func (redis *RedisClient) GetHandler(url, method string) context.Handler {
	return plugins.GetHandler(url, method, redis.app)
}