package components

import (
	"github.com/MexChina/Treasure/application/admin/view/types"
	"html/template"
)

type BoxAttribute struct {
	Name       string
	Header     template.HTML
	Body       template.HTML
	Footer     template.HTML
	Title      template.HTML
	Theme      string
	HeadBorder string
}

func (compo *BoxAttribute) SetTheme(value string) types.BoxAttribute {
	compo.Theme = value
	return compo
}

func (compo *BoxAttribute) SetHeader(value template.HTML) types.BoxAttribute {
	compo.Header = value
	return compo
}

func (compo *BoxAttribute) SetBody(value template.HTML) types.BoxAttribute {
	compo.Body = value
	return compo
}

func (compo *BoxAttribute) SetFooter(value template.HTML) types.BoxAttribute {
	compo.Footer = value
	return compo
}

func (compo *BoxAttribute) SetTitle(value template.HTML) types.BoxAttribute {
	compo.Title = value
	return compo
}

func (compo *BoxAttribute) WithHeadBorder(has bool) types.BoxAttribute {
	if has {
		compo.HeadBorder = "with-border"
	} else {
		compo.HeadBorder = ""
	}
	return compo
}

func (compo *BoxAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "box")
}
