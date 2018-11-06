package controller

import (
	"encoding/json"
	"github.com/MexChina/Treasure/context"
	"github.com/MexChina/Treasure/modules/auth"
	"github.com/MexChina/Treasure/modules/orm"
)

func RecordOperationLog(ctx *context.Context) {
	if user, ok := ctx.UserValue["user"].(auth.User); ok {
		var input []byte
		form := ctx.Request.MultipartForm
		if form != nil {
			input, _ = json.Marshal((*form).Value)
		}

		orm.GetConnection().Exec("insert into goadmin_operation_log (user_id, path, method, ip, input) values (?, ?, ?, ?, ?)", user.ID, ctx.Path(),
			ctx.Method(), ctx.LocalIP(), string(input))
	}
}
