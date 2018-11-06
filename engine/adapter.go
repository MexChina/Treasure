package engine

import (
	"github.com/MexChina/Treasure/plugins"
	"github.com/MexChina/Treasure/template/types"
)

type WebFrameWork interface {
	Use(interface{}, []plugins.Plugin) error
	Content(interface{}, types.GetPanel)
}
