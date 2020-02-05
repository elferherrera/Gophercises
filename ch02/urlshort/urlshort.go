package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathMap map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path

		if dest, ok := pathMap[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		}

		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYAML(yamlBytes)
	if err != nil {
		return nil, err
	}

	pathsToUrls := buildPaths(pathUrls)
	return MapHandler(pathsToUrls, fallback), nil
}

func buildPaths(pathUrls []pathURL) map[string]string {
	pathsToURL := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToURL[pu.Path] = pu.URL
	}

	return pathsToURL
}

func parseYAML(yamlBytes []byte) ([]pathURL, error) {
	var pathUrls []pathURL
	err := yaml.Unmarshal(yamlBytes, &pathUrls)

	if err != nil {
		return nil, nil
	}

	return pathUrls, nil
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
