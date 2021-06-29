package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{

}
var layoutsPath = "./templates/*.layout.tmpl"
var pagesPath = "./templates/*.page.tmpl"


func RenderTemplate(w http.ResponseWriter, tmpl string) {
	_, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("error getting template cache:", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates" + tmpl)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}


	pages, err := filepath.Glob(pagesPath)
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(layoutsPath)
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(layoutsPath)
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}