package components

import (
	"bytes"
	"fmt"
	"github.com/MexChina/Treasure/modules/language"
	"github.com/MexChina/Treasure/template/adminlte/tmpl"
	"html/template"
	"strings"
)

func ComposeHtml(compo interface{}, templateName ...string) template.HTML {
	var text = ""
	for _, v := range templateName {
		text += tmpl.List["components/"+v]
	}

	tmpla, err := template.New("comp").Funcs(template.FuncMap{
		"lang":     language.Get,
		"langHtml": language.GetFromHtml,
	}).Parse(text)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	buffer := new(bytes.Buffer)

	defineName := strings.Replace(templateName[0], "table/", "", -1)
	defineName = strings.Replace(defineName, "form/", "", -1)

	err = tmpla.ExecuteTemplate(buffer, defineName, compo)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
