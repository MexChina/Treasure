package components

import (
	"github.com/MexChina/Treasure/application/admin/view/types"
	"html/template"
)

type ColAttribute struct {
	Name    string
	Content template.HTML
	Size    string
}

func (compo *ColAttribute) SetContent(value template.HTML) types.ColAttribute {
	compo.Content = value
	return compo
}

func (compo *ColAttribute) SetSize(value map[string]string) types.ColAttribute {
	compo.Size = ""
	for key, size := range value {
		compo.Size += "col-" + key + "-" + size + " "
	}
	return compo
}

func (compo *ColAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "col")
}
