package controller

import (
	"bytes"
	"github.com/MexChina/Treasure/modules/context"
	"github.com/MexChina/Treasure/modules/auth"
	"github.com/MexChina/Treasure/modules/menu"
	"net/http"
	"github.com/MexChina/Treasure/application/admin/view"
)

func Auth(ctx *context.Context) {
	password := ctx.Request.FormValue("password")
	username := ctx.Request.FormValue("username")
	if user, ok := auth.Check(password, username); ok {
		auth.SetCookie(ctx, user)
		menu.Unlock()
		ctx.Json(http.StatusOK, map[string]interface{}{
			"code": 200,
			"msg":  "登录成功",
			"url":  Config.PREFIX + Config.INDEX,
		})
		return
	}
	ctx.Json(http.StatusBadRequest, map[string]interface{}{
		"code": 400,
		"msg":  "登录失败",
	})
	return
}

func Logout(ctx *context.Context) {
	auth.DelCookie(ctx)
	ctx.Response.Header.Add("Location", Config.PREFIX+"/login")
	ctx.SetStatusCode(302)
}

func ShowLogin(ctx *context.Context) {
	defer GlobalDeferHandler(ctx)
	buf := new(bytes.Buffer)
	tmpler := view.Adminlte.GetHtml("login")
	tmpler.ExecuteTemplate(buf,"login", struct {
		AssertRootUrl string
	}{Config.PREFIX})
	ctx.WriteString(buf.String())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}
