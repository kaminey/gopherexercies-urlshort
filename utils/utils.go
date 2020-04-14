package utils

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	JSON = "json"
	YAML = "yaml"
)

type Pair struct {
	Path string
	Url  string
}

func getJsonData(filename string) ([]Pair, error) {
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

func getYamlData(filename string) ([]Pair, error) {
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
		return getJsonData(filename)
	}
	return getYamlData(filename)
}

func GetUrlMap(filename, fileType string) (map[string]string, error) {
	pairs, err := getFileData(filename, fileType)
	if err != nil {
		return nil, err
	}
	urlMap := make(map[string]string)
	for _, pair := range pairs {
		urlMap[pair.Path] = pair.Url
	}
	return urlMap, nil
}
