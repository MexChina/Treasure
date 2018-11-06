package components

import (
	"github.com/MexChina/Treasure/template/types"
	"html/template"
)

type LineChartAttribute struct {
	Name   string
	Title  string
	Data   string
	ID     string
	Height int
}

func (compo *LineChartAttribute) SetID(value string) types.LineChartAttribute {
	compo.ID = value
	return compo
}

func (compo *LineChartAttribute) SetTitle(value string) types.LineChartAttribute {
	compo.Title = value
	return compo
}

func (compo *LineChartAttribute) SetHeight(value int) types.LineChartAttribute {
	compo.Height = value
	return compo
}

func (compo *LineChartAttribute) SetData(value string) types.LineChartAttribute {
	compo.Data = value
	return compo
}

func (compo *LineChartAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "line-chart")
}
