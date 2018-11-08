package application
import (
	"github.com/MexChina/Treasure/context"
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