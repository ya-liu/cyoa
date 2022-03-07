package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/ya-liu/cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "The JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	// Decoder takes an io.Reader
	d := json.NewDecoder(f)
	var story cyoa.Story
	// Decode takes a pointer, can't just use story here
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
