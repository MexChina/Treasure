package datamodel

import (
	"github.com/MexChina/Treasure/plugins/admin/models"
	"github.com/MexChina/Treasure/template/types"
)

func GetUserTable() (userTable models.Table) {

	userTable.Info.FieldList = []types.FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Sortable: true,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Name",
			Field:    "name",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Gender",
			Field:    "gender",
			TypeName: "tinyint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				if model.Value == "1" {
					return "man"
				}
				if model.Value == "2" {
					return "women"
				}
				return "unknown"
			},
		},
		{
			Head:     "Phone",
			Field:    "phone",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "City",
			Field:    "city",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	userTable.Info.Table = "users"
	userTable.Info.Title = "Users"
	userTable.Info.Description = "Users"

	userTable.Form.FormList = []types.FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Ip",
			Field:    "ip",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Name",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Gender",
			Field:    "gender",
			TypeName: "tinyint",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Phone",
			Field:    "phone",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "City",
			Field:    "city",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	userTable.Form.Table = "users"
	userTable.Form.Title = "Users"
	userTable.Form.Description = "Users"

	userTable.ConnectionDriver = "mysql"

	return
}
