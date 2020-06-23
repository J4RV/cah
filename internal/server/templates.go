package server

import (
	"html/template"
	"log"
	"net/http"
)

// Template definitions

type tmplID string

var tmplBase = []string{
	"layout/base.gohtml",
	"layout/base-style.gohtml",
	"layout/base-js.gohtml",
	"layout/css-reset.gohtml",
}

const (
	loginPageTmpl      tmplID = "Login"
	gamesPageTmpl      tmplID = "Games"
	createGamePageTmpl tmplID = "Create game"
	lobbyPageTmpl      tmplID = "Lobby"
	ingamePageTmpl     tmplID = "Ingame"
	notFoundPageTmpl   tmplID = "Not found"
)

var templateFiles = map[tmplID][]string{
	loginPageTmpl: append([]string{
		"login.gohtml",
	}, tmplBase...),

	notFoundPageTmpl: append([]string{
		"404.gohtml",
	}, tmplBase...),

	gamesPageTmpl: append([]string{
		"games.gohtml",
		"components/logged-header.gohtml",
	}, tmplBase...),

	lobbyPageTmpl: append([]string{
		"lobby.gohtml",
		"components/logged-header.gohtml",
		"components/import-vue.gohtml",
	}, tmplBase...),

	ingamePageTmpl: append([]string{
		"ingame.gohtml",
		"components/logged-header.gohtml",
		"components/import-vue.gohtml",
	}, tmplBase...),

	createGamePageTmpl: append([]string{
		"create-game.gohtml",
		"components/logged-header.gohtml",
	}, tmplBase...),
}

func initTemplates() {
	// adds the config path prefix to templateFiles
	for id := range templateFiles {
		templatesWithPath := templateFiles[id]
		for i := range templatesWithPath {
			templatesWithPath[i] = config.TemplatePath + templatesWithPath[i]
		}
	}
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
