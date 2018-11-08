package components

import (
	"github.com/MexChina/Treasure/application/admin/view/types"
	"html/template"
)

type BarChartAttribute struct {
	Name  string
	Title string
	Data  string
	ID    string
	Width int
}

func (compo *BarChartAttribute) SetID(value string) types.BarChartAttribute {
	compo.ID = value
	return compo
}

func (compo *BarChartAttribute) SetTitle(value string) types.BarChartAttribute {
	compo.Title = value
	return compo
}

func (compo *BarChartAttribute) SetWidth(value int) types.BarChartAttribute {
	compo.Width = value
	return compo
}

func (compo *BarChartAttribute) SetData(value string) types.BarChartAttribute {
	compo.Data = value
	return compo
}

func (compo *BarChartAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "bar-chart")
}
