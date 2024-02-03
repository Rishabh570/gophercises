package utils

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type URL struct {
	Loc string `xml:"loc"`
}

type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

func GenerateSitemap(links []string) string {
	var sitemap Sitemap
	sitemap.XMLNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

	for _, url := range links {
		sitemap.URLs = append(sitemap.URLs, URL{Loc: url})
	}

	xmlData, err := xml.MarshalIndent(sitemap, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		return "" // TODO: learn proper err handling
	}

	return string(xmlData)
}

func BuildSitemapLinks(url string) []string {
	// sanitize initial url entered by the user
	rootURL := SanitizeURL(url)

	visited := make(map[string]bool)
	queue := []string{rootURL}
	var result []string

	for len(queue) > 0 {
		curr := queue[0]

		// keep skipping urls that are already visited
		for visited[curr] {
			queue = queue[1:]
			curr = queue[0]
		}

		// mark url as visited
		visited[curr] = true
		// remove it from the queue
		queue = queue[1:]

		// Scrape url and get links on the page
		links := GetLinks(curr)

		for link := range links {
			if !visited[link] {
				queue = append(queue, link)
			}
		}

		result = append(result, curr)
	}

	return result
}

func GetLinks(url string) map[string]bool {
	// fetch webpage
	resp, absoluteBaseURL := PerformHTTPGETRequest(url)
	defer resp.Body.Close()

	// Get a io reader
	reader := io.Reader(resp.Body)

	// Parse the html
	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Println("Error while parsing HTML:", err)
		os.Exit(1)
	}

	// collect all the internal and unique links on the page
	links := make(map[string]bool)
	scrapeLinks(doc, absoluteBaseURL, links)

	mutatedLinks := make(map[string]bool)
	for link := range links {
		mutatedLink := SanitizeURL(link)
		mutatedLinks[mutatedLink] = true
	}

	return mutatedLinks
}

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

// Private functions
func scrapeLinks(n *html.Node, absoluteBaseURL string, links map[string]bool) {
	// if current node is an anchor node, extract link
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				hrefLink := strings.TrimSpace(a.Val)
				// Only append to links slice if the link is internal and unique
				if IsInternal(hrefLink) && !IsBlacklistedLink(hrefLink) && !links[hrefLink] {
					links[absoluteBaseURL+hrefLink] = true
					break
				}
			}
		}
	}

	// Go over the child nodes of current node in a DFS fashion
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		scrapeLinks(c, absoluteBaseURL, links)
	}
}
