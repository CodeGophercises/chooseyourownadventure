package models

type Chapter struct {
	Title   string
	Story   []string
	Options []Option
}

type Option struct {
	Text string
	Arc  string
}
