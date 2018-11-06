package components

import (
	"github.com/MexChina/Treasure/modules/menu"
	"github.com/MexChina/Treasure/template/types"
	"html/template"
)

type TreeAttribute struct {
	Name      string
	Tree      []menu.MenuItem
	EditUrl   string
	DeleteUrl string
	OrderUrl  string
}

func (compo *TreeAttribute) SetTree(value []menu.MenuItem) types.TreeAttribute {
	compo.Tree = value
	return compo
}

func (compo *TreeAttribute) SetEditUrl(value string) types.TreeAttribute {
	compo.EditUrl = value
	return compo
}

func (compo *TreeAttribute) SetDeleteUrl(value string) types.TreeAttribute {
	compo.DeleteUrl = value
	return compo
}

func (compo *TreeAttribute) SetOrderUrl(value string) types.TreeAttribute {
	compo.OrderUrl = value
	return compo
}

func (compo *TreeAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "tree")
}

func (compo *TreeAttribute) GetTreeHeader() template.HTML {
	return ComposeHtml(*compo, "tree-header")
}
