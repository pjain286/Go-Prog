package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"flag"
	"./urlshort"
)

func main() {
	fileName := flag.String("json","sample.json","json file to read data")
	flag.Parse()

	jsonFile,err := os.Open(*fileName)
	if err != nil{
		panic(err)
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Read from the json file and create map
	jsn,_ := ioutil.ReadAll(jsonFile)
	jsonHandler, err := urlshort.JSONHandler([]byte(jsn), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}