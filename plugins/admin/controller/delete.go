package controller

import (
	"github.com/MexChina/Treasure/context"
	"github.com/MexChina/Treasure/modules/auth"
	"github.com/MexChina/Treasure/plugins/admin/models"
)

func DeleteData(ctx *context.Context) {

	defer GlobalDeferHandler(ctx)

	//token := ctx.Request.FormValue("_t")
	//
	//if !auth.TokenHelper.CheckToken(token) {
	//	ctx.SetStatusCode(http.StatusBadRequest)
	//	ctx.WriteString(`{"code":400, "msg":"删除失败"}`)
	//	return
	//}

	prefix := ctx.Request.URL.Query().Get("prefix")

	id := ctx.Request.FormValue("id")

	models.TableList[prefix].DeleteDataFromDatabase(prefix, id)

	newToken := auth.TokenHelper.AddToken()

	ctx.WriteString(`{"code":200, "msg":"删除成功", "data":"` + newToken + `"}`)
	return
}
