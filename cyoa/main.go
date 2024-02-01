package main

// TIL: If folder name matches the package name, all entities (fn, struct, interfaces, etc.) in all files inside that folder will be accessible with package name prefix
// For eg. if folder name = storyUtils, and all files inside it are using _package storyUtils_, then
// we can use all entities across all files inside storyUtils folders by doing: storyUtils.<ENTITY_NAME>
// However, if folder name is not equal to package name, then
//   import structure for a file inside that folder becomes: <FILENAME> "<MODULE_NAME/PACKAGE_NAME>"
//   eg. folder storyUtils/, file inside is story.go; main.go will have _story cyoa/storyUtils_ to access entities inside story.go
import (
	"cyoa/controllers"
	"cyoa/storyUtils"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// 0. read file
	filePath := "gopher.json"
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("could not read json file")
		return
	}

	// 1. parse json and store in a struct
	parsedJSON, parsedKeys := storyUtils.ParseJSON(jsonData)
	if parsedJSON == nil {
		log.Fatal("Could not parse json")
		return
	}
	fmt.Println("keys: ", parsedKeys)

	// 2. create a server where landing page should show list of all keys to start the story from
	// 3. each slug will be a link and clicking on it takes you to that story
	// 4. each story page will have title, content, and options
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.AttachLanding(parsedKeys))
	handler := controllers.AttachStoryRoutes(mux, parsedJSON)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}
