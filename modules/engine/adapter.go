package engine

import (
	"github.com/MexChina/Treasure/template/types"
)

type WebFrameWork interface {
	Use(interface{}, []Application) error
	Content(interface{}, types.GetPanel)
}
