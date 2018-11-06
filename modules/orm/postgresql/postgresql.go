package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/MexChina/Treasure/modules/config"
	"github.com/MexChina/Treasure/modules/orm/performer"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

type Postgresql struct {
	SqlDBmap map[string]*sql.DB
	Once     sync.Once
}

func GetPostgresqlDB() *Postgresql {
	return &DB
}

func (db *Postgresql) Query(query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows) {
	queCount := strings.Count(query, "?")
	for i := 1; i < queCount+1; i++ {
		query = strings.Replace(query, "?", "$"+strconv.Itoa(i), 1)
	}
	query = strings.Replace(query, "`", "", -1)
	// TODO: 关键字加双引号
	query = strings.Replace(query, "by order ", `by "order" `, -1)
	fmt.Println("query", query)
	return performer.Query(db.SqlDBmap["default"], query, args...)
}

func (db *Postgresql) Exec(query string, args ...interface{}) sql.Result {
	queCount := strings.Count(query, "?")
	for i := 1; i < queCount+1; i++ {
		query = strings.Replace(query, "?", "$"+strconv.Itoa(i), 1)
	}
	query = strings.Replace(query, "`", "", -1)
	// TODO: 关键字加双引号
	query = strings.Replace(query, "by order ", `by "order" `, -1)
	fmt.Println("query", query)
	return performer.Exec(db.SqlDBmap["default"], query, args...)
}

func (db *Postgresql) InitDB(cfgList map[string]config.Database) {
	db.Once.Do(func() {
		var (
			sqlDB *sql.DB
			err   error
		)

		for conn, cfg := range cfgList {

			query := url.Values{}
			query.Add("sslmode", "disable") // "verify-full"

			u := &url.URL{
				Scheme:   "postgres",
				User:     url.UserPassword(cfg.USER, cfg.PWD),
				Host:     fmt.Sprintf("%s:%s", cfg.HOST, cfg.PORT),
				RawQuery: query.Encode(),
			}

			fmt.Println("connection: ", u.String())

			sqlDB, err = sql.Open("postgres", u.String())

			if err != nil {
				panic(err)
			} else {
				db.SqlDBmap[conn] = sqlDB
			}
		}
	})
}

var DB = Postgresql{
	SqlDBmap: map[string]*sql.DB{},
}
