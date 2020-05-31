package parse

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var defualtTemplate = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Choose Your Own Adventure</title>
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{range .Story}}
            <p>{{.}}</p>
        {{end}}
        <ol>
            {{range .Options}}
            <li><a href="/{{.Arc}}">{{.Text}}</a></li>
            {{end}}
        </ol>
    </body>
</html>`

func HandlePage(s Story) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "" || path == "/" {
			path = "/intro"
		}
		path = path[1:]
		t := template.Must(template.ParseFiles("tmp.html"))
		err := t.Execute(w, s[path])
		if err != nil {
			log.Fatalf("template execution: %s", err)
		}
	}
}

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
