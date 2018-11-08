// Copyright 2018 ChenHonggui.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package template

import (
	"github.com/MexChina/Treasure/template/adminlte"
	"github.com/MexChina/Treasure/template/types"
	"html/template"
	"sync"
)

// Template is the interface which contains methods of ui components.
// It will be used in the plugins for custom the ui.
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

	// Builder methods
	GetTmplList() map[string]string
	GetAssetList(string) []string
	GetTemplate(bool) (*template.Template, string)
}

// The templateMap contains templates registered.
var templateMap = map[string]Template{
	"adminlte": adminlte.GetAdminlte(),
}

// Get the template interface by theme name. If the
// name is not found, it panics.
func Get(theme string) Template {
	if temp, ok := templateMap[theme]; ok {
		return temp
	}
	panic("wrong theme name")
}

var (
	templateMu sync.Mutex
)

// Add makes a template available by the provided theme name.
// If Add is called twice with the same name or if template is nil,
// it panics.
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