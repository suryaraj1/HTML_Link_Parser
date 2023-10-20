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
	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	dfs(doc, "")

	return nil, nil
}

func dfs(n *html.Node, padding string) {
	fmt.Println(padding, n.Data)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
