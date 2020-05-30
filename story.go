package cyoa

import (
	"encoding/json"
	"io/ioutil"
)

type Story map[string]StoryArc

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func ParseJSON() (story Story) {
	jsonFile, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(jsonFile, &story); err != nil {
		panic(err)
	}
	return story
}
