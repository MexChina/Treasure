package admin

import (
	"github.com/MexChina/Treasure/context"
	"github.com/MexChina/Treasure/modules/auth"
	"github.com/MexChina/Treasure/plugins/admin/controller"
	"github.com/MexChina/Treasure/template"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	app.Group(prefix)
	{
		// auth
		app.GET("/login", controller.ShowLogin)
		app.POST("/signin", controller.Auth)

		for _, path := range template.Get("adminlte").GetAssetList() {
			app.GET("/assets"+path, controller.Assert)
		}

		for _, path := range template.GetComp("login").GetAssetList() {
			app.GET("/assets"+path, controller.Assert)
		}

		authenticator := auth.SetPrefix(prefix).SetAuthFailCallback(func(ctx *context.Context) {
			ctx.Write(302, map[string]string{
				"Location": prefix + "/login",
			}, ``)
		}).SetPermissionDenyCallback(func(ctx *context.Context) {
			controller.ShowErrorPage(ctx, "permission denied")
		})

		app.Group("", authenticator.Middleware)
		{
			// auth
			app.GET("/logout", controller.Logout)

			// menus
			app.GET("/menu", controller.ShowMenu)
			app.POST("/menu/delete", controller.DeleteMenu)
			app.POST("/menu/new", controller.NewMenu)
			//app.GET("/menu/new", controller.ShowMenu) // TODO: this is a bug of the tire
			app.POST("/menu/edit", controller.EditMenu)
			app.GET("/menu/edit/show", controller.ShowEditMenu)
			app.POST("/menu/order", controller.MenuOrder)

			// add delete modify query
			app.GET("/info/:prefix", controller.ShowInfo)
			app.GET("/info/:prefix/edit", controller.ShowForm)
			app.GET("/info/:prefix/new", controller.ShowNewForm)
			app.POST("/edit/:prefix", controller.EditForm)
			app.POST("/delete/:prefix", controller.DeleteData)
			app.POST("/new/:prefix", controller.NewForm)
		}
	}

	return app
}
