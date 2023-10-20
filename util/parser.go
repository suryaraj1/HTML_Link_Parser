package util

import "io"

// Type definition for Link
type Link struct {
	Href string
	Text string
}

func Parser(r io.Reader) ([]Link, error) {
	links := []Link{
		{Href: "/example1", Text: "Example Link 1"},
		{Href: "/example2", Text: "Example Link 2"},
		{Href: "/example3", Text: "Example Link 3"},
	}

	return links, nil
}
