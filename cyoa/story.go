package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var defaultHandlerTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
      {{range .Options}}
      <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
    </ul>
  </body>
</html>`

func NewHandler(story Story) http.Handler {
	return handler{story}
}

type handler struct {
	story Story
}

func (h handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := strings.TrimSpace(request.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = strings.TrimPrefix(path, "/")
	if chapter, ok := h.story[path]; ok {
		tpl := template.Must(template.New("").Parse(defaultHandlerTemplate))
		err := tpl.Execute(writer, chapter)
		if err != nil {
			log.Printf("Failed to render the template: %v", err)
			http.Error(writer, "Something went wrong", http.StatusInternalServerError)
		}
		return
	} else {
		http.Error(writer, "Chapter not found", http.StatusNotFound)
	}
}

func JsonStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)
	var story Story
	if err := decoder.Decode(&story); err != nil {
		panic(err)
	}
	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
