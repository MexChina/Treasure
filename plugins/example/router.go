package example

import (
	"github.com/MexChina/Treasure/context"
	"github.com/MexChina/Treasure/modules/auth"
	"github.com/MexChina/Treasure/plugins/admin/controller"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	authenticator := auth.SetPrefix(prefix).SetAuthFailCallback(func(ctx *context.Context) {
		ctx.Write(302, map[string]string{
			"Location": prefix + "/login",
		}, ``)
	}).SetPermissionDenyCallback(func(ctx *context.Context) {
		controller.ShowErrorPage(ctx, "permission denied")
	})

	app.GET(prefix+"/example", authenticator.Middleware(TestHandler))

	if prefix == "" {
		app.GET(prefix+"/", authenticator.Middleware(TestHandler))
	} else {
		app.GET(prefix, authenticator.Middleware(TestHandler))
	}

	return app
}
