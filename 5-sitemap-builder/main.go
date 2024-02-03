package main

import (
	"flag"
	"fmt"
	"sitemap-builder/utils"
)

func main() {
	// take user input - website domain, depth (optional)
	// fetch the page HTML with a GET request
	// scrape the page and collect only internal unique links
	// For each scraped link:
	//  -	ignore if already seen
	//  - otherwise, add it to BFS queue

	domain := flag.String("domain", "https://example.com", "domain to build sitemap for")
	flag.Parse()

	links := utils.BuildSitemapLinks(*domain)

	// fmt.Printf("\n🏁 Results (%d):\n", len(links))
	// for _, link := range links {
	// 	fmt.Println(link)
	// }

	sitemapStringified := utils.GenerateSitemap(links)

	fmt.Println(sitemapStringified)
}
