package components

import (
	"github.com/MexChina/Treasure/application/admin/view/types"
	"html/template"
)

type PieChartAttribute struct {
	Name   string
	ID     string
	Height int
	Data   string
	Title  string
}

func (compo *PieChartAttribute) SetID(value string) types.PieChartAttribute {
	compo.ID = value
	return compo
}

func (compo *PieChartAttribute) SetTitle(value string) types.PieChartAttribute {
	compo.Title = value
	return compo
}

func (compo *PieChartAttribute) SetData(value string) types.PieChartAttribute {
	compo.Data = value
	return compo
}

func (compo *PieChartAttribute) SetHeight(value int) types.PieChartAttribute {
	compo.Height = value
	return compo
}

func (compo *PieChartAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "pie-chart")
}
