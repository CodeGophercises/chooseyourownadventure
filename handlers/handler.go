package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jatin-malik/chooseyourownadventure/models"
)

// Story contains n chapters.
// Each chapter will have a web page in our application.
// The name of the chapter will be the path to web page.
// Title will be html title, story in a <p> maybe and options are buttons with text as button text and arc as href to navigate to other chapter.

type StoryHandler struct {
	ChapterDataMap map[string]models.Chapter
}

func (s *StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Parse request
	path := r.URL.Path[1:]
	template, err := template.ParseFiles("templates/chapter.html")
	if err != nil {
		log.Fatal("error parsing html file", err.Error())
	}

	if data, ok := s.ChapterDataMap[path]; ok {
		template.Execute(w, data)
	} else {
		// redirect to intro
		http.Redirect(w, r, "/intro", http.StatusSeeOther)
	}
}
