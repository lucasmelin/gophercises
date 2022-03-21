package urlshort

import (
	"testing"
)

var testYaml = []byte(`
- path: /lucas
  url: https://lucasmelin.com
- path: /timeline
  url: https://timeline.lucasmelin.com
`)

func TestYamlParser(t *testing.T) {
	t.Run("Test yaml to map", func(t *testing.T) {
		paths, err := parseYaml(testYaml)
		if err != nil {
			t.Error("Error parsing YAML: ", err)
		}
		mapPaths := buildMap(paths)
		for _, up := range paths {
			if mapPaths[up.Path] != up.Url {
				t.Errorf("Expected %s, got %s", up.Url, mapPaths[up.Path])
			}
		}
	})
}
