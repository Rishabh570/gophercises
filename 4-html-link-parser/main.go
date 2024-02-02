package main

import (
	"fmt"
	"html-link-parser/models"
	"html-link-parser/utils"
	"io"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// Reads the HTML file from disk
	file := utils.ReadHTMLFile()

	// execute when main() returns
	defer file.Close()

	// Get a reader
	reader := io.Reader(file)

	// Parse the html
	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Println("Error while parsing HTML:", err)
		os.Exit(1)
	}

	// Scrape the HTML for anchor tags
	utils.ScrapeHTML(doc)

	// print output
	for i, anchor := range models.Anchors {
		fmt.Printf("[Anchor %d]: Link = %s, Text = %s\n", i, anchor.Href, anchor.Text)
	}
}
