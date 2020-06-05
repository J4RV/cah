package server

import (
	"html/template"
	"io"
	"log"
)

type tmplID int

const tmplDir = "templates/"

var tmplBase = []string{
	tmplDir + "layout/base.gohtml",
	tmplDir + "layout/base-style.gohtml",
	tmplDir + "layout/base-js.gohtml",
	tmplDir + "layout/css-reset.gohtml",
}

// Template definitions

const (
	loginPageTmpl tmplID = iota
)

var templateFiles = map[tmplID][]string{
	loginPageTmpl: append(tmplBase, tmplDir+"login.gohtml"),
}

// Functions to be called from outside this file to render the templates:

var compiledTemplates = map[tmplID]*template.Template{}

func execTemplate(id tmplID, w io.Writer, data interface{}) {
	if devMode {
		template.Must(template.ParseFiles(templateFiles[id]...)).Execute(w, data)
		return
	}

	if compiledTemplates[id] == nil {
		log.Println("Compiling template with id: ", id)
		compiledTemplates[id] = template.Must(template.ParseFiles(templateFiles[id]...))
	}
	compiledTemplates[id].Execute(w, data)
}
