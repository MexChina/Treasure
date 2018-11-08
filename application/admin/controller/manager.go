package controller

import (
	"github.com/MexChina/Treasure/modules/auth"
	"github.com/MexChina/Treasure/modules/orm"
)

func NewManager(dataList map[string][]string) {

	// 更新管理员表
	result := orm.GetConnection().Exec("insert into users (username, password, name, avatar) values (?, ?, ?, ?)",
		dataList["username"][0], auth.EncodePassword([]byte(dataList["password"][0])), dataList["name"][0], dataList["avatar"][0])

	id, _ := result.LastInsertId()

	// 插入管理员角色表
	for i := 0; i < len(dataList["role_id[]"]); i++ {
		if dataList["role_id[]"][i] != "" {
			orm.GetConnection().Exec("insert into role_users (role_id, user_id) values (?, ?)",
				dataList["role_id[]"][i], id)
		}
	}

	// 更新管理员权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			orm.GetConnection().Exec("insert into user_permissions (permission_id, user_id) values (?, ?)",
				dataList["permission_id[]"][i], id)
		}
	}
}

func EditManager(dataList map[string][]string) {

	// 更新管理员表
	orm.GetConnection().Exec("update users set username = ?, password = ?, name = ?, avatar = ? where id = ?",
		dataList["username"][0], auth.EncodePassword([]byte(dataList["password"][0])), dataList["name"][0],
		dataList["avatar"][0], dataList["id"][0])

	// 插入管理员角色表
	for i := 0; i < len(dataList["role_id[]"]); i++ {
		if dataList["role_id[]"][i] != "" {
			checkRole, _ := orm.GetConnection().Query("select * from role_users where role_id = ? and user_id = ?",
				dataList["role_id[]"][i], dataList["id"][0])
			if len(checkRole) < 1 {
				orm.GetConnection().Exec("insert into role_users (role_id, user_id) values (?, ?)",
					dataList["role_id[]"][i], dataList["id"][0])
			}
		}
	}

	// 更新管理员权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			checkPermission, _ := orm.GetConnection().Query("select * from user_permissions where permission_id = ? and user_id = ?",
				dataList["permission_id[]"][i], dataList["id"][0])
			if len(checkPermission) < 1 {
				orm.GetConnection().Exec("insert into user_permissions (permission_id, user_id) values (?, ?)",
					dataList["permission_id[]"][i], dataList["id"][0])
			}
		}
	}
}

func NewRole(dataList map[string][]string) {
	// 更新管理员角色表
	result := orm.GetConnection().Exec("insert into roles (name, slug) values (?, ?)",
		dataList["name"][0], dataList["slug"][0])

	id, _ := result.LastInsertId()

	// 更新管理员角色权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			orm.GetConnection().Exec("insert into role_permissions (permission_id, role_id) values (?, ?)",
				dataList["permission_id[]"][i], id)
		}
	}
}

func EditRole(dataList map[string][]string) {
	// 更新管理员角色表
	orm.GetConnection().Exec("update roles set name = ?, slug = ? where id = ?",
		dataList["name"][0], dataList["slug"][0], dataList["id"][0])

	// 更新管理员角色权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			checkPermission, _ := orm.GetConnection().Query("select * from role_permissions where permission_id = ? and role_id = ?",
				dataList["permission_id[]"][i], dataList["id"][0])
			if len(checkPermission) < 1 {
				orm.GetConnection().Exec("insert into role_permissions (permission_id, role_id) values (?, ?)",
					dataList["permission_id[]"][i], dataList["id"][0])
			}
		}
	}
}
