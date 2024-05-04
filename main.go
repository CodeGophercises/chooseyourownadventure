package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jatin-malik/chooseyourownadventure/handlers"
	"github.com/jatin-malik/chooseyourownadventure/models"
)

var mode = flag.String("mode", "cmd", "mode of story telling (cmd or web)")

func cmdStoryMode(story map[string]models.Chapter) {
	fmt.Println("Welcome to Choose Your Own Adventure. This is gonna be a fun ride.\n\n")
	curChapter := "intro"
	for {
		chapterData := story[curChapter]
		fmt.Println()
		chapterData.RenderOnScreen()
		if len(chapterData.Options) == 0 {
			break
		} else {
			// take user input
			var userOption int
			fmt.Printf("\n\n>>")
			fmt.Scanf("%d\n", &userOption)
			if userOption > len(chapterData.Options) {
				fmt.Println("Not sure if there is a path down that option.")
				break
			}
			curChapter = chapterData.Options[userOption-1].Arc
		}

	}
	fmt.Println("\n\n\nSee ya soon. Thanks for playing!")
}

func main() {
	flag.Parse()
	file, err := os.Open("gopher.json")
	if err != nil {
		log.Fatalf("error while opening %s", file.Name())
	}
	story_bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("error while reading from story file")
	}
	story := make(map[string]models.Chapter)
	if err := json.Unmarshal(story_bytes, &story); err != nil {
		log.Fatal("error while unmarshalling story")
	}
	if *mode == "web" {
		handler := handlers.StoryHandler{story}
		fmt.Println("Starting server on port 8080")
		log.Fatal(http.ListenAndServe(":8080", &handler))
	} else {
		// Command line mode
		cmdStoryMode(story)
	}
}
