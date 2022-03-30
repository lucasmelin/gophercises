package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lucasmelin/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 8000, "the port to start the web application on")
	filenamePtr := flag.String("file", "gopher.json", "a JSON file with a Choose Your Own Adventure story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *filenamePtr)

	file, err := os.Open(*filenamePtr)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the story file: %s", *filenamePtr))
	}
	story, err := cyoa.JsonStory(file)
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the story: %s", *filenamePtr))
	}
	handler := cyoa.NewHandler(story, nil)
	fmt.Printf("Starting the server on port %d\n", *port)
	// This will log if the server returns a value, which shouldn't normally happen
	// since the server should run indefinitely
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
