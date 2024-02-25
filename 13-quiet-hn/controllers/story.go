package controllers

import (
	"html/template"
	"net/http"
	"quietHN/hn"
	"quietHN/models"
	"quietHN/services"
	"time"
)

var cache models.Cache
var client hn.Client

func GetStories(numStories int, tpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		stories, err := services.GetStories(&client, numStories, &cache)
		if err != nil {
			http.Error(w, "Failed to load top stories", http.StatusInternalServerError)
			return
		}

		data := models.TemplateData{
			Stories: stories,
			Time:    time.Since(start),
		}

		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
	})
}
