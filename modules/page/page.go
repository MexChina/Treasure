package page

import (
	"bytes"
	"github.com/MexChina/Treasure/modules/context"
	"github.com/MexChina/Treasure/modules/auth"
	"github.com/MexChina/Treasure/modules/config"
	"github.com/MexChina/Treasure/modules/menu"
	"github.com/MexChina/Treasure/application/admin/view"
	"github.com/MexChina/Treasure/application/admin/view/types"
)

// SetPageContent set and return the panel of page content.
func SetPageContent(ctx *context.Context, c func() types.Panel) {
	user := ctx.UserValue["user"].(auth.User)
	panel := c()
	globalConfig := config.Get()
	tmpl, tmplName := view.Get(globalConfig.THEME).GetTemplate(ctx.Request.Header.Get("X-PJAX") == "true")
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: menu.GetGlobalMenu(user),
		System: types.SystemInfo{
			globalConfig.VERSION,
		},
		Panel:         panel,
		AssertRootUrl: "/" + globalConfig.PREFIX,
		Title:         globalConfig.TITLE,
		Logo:          globalConfig.LOGO,
		MiniLogo:      globalConfig.MINILOGO,
	})
	ctx.WriteString(buf.String())
}
