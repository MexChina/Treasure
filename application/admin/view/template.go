package view

import (
	"github.com/MexChina/Treasure/template/adminlte"
	"github.com/MexChina/Treasure/application/admin/view/types"
	"html/template"
	"sync"
)

type Template interface {
	// Components
	Form() types.FormAttribute
	Box() types.BoxAttribute
	Col() types.ColAttribute
	Image() types.ImgAttribute
	SmallBox() types.SmallBoxAttribute
	Label() types.LabelAttribute
	Row() types.RowAttribute
	Table() types.TableAttribute
	DataTable() types.DataTableAttribute
	Tree() types.TreeAttribute
	InfoBox() types.InfoBoxAttribute
	Paginator() types.PaginatorAttribute
	AreaChart() types.AreaChartAttribute
	ProgressGroup() types.ProgressGroupAttribute
	LineChart() types.LineChartAttribute
	BarChart() types.BarChartAttribute
	ProductList() types.ProductListAttribute
	Description() types.DescriptionAttribute
	Alert() types.AlertAttribute
	PieChart() types.PieChartAttribute
	ChartLegend() types.ChartLegendAttribute
	Tabs() types.TabsAttribute

	GetTmplList() map[string]string
	GetAssetList(string) []string
	GetTemplate(bool) (*template.Template, string)
}

var templateMap = map[string]Template{
	"adminlte": adminlte.GetAdminlte(),
}

func Get(theme string) Template {
	if temp, ok := templateMap[theme]; ok {
		return temp
	}
	panic("wrong theme name")
}

var (
	templateMu sync.Mutex
)

func Add(name string, temp Template) {
	templateMu.Lock()
	defer templateMu.Unlock()
	if temp == nil {
		panic("template is nil")
	}
	if _, dup := templateMap[name]; dup {
		panic("add template twice " + name)
	}
	templateMap[name] = temp
}