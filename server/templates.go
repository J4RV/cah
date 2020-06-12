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
	gamesPageTmpl
	lobbyPageTmpl
)

var templateFiles = map[tmplID][]string{
	loginPageTmpl: append(tmplBase, tmplDir+"login.gohtml"),
	gamesPageTmpl: append(tmplBase, tmplDir+"games.gohtml"),
	lobbyPageTmpl: append(tmplBase, tmplDir+"lobby.gohtml"),
}

// Functions to be called from outside this file to render the templates:

var compiledTemplates = map[tmplID]*template.Template{}

func execTemplate(id tmplID, w io.Writer, data interface{}) {
	if compiledTemplates[id] == nil {
		compiledTemplates[id] = parseTemplate(id)
	}

	err := compiledTemplates[id].Execute(w, data)
	if err != nil {
		log.Println("Error while executing template", id, err)
	}

	if devMode {
		// In development mode, dont store the compiled templates
		compiledTemplates[id] = nil
	}
}

func parseTemplate(id tmplID) *template.Template {
	log.Println("Parsing template with id: ", id)
	return template.Must(template.New("base.gohtml").Funcs(tmplFuncMap).ParseFiles(templateFiles[id]...))
}

var tmplFuncMap = template.FuncMap{
	"safeHTML": safeHTML,
	"safeCSS":  safeCSS,
}

func safeHTML(b string) template.HTML {
	return template.HTML(b)
}

func safeCSS(b string) template.CSS {
	return template.CSS(b)
}
