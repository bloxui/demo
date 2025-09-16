package main

import (
	"fmt"

	. "github.com/plainkit/html"
)

func main() {
	page := Html(
		Lang("en"),
		Head(
			HeadTitle(T("My Page")),
			Meta(Charset("UTF-8")),
			HeadStyle(T(".intro { color: blue; }")),
		),
		Body(
			H1(T("Hello, World!")),
			P(T("Built with Plain"), Class("intro")),
		),
	)

	fmt.Println("<!DOCTYPE html>")
	fmt.Println(Render(page))
}
