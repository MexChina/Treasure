package engine

import (
	"bytes"
	"errors"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/MexChina/Treasure/modules/context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"github.com/MexChina/Treasure/application/admin/view/types"
	"github.com/MexChina/Treasure/modules/config"
	"github.com/MexChina/Treasure/modules/auth"
	"github.com/MexChina/Treasure/template"
	template2 "html/template"
	"github.com/MexChina/Treasure/modules/menu"
)

type Fasthttp struct {
}

func init()  {
	Register(new(Fasthttp))
}

func (fast *Fasthttp) Use(router interface{}, apps []Application) error {
	var (
		eng *fasthttprouter.Router
		ok     bool
	)
	if eng, ok = router.(*fasthttprouter.Router); !ok {
		return errors.New("wrong parameter")
	}

	for _, app := range apps {
		var plugCopy Application
		plugCopy = app
		for _, req := range app.GetRequest() {
			eng.Handle(strings.ToUpper(req.Method), req.URL, func(c *fasthttp.RequestCtx) {
				httpreq := Convertor(c)
				ctx := context.NewContext(httpreq)

				var params = make(map[string]string, 0)
				c.VisitUserValues(func(i []byte, i2 interface{}) {
					if value, ok := i2.(string); ok {
						params[string(i)] = value
					}
				})

				for key, value := range params {
					if httpreq.URL.RawQuery == "" {
						httpreq.URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + value
					} else {
						httpreq.URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + value
					}
				}

				plugCopy.GetHandler(string(c.Path()), strings.ToLower(string(c.Method())))(ctx)
				for key, head := range ctx.Response.Header {
					c.Response.Header.Set(key, head[0])
				}
				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					buf.ReadFrom(ctx.Response.Body)
					c.WriteString(buf.String())
				}
				c.Response.SetStatusCode(ctx.Response.StatusCode)
			})
		}
	}

	return nil
}

func Convertor(ctx *fasthttp.RequestCtx) *http.Request {
	var r http.Request

	body := ctx.PostBody()
	r.Method = string(ctx.Method())
	r.Proto = "HTTP/1.1"
	r.ProtoMajor = 1
	r.ProtoMinor = 1
	r.RequestURI = string(ctx.RequestURI())
	r.ContentLength = int64(len(body))
	r.Host = string(ctx.Host())
	r.RemoteAddr = ctx.RemoteAddr().String()

	hdr := make(http.Header)
	ctx.Request.Header.VisitAll(func(k, v []byte) {
		sk := string(k)
		sv := string(v)
		switch sk {
		case "Transfer-Encoding":
			r.TransferEncoding = append(r.TransferEncoding, sv)
		default:
			hdr.Set(sk, sv)
		}
	})
	r.Header = hdr
	r.Body = &netHTTPBody{body}
	rURL, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		ctx.Logger().Printf("cannot parse requestURI %q: %s", r.RequestURI, err)
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return &r
	}
	r.URL = rURL
	return &r
}

type netHTTPBody struct {
	b []byte
}

func (r *netHTTPBody) Read(p []byte) (int, error) {
	if len(r.b) == 0 {
		return 0, io.EOF
	}
	n := copy(p, r.b)
	r.b = r.b[n:]
	return n, nil
}

func (r *netHTTPBody) Close() error {
	r.b = r.b[:0]
	return nil
}


func (fast *Fasthttp) Content(contextInterface interface{}, c types.GetPanel) {

	var (
		ctx *fasthttp.RequestCtx
		ok  bool
	)
	if ctx, ok = contextInterface.(*fasthttp.RequestCtx); !ok {
		panic("wrong parameter")
	}

	globalConfig := config.Get()

	sesKey := string(ctx.Request.Header.Cookie("go_admin_session"))

	if sesKey == "" {
		ctx.Redirect("/"+globalConfig.PREFIX+"/login", http.StatusFound)
		return
	}

	userId, ok := auth.Driver.Load(sesKey)["user_id"]

	if !ok {
		ctx.Redirect("/"+globalConfig.PREFIX+"/login", http.StatusFound)
		return
	}

	user, ok := auth.GetCurUserById(userId.(string))

	if !ok {
		ctx.Redirect("/"+globalConfig.PREFIX+"/login", http.StatusFound)
		return
	}

	var panel types.Panel

	if !auth.CheckPermissions(user, string(ctx.Path()), string(ctx.Method())) {
		alert := template.Get(globalConfig.THEME).Alert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> Error!`)).
			SetTheme("warning").SetContent(template2.HTML("没有权限")).GetContent()

		panel = types.Panel{
			Content:     alert,
			Description: "Error",
			Title:       "Error",
		}
	} else {
		panel = c()
	}

	tmpl, tmplName := template.Get(globalConfig.THEME).GetTemplate(string(ctx.Request.Header.Peek("X-PJAX")) == "true")

	ctx.Request.Header.Add("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: menu.GetGlobalMenu(user),
		System: types.SystemInfo{
			globalConfig.VERSION,
		},
		Panel:         panel,
		AssertRootUrl: "/" + globalConfig.PREFIX,
		Title:         globalConfig.TITLE,
		Logo:          globalConfig.LOGO,
		MiniLogo:      globalConfig.MINILOGO,
	})
	ctx.WriteString(buf.String())
}