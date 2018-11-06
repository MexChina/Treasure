package page

import (
	"bytes"
	"github.com/MexChina/Treasure/context"
	"github.com/MexChina/Treasure/modules/auth"
	"github.com/MexChina/Treasure/modules/config"
	"github.com/MexChina/Treasure/modules/menu"
	"github.com/MexChina/Treasure/template"
	"github.com/MexChina/Treasure/template/types"
)

// SetPageContent set and return the panel of page content.
func SetPageContent(ctx *context.Context, c func() types.Panel) {
	user := ctx.UserValue["user"].(auth.User)

	panel := c()

	globalConfig := config.Get()

	tmpl, tmplName := template.Get(globalConfig.THEME).GetTemplate(ctx.Request.Header.Get("X-PJAX") == "true")

	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: menu.GetGlobalMenu(user),
		System: types.SystemInfo{
			"0.0.1",
		},
		Panel:         panel,
		AssertRootUrl: "/" + globalConfig.PREFIX,
		Title:         globalConfig.TITLE,
		Logo:          globalConfig.LOGO,
		MiniLogo:      globalConfig.MINILOGO,
	})
	ctx.WriteString(buf.String())

}