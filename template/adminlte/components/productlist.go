package components

import (
	"github.com/MexChina/Treasure/template/types"
	"html/template"
)

type ProductListAttribute struct {
	Name string
	Data []map[string]string
}

func (compo *ProductListAttribute) SetData(value []map[string]string) types.ProductListAttribute {
	compo.Data = value
	return compo
}

func (compo *ProductListAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "productlist")
}
