package cyoa

import (
	"encoding/json"
	"io"
)

func JsonStory(r io.Reader) (Story, error) {
	// Decoder takes an io.Reader
	d := json.NewDecoder(r)
	var story Story
	// Decode takes a pointer, can't just use story here
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
