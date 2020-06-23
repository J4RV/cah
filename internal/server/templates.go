package server

import (
	"html/template"
	"log"
	"net/http"
)

type tmplID string

var tmplBase = []string{
	config.TemplatePath + "components/layout/base.gohtml",
	config.TemplatePath + "components/layout/base-style.gohtml",
	config.TemplatePath + "components/layout/base-js.gohtml",
	config.TemplatePath + "components/layout/css-reset.gohtml",
}

// Template definitions

const (
	loginPageTmpl      tmplID = "Login"
	gamesPageTmpl      tmplID = "Games"
	createGamePageTmpl tmplID = "Create game"
	lobbyPageTmpl      tmplID = "Lobby"
	ingamePageTmpl     tmplID = "Ingame"
	notFoundPageTmpl   tmplID = "Not found"
)

// FIXME reduce verbosity
var templateFiles = map[tmplID][]string{
	loginPageTmpl:      append(tmplBase, config.TemplatePath+"login.gohtml"),
	notFoundPageTmpl:   append(tmplBase, config.TemplatePath+"404.gohtml"),
	gamesPageTmpl:      append(tmplBase, config.TemplatePath+"games.gohtml", config.TemplatePath+"components/logged-header.gohtml"),
	lobbyPageTmpl:      append(tmplBase, config.TemplatePath+"lobby.gohtml", config.TemplatePath+"components/logged-header.gohtml", config.TemplatePath+"components/import-vue.gohtml"),
	ingamePageTmpl:     append(tmplBase, config.TemplatePath+"ingame.gohtml", config.TemplatePath+"components/logged-header.gohtml", config.TemplatePath+"components/import-vue.gohtml"),
	createGamePageTmpl: append(tmplBase, config.TemplatePath+"create-game.gohtml", config.TemplatePath+"components/logged-header.gohtml"),
}

// Functions to be called from outside this file to render the templates:

var compiledTemplates = map[tmplID]*template.Template{}

func execTemplate(id tmplID, w http.ResponseWriter, status int, data interface{}) {
	if compiledTemplates[id] == nil {
		compiledTemplates[id] = parseTemplate(id)
	}

	w.WriteHeader(status)
	err := compiledTemplates[id].Execute(w, data)
	if err != nil {
		logError.Println("trying to executing template", id, err)
	}

	if devMode {
		// In development mode, dont store the compiled templates
		compiledTemplates[id] = nil
	}
}

func simpleTmplHandler(id tmplID, status int) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		execTemplate(id, w, status, nil)
	}
}

// File internals

func parseTemplate(id tmplID) *template.Template {
	log.Println("Parsing template:", id)
	return template.Must(template.New("base.gohtml").Funcs(tmplFuncMap).ParseFiles(templateFiles[id]...))
}

var tmplFuncMap = template.FuncMap{
	"safeHTML": safeHTML,
	"safeCSS":  safeCSS,
	"devMode":  func() bool { return devMode },
}

func safeHTML(b string) template.HTML {
	return template.HTML(b)
}

func safeCSS(b string) template.CSS {
	return template.CSS(b)
}
