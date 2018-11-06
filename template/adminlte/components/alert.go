package components

import (
	"github.com/MexChina/Treasure/template/types"
	"html/template"
)

type AlertAttribute struct {
	Name    string
	Theme   string
	Title   template.HTML
	Content template.HTML
}

func (compo *AlertAttribute) SetTheme(value string) types.AlertAttribute {
	compo.Theme = value
	return compo
}

func (compo *AlertAttribute) SetTitle(value template.HTML) types.AlertAttribute {
	compo.Title = value
	return compo
}

func (compo *AlertAttribute) SetContent(value template.HTML) types.AlertAttribute {
	compo.Content = value
	return compo
}

func (compo *AlertAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "alert")
}
