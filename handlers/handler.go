package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jatin-malik/chooseyourownadventure/models"
)

var chapterTmpl *template.Template

func init() {
	var err error
	chapterTmpl, err = template.ParseFiles("templates/chapter.html")
	if err != nil {
		log.Fatal("Error parsing chapter template:", err)
	}
}

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
	if data, ok := s.ChapterDataMap[path]; ok {
		chapterTmpl.Execute(w, data)
	} else {
		// redirect to intro
		http.Redirect(w, r, "/intro", http.StatusSeeOther)
	}
}
