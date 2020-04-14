package utils

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		url, ok := pathsToUrls[path]
		if !ok {
			fallback.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, url, http.StatusFound)
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
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

func buildMap(yml []byte) (map[string]string, error) {
	var mpSlice []map[string]string
	err := yaml.Unmarshal(yml, &mpSlice)
	if err != nil {
		return nil, err
	}
	mp := make(map[string]string)
	for _, val := range mpSlice {
		mp[val["path"]] = val["url"]
	}
	return mp, nil
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	mp, err := buildMap(yml)
	if err != nil {
		return nil, err
	}
	return MapHandler(mp, fallback), nil
}
