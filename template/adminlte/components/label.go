package components

import (
	"github.com/MexChina/Treasure/application/admin/view/types"
	"html/template"
)

type LabelAttribute struct {
	Name    string
	Color   string
	Content string
}

func (compo *LabelAttribute) SetContent(value string) types.LabelAttribute {
	compo.Content = value
	return compo
}

func (compo *LabelAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "label")
}
