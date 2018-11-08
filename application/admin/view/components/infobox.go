package components

import (
	"github.com/MexChina/Treasure/application/admin/view/types"
	"html/template"
)

type InfoBoxAttribute struct {
	Name    string
	Icon    string
	Text    string
	Number  template.HTML
	Content string
	Color   string
}

func (compo *InfoBoxAttribute) SetIcon(value string) types.InfoBoxAttribute {
	compo.Icon = value
	return compo
}

func (compo *InfoBoxAttribute) SetText(value string) types.InfoBoxAttribute {
	compo.Text = value
	return compo
}

func (compo *InfoBoxAttribute) SetNumber(value template.HTML) types.InfoBoxAttribute {
	compo.Number = value
	return compo
}

func (compo *InfoBoxAttribute) SetContent(value string) types.InfoBoxAttribute {
	compo.Content = value
	return compo
}

func (compo *InfoBoxAttribute) SetColor(value string) types.InfoBoxAttribute {
	compo.Color = value
	return compo
}

func (compo *InfoBoxAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "infobox")
}
