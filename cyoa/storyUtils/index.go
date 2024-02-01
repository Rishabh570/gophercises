package storyUtils

import (
	"encoding/json"
	"fmt"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryObj struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

func ParseJSON(jsonData []byte) (map[string]StoryObj, []string) {
	var data map[string]StoryObj

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, nil
	}

	var keys []string
	for key := range data {
		keys = append(keys, key)
	}

	return data, keys
}
