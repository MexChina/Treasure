package controller

import (
	"bytes"
	"github.com/MexChina/Treasure/modules/context"
	"github.com/MexChina/Treasure/modules/auth"
	"github.com/MexChina/Treasure/modules/menu"
	"github.com/MexChina/Treasure/application/admin/models"
	"github.com/MexChina/Treasure/template"
	"github.com/MexChina/Treasure/template/types"
	"net/http"
	"path"
	"strings"
	"github.com/MexChina/Treasure/modules/logger"
	"os"
	"io/ioutil"
	"path/filepath"
)

// 显示列表
func ShowInfo(ctx *context.Context) {

	defer GlobalDeferHandler(ctx)

	user := ctx.UserValue["user"].(auth.User)

	prefix := ctx.Request.URL.Query().Get("prefix")

	page := ctx.Request.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	pageSize := ctx.Request.URL.Query().Get("pageSize")
	if pageSize == "" {
		pageSize = "10"
	}

	sortField := ctx.Request.URL.Query().Get("sort")
	if sortField == "" {
		sortField = "id"
	}
	sortType := ctx.Request.URL.Query().Get("sort_type")
	if sortType == "" {
		sortType = "desc"
	}

	thead, infoList, paginator, title, description := models.TableList[prefix].GetDataFromDatabase(map[string]string{
		"page":      page,
		"path":      ctx.Path(),
		"sortField": sortField,
		"sortType":  sortType,
		"prefix":    prefix,
		"pageSize":  pageSize,
	})

	editUrl := Config.PREFIX + "/info/" + prefix + "/edit" + GetRouteParameterString(page, pageSize, sortType, sortField)
	newUrl := Config.PREFIX + "/info/" + prefix + "/new" + GetRouteParameterString(page, pageSize, sortType, sortField)
	deleteUrl := Config.PREFIX + "/delete/" + prefix

	tmpl, tmplName := template.Get("adminlte").GetTemplate(ctx.Request.Header.Get("X-PJAX") == "true")

	menu.GlobalMenu.SetActiveClass(ctx.Path())

	dataTable := template.Get(Config.THEME).DataTable().SetInfoList(infoList).SetThead(thead).SetEditUrl(editUrl).SetNewUrl(newUrl).SetDeleteUrl(deleteUrl)
	table := dataTable.GetContent()

	box := template.Get(Config.THEME).Box().
		SetBody(table).
		SetHeader(dataTable.GetDataTableHeader()).
		WithHeadBorder(false).
		SetFooter(paginator.GetContent()).
		GetContent()

	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: menu.GetGlobalMenu(user),
		System: types.SystemInfo{
			Config.VERSION,
		},
		Panel: types.Panel{
			Content:     box,
			Description: description,
			Title:       title,
		},
		AssertRootUrl: Config.PREFIX,
		Title:         Config.TITLE,
		Logo:          Config.LOGO,
		MiniLogo:      Config.MINILOGO,
	})
	ctx.WriteString(buf.String())
}

func Assert(ctx *context.Context) {
	filepathstr := Config.ASSETS + strings.Replace(ctx.Request.URL.Path, Config.PREFIX, "", 1)
	pathsss, _ := filepath.Abs(filepathstr);
	fin, err := os.Open(pathsss)
	defer fin.Close()
	if err != nil {
		logger.Error("static resource:", err)
	}
	data, err := ioutil.ReadAll(fin)

	fileSuffix := path.Ext(filepathstr)
	fileSuffix = strings.Replace(fileSuffix, ".", "", -1)

	var contentType = ""
	if fileSuffix == "css" || fileSuffix == "js" {
		contentType = "text/" + fileSuffix + "; charset=utf-8"
	} else {
		contentType = "image/" + fileSuffix
	}

	if err != nil {
		logger.Error("asset err", err)
		ctx.Write(http.StatusNotFound, map[string]string{}, "")
	} else {
		ctx.Write(http.StatusOK, map[string]string{
			"content-type": contentType,
		}, string(data))
	}
}

func GetRouteParameterString(page, pageSize, sortType, sortField string) string {
	return "?page=" + page + "&pageSize=" + pageSize + "&sort=" + sortField + "&sort_type=" + sortType
}
