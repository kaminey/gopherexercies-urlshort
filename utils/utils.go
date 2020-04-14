package utils

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// constants
const (
	JSON = "json"
	YAML = "yaml"
)

// Pair is used to store path and related url combination
type Pair struct {
	Path string
	URL  string
}

func getJSONData(filename string) ([]Pair, error) {
	var pairs []Pair

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &pairs)
	if err != nil {
		return nil, err
	}

	return pairs, nil
}

func getYAMLData(filename string) ([]Pair, error) {
	var pairs []Pair

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &pairs)
	if err != nil {
		return nil, err
	}

	return pairs, nil
}

func getFileData(filename, fileType string) ([]Pair, error) {
	if fileType == JSON {
		return getJSONData(filename)
	}
	return getYAMLData(filename)
}

// GetURLMap returns a map of path and related url
func GetURLMap(filename, fileType string) (map[string]string, error) {
	pairs, err := getFileData(filename, fileType)
	if err != nil {
		return nil, err
	}
	urlMap := make(map[string]string)
	for _, pair := range pairs {
		urlMap[pair.Path] = pair.URL
	}
	return urlMap, nil
}
