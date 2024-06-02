package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

type APICredentials struct {
	Key    string `json:"api_key"`
	Secret string `json:"api_secret"`
}

func main() {
	// 1. load secrets
	file, err := os.Open("secrets.json")
	if err != nil {
		fmt.Println("Error while opening file:", err)
		os.Exit(1)
	}

	defer file.Close()

	// Decode the JSON data into a struct
	var creds APICredentials
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&creds); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Print the decoded data
	fmt.Println("Name:", creds.Key)
	fmt.Println("Secret:", creds.Secret)

	// 2. encode as per RFC 1738
	encodedAPIKey := url.QueryEscape(creds.Key)
	encodedAPISecret := url.QueryEscape(creds.Secret)
	fmt.Println(encodedAPIKey)
	fmt.Println(encodedAPISecret)

	mergedCreds := encodedAPIKey + ":" + encodedAPISecret
	fmt.Println("mergedCreds:", mergedCreds)

	base64EncodedMergedCreds := base64.StdEncoding.EncodeToString([]byte(mergedCreds))
	fmt.Println(base64EncodedMergedCreds)
}
