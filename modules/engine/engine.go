package engine

import (
	"github.com/MexChina/Treasure/modules/config"
	"github.com/MexChina/Treasure/application/admin/view/types"
	"github.com/MexChina/Treasure/modules/logger"
)

type Engine struct {
	ApplicationList []Application
	Adapter    WebFrameWork
}

// Default return the default engine instance.
func Default() *Engine {
	return &Engine{
		Adapter: DefaultAdapter,
	}
}

// Use enable the adapter.
func (eng *Engine) Use(router interface{}) error {
	if eng.Adapter == nil {
		logger.Painc("adapter is nil, import the default adapter or use AddAdapter method add the adapter")
	}
	return eng.Adapter.Use(router, eng.ApplicationList)
}

//addLogger set the global logger
func (eng *Engine) AddLogger(cfg logger.LogConfig) *Engine{
	logger.SetLogger(cfg)
	return eng
}

// AddPlugins add the plugins and initialize them.
func (eng *Engine) AddApplication(apps ...Application) *Engine {
	for _, app := range apps {
		app.InitApplication()
	}
	eng.ApplicationList = append(eng.ApplicationList, apps...)
	return eng
}

// AddConfig set the global config.
func (eng *Engine) AddConfig(cfg config.Config) *Engine {
	config.Set(cfg)
	return eng
}

// AddAdapter add the adapter of engine.
func (eng *Engine) AddAdapter(ada WebFrameWork) *Engine {
	eng.Adapter = ada
	return eng
}

var DefaultAdapter WebFrameWork

func Register(ada WebFrameWork)  {
	if ada == nil {
		panic("adapter is nil")
	}
	DefaultAdapter = ada
}

func Content(ctx interface{}, panel types.GetPanel) {
	if DefaultAdapter == nil {
		panic("adapter is nil")
	}
	DefaultAdapter.Content(ctx, panel)
}