package controller

import (
	"github.com/MexChina/Treasure/modules/context"
	"github.com/MexChina/Treasure/modules/auth"
	"github.com/MexChina/Treasure/application/admin/models"
)

func DeleteData(ctx *context.Context) {
	defer GlobalDeferHandler(ctx)
	prefix := ctx.Request.URL.Query().Get("prefix")
	id := ctx.Request.FormValue("id")
	models.TableList[prefix].DeleteDataFromDatabase(prefix, id)
	newToken := auth.TokenHelper.AddToken()
	ctx.WriteString(`{"code":200, "msg":"删除成功", "data":"` + newToken + `"}`)
	return
}
