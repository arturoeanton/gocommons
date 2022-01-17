package main

import (
	"fmt"
	"text/template"

	"github.com/arturoeanton/gocommons/snippet"
)

func main() {

	fmt.Println("Example QueryBatery")

	qb := snippet.LoadFile("test.sql")

	qt := qb.GetSnippet("test1")

	fmt.Println("test1:\n" + qt.Get())

	qt = qb.GetSnippet("test2")

	fmt.Println("test2:\n" +
		qt.Escape(template.HTMLEscapeString).
			Param("id", "hola").
			Param("name", "Elias\"--").
			Param("age", 41).
			Get(),
	)

}
