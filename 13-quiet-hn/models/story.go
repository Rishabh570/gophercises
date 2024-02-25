package models

import (
	"quietHN/hn"
	"time"
)

// item is the same as the hn.Item, but adds the Host field
type Item struct {
	hn.Item
	Host string
}

type TemplateData struct {
	Stories []Item
	Time    time.Duration
}

type Result struct {
	Index int
	Item  Item
	Err   error
}
