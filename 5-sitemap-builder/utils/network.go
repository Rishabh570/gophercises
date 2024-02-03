package utils

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

func PerformHTTPGETRequest(link string) (*http.Response, string) {
	// fmt.Println("[PerformHTTPGETRequest] domain to GET: ", link)
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal("Could not fetch page")
		os.Exit(1)
	}

	reqURL := resp.Request.URL

	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	absoluteBaseURL := baseURL.String()

	// res, e := io.ReadAll(resp.Body)
	// if e != nil {
	// 	log.Fatal("Could not read response body")
	// 	os.Exit(1)
	// }

	// defer resp.Body.Close()
	// return string(res)
	return resp, absoluteBaseURL
}
