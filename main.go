package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jatin-malik/chooseyourownadventure/handlers"
	"github.com/jatin-malik/chooseyourownadventure/models"
)

func main() {

	// We have the story in a file called `gopher.json`
	file, err := os.Open("gopher.json")
	if err != nil {
		log.Fatalf("error while opening %s", file.Name())
	}
	story_bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("error while reading from story file")
	}
	// Decode the json story in a Go type
	story := make(map[string]models.Chapter)
	if err := json.Unmarshal(story_bytes, &story); err != nil {
		log.Fatal("error while unmarshalling story")
	}

	// We have the story in a map now
	//fmt.Println(story)

	// Solve for the presentation how does each chapter render on the web page
	// Use html/template for data driven html templates

	handler := handlers.StoryHandler{story}
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", &handler))
}
