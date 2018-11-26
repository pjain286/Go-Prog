package urlshort

import (
	"fmt"
	"net/http"
	"encoding/json"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request){
		path := r.URL.Path
		if dest,ok := pathsToUrls[path]; ok{
			http.Redirect(w,r,dest,http.StatusFound)
			return
		}
		fallback.ServeHTTP(w,r)
	}
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//     { "Path" : "/some-path",
//       "Url": "https://www.some-url.com/demo"
//	 	}
//
// The only errors that can be returned all related to having
// invalid JSON data.


type JSONMap struct{
	Path string
	Url string 
}

func parseJSON(jsn []byte) (objects []JSONMap,err error){
	err = json.Unmarshal(jsn,&objects)
	if err!=nil{
		fmt.Printf("Cannot unmarshal data: %v", err)
	}
	return objects,err
}

func buildMap(objects []JSONMap) map[string]string{
	pathsToUrls := make(map[string]string)
	for _,pu := range objects{
		pathsToUrls[pu.Path] = pu.Url
	}
	return pathsToUrls
}

func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	
	parsedJson,err := parseJSON(jsn)
	if err!=nil{
		return nil,err
	}
	pathMAP := buildMap(parsedJson)
	return MapHandler(pathMAP,fallback),nil
}