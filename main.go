package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/suryaraj1/html-link-parser/util"
)

func main() {
	fmt.Println("hello world")

	r := strings.NewReader("./testDocuments/test1.html")
	links, err := util.Parser(r)

	if err != nil {
		panic(err)
	}

	for i, link := range links {
		fmt.Println(strconv.Itoa(i) + " " + link.Href)
	}
}
