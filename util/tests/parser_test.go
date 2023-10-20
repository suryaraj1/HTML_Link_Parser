package tests

import (
	"strings"
	"testing"

	"github.com/suryaraj1/html-link-parser/util"
)

func TestParser(t *testing.T) {
	htmlContent := `
		<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<a href="https://example.com">Example Link 1</a>
		<a href="https://example.org">Example Link 2</a>
	</body>
	</html>
	`

	r := strings.NewReader(htmlContent)
	links, err := util.Parser(r)

	if err != nil {
		t.Errorf("Error parsing HTML: %v", err)
	}

	expectedLinks := []util.Link{
		{Href: "https://example.com", Text: "Example Link 1"},
		{Href: "https://example.org", Text: "Example Link 2"},
	}

	if len(links) != len(expectedLinks) {
		t.Errorf("Expected %d links, but got %d", len(expectedLinks), len(links))
	}

	for i, link := range links {
		if link.Href != expectedLinks[i].Href || link.Text != expectedLinks[i].Text {
			t.Errorf("Mismatch in link %d. Expected: %+v, Actual: %+v", i, expectedLinks[i], link)
		}
	}
}
