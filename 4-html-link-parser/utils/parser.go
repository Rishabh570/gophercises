package utils

import (
	"flag"
	"fmt"
	"html-link-parser/models"
	"os"

	"golang.org/x/net/html"
)

func ReadHTMLFile() *os.File {
	path := flag.String("path", "sample-htmls/ex4.html", "path of the sample HTML file")
	flag.Parse()

	// Open the HTML file
	file, err := os.Open(*path)
	if err != nil {
		fmt.Println("Error while opening file:", err)
		os.Exit(1)
	}

	return file
}

func ScrapeHTML(n *html.Node) {
	// if current node is an anchor node, extract link and text
	if n.Type == html.ElementNode && n.Data == "a" {
		var myAnchor models.Anchor

		for _, a := range n.Attr {
			if a.Key == "href" {
				myAnchor.Href = a.Val
			}
		}

		ExtractAnchorContent(n.FirstChild, &myAnchor)
		models.Anchors = append(models.Anchors, myAnchor)
	}

	// Go over the child nodes of current node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ScrapeHTML(c)
	}
}

func ExtractAnchorContent(n *html.Node, ma *models.Anchor) {
	for c := n; c != nil; c = c.FirstChild {
		if n.Type == html.TextNode {
			ma.Text = append(ma.Text, n.Data)
			return
		}
	}
}
