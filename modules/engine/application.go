package engine
import (
	"github.com/MexChina/Treasure/modules/context"
)

type Application interface {
	GetRequest() []context.Path
	GetHandler(url, method string) context.Handler
	InitApplication()
}

func GetHandler(url, method string, app *context.App) context.Handler {
	handler := app.Find(url, method)
	if handler == nil {
		panic("handler not found")
	}
	return handler
}