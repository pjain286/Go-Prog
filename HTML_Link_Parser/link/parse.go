package link

import (
	"io"
	"golang.org/x/net/html"
)

/**Struct to represent Link*/
type Link struct {
	Href string
	Text string
}

/**Parse Function parses the input stream and return []Link*/
/**Input - *io.Reader*/
/**Output - []Link */
/**NOTE : The caller must make sure that io.Reader stream is UTF-8 Encoded.*/
func Parse(file io.Reader) ([]Link,error){
	parseTree,err := html.Parse(file)
	if err != nil{
		return nil,err
	}
	result := DFS(parseTree)
	return result,nil
}

/**DFS to identify <a> nodes and process them*/
/**Input - *html.Node*/
/**Output - []Link*/
func DFS(n *html.Node) []Link{
	
	var result []Link
	var res Link
	if n.Type == html.ElementNode && n.Data == "a"{
			for _,attr := range n.Attr{
				if attr.Key == "href"{
					res.Href = attr.Val
				}
				res.Text = parseText(n)
			}
			result = append(result,res)
			return result
	}

	for c := n.FirstChild; c != nil ; c = c.NextSibling{
		result = append(result,DFS(c)...)
	}

	return result
}


/**Parses the text under a html.Node*/
/**Input - *html.Node*/
/**OUtput - string*/
func parseText(n *html.Node) string{
	var result string
	if n.Type == html.TextNode {
		result = n.Data
		return result
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling{
		result = result + " " + parseText(c)
	}

	return result
}