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
	<style>
	  body {
		  font-family: Helvetica, Arial;
	  }
	  h1 {
		  text-align: center;
		  position: relative;
	  }
	  .page {
		  width: 80%;
		  max-width: 500px;
		  margin: auto;
		  margin-top: 40px;
		  margin-bottom: 40px;
		  padding: 80px;
		  background: #FFFCF6;
		  border: 1px solid #eee;
		  box-shadow: 0 10px 6px -6px #777;
	  }
	  ul {
		  border-top: 1px dotted #ccc;
		  padding: 10px 0 0 0;
		  -webkit-padding-start: 0;
	  }
	  li {
		  padding-top: 10px;
	  }
	  a,
	  a:visited {
		  text-decoration: none;
		  color: #6295b5;
	  }
	  a:active,
	  a:hover {
		  color: #7792a2;
	  }
	  p {
		  text-indent: 1em;
	  }
	</style>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      <ul>
      {{range .Options}}
        <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
      </ul>
	</section>
  </body>
</html>`

type HandlerOption func(h *handler)

// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
// Functional options - customization is performed with functions that
// operate on the handler itself
func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func NewHandler(story Story, options ...HandlerOption) http.Handler {
	t := template.Must(template.New("").Parse(defaultHandlerTemplate))
	h := handler{story, t}
	// Apply functional options
	for _, opt := range options {
		opt(&h)
	}
	return h
}

type handler struct {
	story Story
	t     *template.Template
}

func (h handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := strings.TrimSpace(request.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = strings.TrimPrefix(path, "/")
	if chapter, ok := h.story[path]; ok {
		err := h.t.Execute(writer, chapter)
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
