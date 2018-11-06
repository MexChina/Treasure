package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/MexChina/Treasure/modules/config"
	"github.com/MexChina/Treasure/modules/orm/performer"
	"sync"
)

type Sqlite struct {
	SqlDBmap map[string]*sql.DB
	Once     sync.Once
}

func GetSqliteDB() *Sqlite {
	return &DB
}

func (db *Sqlite) Query(query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows) {
	return performer.Query(db.SqlDBmap["default"], query, args...)
}

func (db *Sqlite) Exec(query string, args ...interface{}) sql.Result {
	return performer.Exec(db.SqlDBmap["default"], query, args...)
}

func (db *Sqlite) InitDB(cfgList map[string]config.Database) {
	db.Once.Do(func() {
		var (
			sqlDB *sql.DB
			err   error
		)

		for conn, cfg := range cfgList {
			sqlDB, err = sql.Open("sqlite3", cfg.FILE)

			if err != nil {
				panic(err)
			} else {
				db.SqlDBmap[conn] = sqlDB
			}
		}
	})
}

var DB = Sqlite{
	SqlDBmap: map[string]*sql.DB{},
}
