package models

import (
	"fmt"
	"strings"
)

type Chapter struct {
	Title   string
	Story   []string
	Options []Option
}

type Option struct {
	Text string
	Arc  string
}

func (c *Chapter) RenderOnScreen() {
	fmt.Println(c.Title)
	fmt.Printf("\n")
	fmt.Println(strings.Join(c.Story, " "))
	fmt.Println()
	for i, option := range c.Options {
		fmt.Printf("Press %d to %s\n", i+1, option.Text)
	}
}
