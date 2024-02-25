package helpers

import (
	"net/url"
	"quietHN/hn"
	"quietHN/models"
	"strings"
)

func ParseHNItem(hnItem hn.Item) models.Item {
	ret := models.Item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

func IsStoryLink(item models.Item) bool {
	return item.Type == "story" && item.URL != ""
}
