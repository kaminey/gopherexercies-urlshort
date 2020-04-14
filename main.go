package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"urlshort/utils"
)

func main() {

	yamlFile := flag.String("yaml", "data/url.yaml", "location of yaml file.")
	jsonFile := flag.String("csv", "data/url.json", "location of csv file.")

	flag.Parse()

	mux := defaultMux()

	pathsToUrls, err := utils.GetUrlMap(*jsonFile, utils.JSON)
	if err != nil {
		log.Fatalln(err)
	}
	mapHandler := utils.MapHandler(pathsToUrls, mux)

	pathsToUrls, err = utils.GetUrlMap(*yamlFile, utils.YAML)
	if err != nil {
		log.Fatalln(err)
	}
	mapHandler = utils.MapHandler(pathsToUrls, mapHandler)

	fmt.Println("started server :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
