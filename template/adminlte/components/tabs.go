package components

import (
	"github.com/MexChina/Treasure/template/types"
	"html/template"
)

type TabsAttribute struct {
	Name string
	Data []map[string]template.HTML
}

func (compo *TabsAttribute) SetData(value []map[string]template.HTML) types.TabsAttribute {
	compo.Data = value
	return compo
}

func (compo *TabsAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "tabs")
}
