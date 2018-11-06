package connections

import (
	"database/sql"
	"github.com/MexChina/Treasure/modules/config"
	"github.com/MexChina/Treasure/modules/connections/mysql"
	"github.com/MexChina/Treasure/modules/connections/postgresql"
	"github.com/MexChina/Treasure/modules/connections/sqlite"
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
