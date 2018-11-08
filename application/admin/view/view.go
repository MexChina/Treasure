package view

import (
	"github.com/MexChina/Treasure/modules/language"
	"github.com/MexChina/Treasure/modules/menu"
	"github.com/MexChina/Treasure/modules/config"
	"github.com/MexChina/Treasure/modules/logger"
	"github.com/MexChina/Treasure/application/admin/view/components"
	"github.com/MexChina/Treasure/application/admin/view/types"
	"html/template"
	"path/filepath"
	"os"
	"io/ioutil"
)

type Theme struct {
	Name string
}

var Adminlte = Theme{
	Name: "adminlte",
}

func GetAdminlte() *Theme {
	return &Adminlte
}

func (*Theme) GetTemplate(isPjax bool) (tmpler *template.Template, name string) {
	var (
		err error
	)

	if !isPjax {
		name = "layout"
		tmpler, err = template.New("layout").Funcs(template.FuncMap{
			"lang":     language.Get,
			"langHtml": language.GetFromHtml,
		}).Parse(GetHtmlStr("layout") + GetHtmlStr("head") + GetHtmlStr("header") + GetHtmlStr("sidebar") + GetHtmlStr("footer") + GetHtmlStr("js") + GetHtmlStr("menu") + GetHtmlStr("admin_panel") + GetHtmlStr("content"))
	} else {
		name = "content"
		tmpler, err = template.New("content").Funcs(template.FuncMap{
			"lang":     language.Get,
			"langHtml": language.GetFromHtml,
		}).Parse(GetHtmlStr("admin_panel") + GetHtmlStr("content"))
	}

	if err != nil {
		logger.Error(err)
	}

	return
}

func (*Theme) GetHtml(name string) (tmpler *template.Template) {
	cfg := config.Get()
	pathsss, _ := filepath.Abs(cfg.ASSETS+"/application/admin/view/html/"+name+".html");
	tmpler,err := template.ParseFiles(pathsss)
	if err != nil{
		logger.Error(err)
	}
	return
}

func GetHtmlStr(name string) string {
	cfg := config.Get()
	pathsss, _ := filepath.Abs(cfg.ASSETS+"/application/admin/view/html/"+name+".html");
	fin, err := os.Open(pathsss)
	defer fin.Close()
	if err != nil {
		logger.Error("static resource:", err)
		return ""
	}
	data, err := ioutil.ReadAll(fin)
	if err != nil{
		logger.Error(err)
		return ""
	}
	return string(data)
}

func (*Theme) GetAssetList(name string) []string {
	return asserts[name]
}

func (*Theme) Box() types.BoxAttribute {
	return &components.BoxAttribute{
		Name:       "box",
		Header:     template.HTML(""),
		Body:       template.HTML(""),
		Footer:     template.HTML(""),
		Title:      "",
		HeadBorder: "",
	}
}

func (*Theme) Col() types.ColAttribute {
	return &components.ColAttribute{
		Name:    "col",
		Size:    "col-md-2",
		Content: "",
	}
}

func (*Theme) Form() types.FormAttribute {
	return &components.FormAttribute{
		Name:    "form",
		Content: []types.FormStruct{},
		Url:     "/",
		Method:  "post",
		InfoUrl: "",
		Title:   "edit",
	}
}

func (*Theme) Image() types.ImgAttribute {
	return &components.ImgAttribute{
		Name:   "image",
		Witdh:  "50",
		Height: "50",
		Src:    "",
	}
}

func (*Theme) SmallBox() types.SmallBoxAttribute {
	return &components.SmallBoxAttribute{
		Name:  "smallbox",
		Title: "标题",
		Value: "值",
		Url:   "/",
		Color: "aqua",
	}
}

func (*Theme) InfoBox() types.InfoBoxAttribute {
	return &components.InfoBoxAttribute{
		Name:   "infobox",
		Text:   "标题",
		Icon:   "ion-ios-cart-outline",
		Number: "90",
		Color:  "red",
	}
}

func (*Theme) AreaChart() types.AreaChartAttribute {
	return &components.AreaChartAttribute{
		Name: "area-chart",
	}
}

func (*Theme) ProgressGroup() types.ProgressGroupAttribute {
	return &components.ProgressGroupAttribute{
		Name: "progress-group",
	}
}

func (*Theme) Description() types.DescriptionAttribute {
	return &components.DescriptionAttribute{
		Name: "description",
	}
}

func (*Theme) PieChart() types.PieChartAttribute {
	return &components.PieChartAttribute{
		Name: "pie-chart",
	}
}

func (*Theme) LineChart() types.LineChartAttribute {
	return &components.LineChartAttribute{
		Name: "line-chart",
	}
}

func (*Theme) BarChart() types.BarChartAttribute {
	return &components.BarChartAttribute{
		Name: "bar-chart",
	}
}

func (*Theme) ChartLegend() types.ChartLegendAttribute {
	return &components.ChartLegendAttribute{
		Name: "chart-legend",
	}
}

func (*Theme) ProductList() types.ProductListAttribute {
	return &components.ProductListAttribute{
		Name: "productlist",
	}
}

func (*Theme) Tabs() types.TabsAttribute {
	return &components.TabsAttribute{
		Name: "tabs",
	}
}

func (*Theme) Alert() types.AlertAttribute {
	return &components.AlertAttribute{
		Name: "alert",
	}
}

func (*Theme) Label() types.LabelAttribute {
	return &components.LabelAttribute{
		Name:    "label",
		Color:   "success",
		Content: "",
	}
}

func (*Theme) Paginator() types.PaginatorAttribute {
	return &components.PaginatorAttribute{
		Name: "paginator",
	}
}

func (*Theme) Row() types.RowAttribute {
	return &components.RowAttribute{
		Name:    "row",
		Content: "",
	}
}

func (*Theme) Table() types.TableAttribute {
	return &components.TableAttribute{
		Name:     "table",
		Thead:    []map[string]string{},
		InfoList: []map[string]template.HTML{},
		Type:     "normal",
	}
}

func (theme *Theme) DataTable() types.DataTableAttribute {
	return &components.DataTableAttribute{
		TableAttribute: *(theme.Table().SetType("data-table").(*components.TableAttribute)),
		EditUrl:        "",
		NewUrl:         "",
	}
}

func (*Theme) Tree() types.TreeAttribute {
	return &components.TreeAttribute{
		Name: "tree",
		Tree: []menu.MenuItem{},
	}
}
