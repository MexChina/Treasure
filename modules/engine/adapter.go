package engine

import (
	"github.com/MexChina/Treasure/application/admin/view/types"
)

type WebFrameWork interface {
	Use(interface{}, []Application) error
	Content(interface{}, types.GetPanel)
}
