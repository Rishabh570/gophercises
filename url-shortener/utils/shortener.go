package utils

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// serve if we can
		if dest, ok := (pathsToUrls[r.URL.Path]); ok {
			http.Redirect(w, r, dest, http.StatusFound)
		}
		// otherwise, fallback to defaultMux
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYAML, err := parseYAMLFile(yml)
	if err != nil {
		return nil, err
	}

	mappedYAML := buildMap(parsedYAML)
	return MapHandler(mappedYAML, fallback), nil
}

type Config struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYAMLFile(yml []byte) ([]Config, error) {
	var config []Config
	err := yaml.Unmarshal(yml, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	fmt.Println()
	for _, item := range config {
		fmt.Printf("Path: %s, URL: %s\n", item.Path, item.URL)
	}
	fmt.Println("Config : ")
	fmt.Println(config)
	return config, nil
}

func buildMap(config []Config) map[string]string {
	mapping := make(map[string]string)
	for _, item := range config {
		mapping[item.Path] = item.URL
	}
	return mapping
}
