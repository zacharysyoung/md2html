package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/renderer/html"
)

func main() {
	source, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	p := parser.New()
	r := html.New()

	doc := p.Parse(source)
	if err := r.Render(&buf, source, doc); err != nil {
		panic(err)
	}

	fmt.Println(`<html>
    <head>
        <style type="text/css">
            :root {
                color-scheme: light dark;
            }
        </style>
    </head>
    <body>`)

	fmt.Println(buf.String())

	fmt.Println(`</body>
	</html>
	`)
}
