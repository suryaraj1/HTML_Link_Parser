package util

import (
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
	links := []Link{}

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			// End of Doc
			return links, nil
		}

		tagName, _ := z.TagName()
		if string(tagName) == "a" {
			_, val, _ := z.TagAttr()

			tt = z.Next()

			links = append(links, Link{
				Href: string(val),
				Text: string(z.Text()),
			})
		}
	}
}
