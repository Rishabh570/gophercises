package controllers

import (
	"cyoa/storyUtils"
	"net/http"
	"strings"
	"text/template"
)

func AttachLanding(parsedKeys []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Keys []string
		}{
			Keys: parsedKeys,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func AttachStoryRoutes(mux http.Handler, storyObj map[string]storyUtils.StoryObj) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimLeft(r.URL.Path, "/")

		if dest, ok := storyObj[path]; ok {
			tmpl, err := template.ParseFiles("template.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, dest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		mux.ServeHTTP(w, r)
	}
}
