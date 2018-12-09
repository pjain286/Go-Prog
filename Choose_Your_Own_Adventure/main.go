package main

import(
	"fmt"
	"encoding/json"
	"html/template"
	"io"
	"os"
	"net/http"
	"flag"
)

/**Structs for decoding the json format */
type Chapter struct{
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []struct{
			Text string `json:"text"`
			Arc string `json:"arc"`
	}`json:"options"`
}
/**Global variable which stores all the Chapters */
type Story map[string]Chapter
var stories Story

/**Json Handler that converts the json into a map of Chapters */
func JsonHandler(x io.Reader)(Story,error){
	d := json.NewDecoder(x)
	var story Story
	if err := d.Decode(&story); err != nil{
		return nil,err
	}
	return story,nil
}

/**HTTP handler function for web application */
func indexHandler(w http.ResponseWriter,r *http.Request){
	t,_ := template.ParseFiles("sample.html")
	x := r.URL.Path
	if len(x)==1{
		x="/intro"
	}
	x = x[1:len(x)]
	err := t.Execute(w,stories[x])
	if err!=nil{
		fmt.Println(err)
	}
}
/**Display story on CLI interface */
func commandHandler(){
	var inp int
	chapter := "intro"
	
	for{
		fmt.Printf("\n\n\n");
		fmt.Println(stories[chapter].Title)
		for i := 0; i < len(stories[chapter].Story); i++{
			fmt.Println(stories[chapter].Story[i])	
		}
		
		if chapter=="home"{
			break
		}

		fmt.Println("Options : ")
		for i := 0; i < len(stories[chapter].Options); i++{
			fmt.Println(i+1," : ",stories[chapter].Options[i].Text)
		}
		fmt.Printf("Choose One : ")
		fmt.Scanf("%d",&inp)
		inp = inp -1
		chapter = stories[chapter].Options[inp].Arc;
	}
}

func main(){
	
	//Command line arguments
	fileName := flag.String("json","cyoa.json","Json file to read stories from.")
	appType := flag.String("type","web","Type of applications (options: \"cli\" , \"web\" )")
	flag.Parse()

	// Open the JSON file and parse it
	jsonFile,err := os.Open(*fileName)
	if err != nil{
		panic(err)
	}
	
	stories,err = JsonHandler(jsonFile)
	if err != nil{
		panic(err)
	}

	// Run either as web application or CLI 
	if *appType == "web"{	
		http.HandleFunc("/",indexHandler)
		fmt.Println("Server Running on localhost:8080")
		http.ListenAndServe(":8080",nil)
	}else if *appType == "cli"{
		commandHandler()
	}else{
		fmt.Println("Wrong command line arguments !! Use \" ./main --help\" for options !!")
	}

}

