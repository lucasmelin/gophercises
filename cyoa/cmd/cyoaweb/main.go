package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lucasmelin/gophercises/cyoa"
)

func main() {
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
	fmt.Printf("%+v\n", story)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
