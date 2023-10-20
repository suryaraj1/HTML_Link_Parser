package main

import (
	"fmt"
	"strings"

	"github.com/suryaraj1/html-link-parser/util"
)

var exampleHtml = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <a href="https://github.com/gophercises/link">Hi, I am a Link</a>
    <button>
        <a href="https://google.com">Hello gopher</a>
    </button>
</body>
</html>
`

func main() {
	fmt.Println("hello world")

	r := strings.NewReader(exampleHtml)
	links, err := util.Parser(r)

	if err != nil {
		panic(err)
	}

	for _, link := range links {
		fmt.Println(link.Href + " " + link.Text)
	}
}
