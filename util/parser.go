package util

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Type definition for Link
type Link struct {
	Href string
	Text string
}

func Parser(r io.Reader) ([]Link, error) {
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			// ...
			return nil, nil
		}

		tagName, _ := z.TagName()
		fmt.Println(string(tagName))
	}

	return []Link{
		{Href: "example 1", Text: "i am example 1"},
	}, nil
}
