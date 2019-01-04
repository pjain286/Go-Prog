package main

import (
	"fmt"
	"os"
	"flag"
	"./link"
)

func main() {
	fileName := flag.String("html","sample.html","html file to parse and print hyperlink data")
	flag.Parse()

	htmlFile,err := os.Open(*fileName)
	if err!=nil {
		panic(err)
	}

	result,_ := link.Parse(htmlFile)
	for _,res := range result{
		fmt.Printf("Href - %s\n",res.Href)
		fmt.Printf("Text - %s\n\n",res.Text)
	}
}