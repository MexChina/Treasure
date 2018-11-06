package components

import (
	"github.com/MexChina/Treasure/template/types"
	"html/template"
)

type RowAttribute struct {
	Name    string
	Content template.HTML
}

func (compo *RowAttribute) SetContent(value template.HTML) types.RowAttribute {
	compo.Content = value
	return compo
}

func (compo *RowAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "row")
}
