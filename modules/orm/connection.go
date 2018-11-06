package orm

import (
	"database/sql"
	"github.com/MexChina/Treasure/modules/config"
	"github.com/MexChina/Treasure/modules/orm/mysql"
	"github.com/MexChina/Treasure/modules/orm/postgresql"
	"github.com/MexChina/Treasure/modules/orm/sqlite"
)

type Connection interface {
	Query(query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows)
	Exec(query string, args ...interface{}) sql.Result
	InitDB(cfg map[string]config.Database)
}

func GetConnectionByDriver(driver string) Connection {
	switch driver {
	case "mysql":
		return mysql.GetMysqlDB()
	case "sqlite":
		return sqlite.GetSqliteDB()
	case "postgresql":
		return postgresql.GetPostgresqlDB()
	default:
		panic("driver not found!")
	}
}

func GetConnection() Connection {
	return GetConnectionByDriver(config.Get().DATABASE[0].DRIVER)
}
