package engine

import (
	"github.com/MexChina/Treasure/template/types"
	"github.com/MexChina/Treasure/application"
)

type WebFrameWork interface {
	Use(interface{}, []application.Plugin) error
	Content(interface{}, types.GetPanel)
}
