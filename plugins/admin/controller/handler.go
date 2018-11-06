package controller

import (
	"bytes"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/mgutz/ansi"
	"github.com/MexChina/Treasure/context"
	"github.com/MexChina/Treasure/modules/auth"
	"github.com/MexChina/Treasure/modules/menu"
	"github.com/MexChina/Treasure/plugins/admin/models"
	"github.com/MexChina/Treasure/template"
	"github.com/MexChina/Treasure/template/types"
	template2 "html/template"
	"log"
	"net/http"
	"regexp"
	"runtime/debug"
	"strconv"
)

// 全局错误处理
func GlobalDeferHandler(ctx *context.Context) {

	log.Println("[GoAdmin]",
		ansi.Color(" "+strconv.Itoa(ctx.Response.StatusCode)+" ", "white:blue"),
		ansi.Color(" "+string(ctx.Method()[:])+"   ", "white:blue+h"),
		ctx.Path())

	RecordOperationLog(ctx)

	if err := recover(); err != nil {
		fmt.Println(err)
		fmt.Println(string(debug.Stack()[:]))

		var (
			errMsg     string
			mysqlError *mysql.MySQLError
			ok         bool
		)
		if errMsg, ok = err.(string); !ok {
			if mysqlError, ok = err.(*mysql.MySQLError); ok {
				errMsg = mysqlError.Error()
			} else {
				errMsg = fmt.Sprint("%v", err)
			}
		}

		alert := template.Get(Config.THEME).Alert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> Error!`)).
			SetTheme("warning").SetContent(template2.HTML(errMsg)).GetContent()

		if ok, _ = regexp.Match("/edit(.*)", []byte(ctx.Path())); ok {

			user := ctx.UserValue["user"].(auth.User)

			prefix := ctx.Request.URL.Query().Get("prefix")

			id := ctx.Request.URL.Query().Get("id")

			formData, title, description := models.TableList[prefix].GetDataFromDatabaseWithId(prefix, id)

			tmpl, tmplName := template.Get("adminlte").GetTemplate(ctx.Request.Header.Get("X-PJAX") == "true")

			path := ctx.Path()
			menu.GlobalMenu.SetActiveClass(path)

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

			ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")

			queryParam := "?page=" + page + "&pageSize=" + pageSize + "&sort=" + sortField + "&sort_type=" + sortType

			buf := new(bytes.Buffer)
			tmpl.ExecuteTemplate(buf, tmplName, types.Page{
				User: user,
				Menu: menu.GetGlobalMenu(user),
				System: types.SystemInfo{
					"0.0.1",
				},
				Panel: types.Panel{
					Content: alert + template.Get(Config.THEME).Form().
						SetContent(formData).
						SetPrefix(Config.PREFIX).
						SetUrl(Config.PREFIX+"/edit/"+prefix).
						SetToken(auth.TokenHelper.AddToken()).
						SetInfoUrl(Config.PREFIX+"/info/"+prefix+queryParam).
						GetContent(),
					Description: description,
					Title:       title,
				},
				AssertRootUrl: Config.PREFIX,
				Title:         Config.TITLE,
				Logo:          Config.LOGO,
				MiniLogo:      Config.MINILOGO,
			})
			ctx.WriteString(buf.String())
			ctx.Response.Header.Add("X-PJAX-URL", Config.PREFIX+"/info/"+prefix+"/new"+queryParam)
			return
		}

		if ok, _ = regexp.Match("/new(.*)", []byte(ctx.Path())); ok {
			prefix := ctx.Request.URL.Query().Get("prefix")

			user := ctx.UserValue["user"].(auth.User)

			tmpl, tmplName := template.Get("adminlte").GetTemplate(ctx.Request.Header.Get("X-PJAX") == "true")

			path := ctx.Path()
			menu.GlobalMenu.SetActiveClass(path)

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

			ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")

			queryParam := "?page=" + page + "&pageSize=" + pageSize + "&sort=" + sortField + "&sort_type=" + sortType

			buf := new(bytes.Buffer)
			tmpl.ExecuteTemplate(buf, tmplName, types.Page{
				User: user,
				Menu: menu.GetGlobalMenu(user),
				System: types.SystemInfo{
					"0.0.1",
				},
				Panel: types.Panel{
					Content: alert + template.Get(Config.THEME).Form().
						SetPrefix(Config.PREFIX).
						SetContent(models.GetNewFormList(models.TableList[prefix].Form.FormList)).
						SetUrl(Config.PREFIX+"/new/"+prefix).
						SetToken(auth.TokenHelper.AddToken()).
						SetInfoUrl(Config.PREFIX+"/info/"+prefix+queryParam).
						GetContent(),
					Description: models.TableList[prefix].Form.Description,
					Title:       models.TableList[prefix].Form.Title,
				},
				AssertRootUrl: Config.PREFIX,
				Title:         Config.TITLE,
				Logo:          Config.LOGO,
				MiniLogo:      Config.MINILOGO,
			})
			ctx.WriteString(buf.String())
			ctx.Response.Header.Add("X-PJAX-URL", Config.PREFIX+"/info/"+prefix+"/new"+queryParam)
			return
		}

		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetContentType("application/json")
		ctx.WriteString(`{"code":500, "msg":"` + errMsg + `"}`)
		return
	}
}
