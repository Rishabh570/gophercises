package utils

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func GetSlug(link string) string {
	return "1"
}

func GetFullLinkFromSlug(domain string, slug string) string {
	fmt.Println("[GetFullLinkFromSlug] domain = ", domain)
	fmt.Println("[GetFullLinkFromSlug] slug = ", slug)
	// fullLink := url.URL{
	// 	Host: domain,
	// 	Path: slug,
	// }
	return domain + slug
}

func GetDomain(link string, domain string) string {
	url, err := url.Parse(link)
	if err != nil {
		log.Fatal("[GetDomain] Could not parse link: ", err)
		os.Exit(1)
	}

	if url.Hostname() == "" {
		return domain
	}
	hostname := strings.TrimPrefix(url.Hostname(), "www.")
	// return url.Scheme + "://" + url.Hostname()
	return hostname
}

func IsInternal(link string) bool {
	return strings.HasPrefix(link, "/")
	// linkDomain := GetDomain(link, domain)
	// // fmt.Println("linkDomain: ", linkDomain)
	// // fmt.Println("domain: ", domain)
	// return linkDomain == domain
}

func IsBlacklistedLink(link string) bool {
	return strings.HasPrefix(link, "mailto:") || strings.HasPrefix(link, "#")
}

func SanitizeURL(URL string) string {
	rootURL := strings.TrimRight(URL, "/")
	rootURL = strings.Replace(rootURL, "www.", "", 1)
	return rootURL
}
