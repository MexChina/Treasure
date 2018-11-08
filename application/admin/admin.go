package admin

import (
	"github.com/MexChina/Treasure/modules/context"
	"github.com/MexChina/Treasure/modules/config"
	"github.com/MexChina/Treasure/application"
	"github.com/MexChina/Treasure/application/admin/controller"
	"github.com/MexChina/Treasure/application/admin/models"
	"github.com/MexChina/Treasure/modules/orm"
)

type Admin struct {
	app      *context.App
	tableCfg map[string]models.TableGenerator
}

func (admin *Admin) InitApplication() {

	cfg := config.Get()

	// Init database
	for _, databaseCfg := range cfg.DATABASE {
		orm.GetConnectionByDriver(databaseCfg.DRIVER).InitDB(map[string]config.Database{
			"default": databaseCfg,
		})
	}

	// Init router
	App.app = InitRouter("/" + cfg.PREFIX)

	models.SetGenerators(map[string]models.TableGenerator{
		"manager":    models.GetManagerTable,
		"permission": models.GetPermissionTable,
		"roles":      models.GetRolesTable,
		"op":         models.GetOpTable,
		"menu":       models.GetMenuTable,
	})
	models.SetGenerators(admin.tableCfg)
	models.InitTableList()

	cfg.PREFIX = "/" + cfg.PREFIX
	controller.SetConfig(cfg)

}

var App = new(Admin)

//func NewAdmin(tableCfg map[string]models.TableGenerator) *Admin {
func NewAdmin() *Admin {
	//App.tableCfg = tableCfg
	return App
}

func (admin *Admin) GetRequest() []context.Path {
	return admin.app.Requests
}

func (admin *Admin) GetHandler(url, method string) context.Handler {
	return application.GetHandler(url, method, admin.app)
}
