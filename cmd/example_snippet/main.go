package main

import (
	"fmt"
	"text/template"

	"github.com/arturoeanton/gocommons/bloom"
	"github.com/arturoeanton/gocommons/hash"
	"github.com/arturoeanton/gocommons/snippet"
)

func main() {
	/*defer func() {
		e := recover()
		log.Println("Recovering from panic:", e)
	}()
	//*/

	fmt.Println("Example Snippet")

	qb := snippet.NewSnippetStorage().Escape(template.HTMLEscapeString).Comment("--").LoadFile("test.sql")

	qt := qb.GetSnippet("test1")

	fmt.Println("test1:\n" + qt.Get())

	qt = qb.GetSnippet("test2")

	fmt.Println("test2:\n" +
		qt.
			Param("id", "hola").
			Param("name", "Elias\"--").
			Param("age", 41).
			Get(),
	)
	fmt.Println(hash.HashStringUint64("HelloWorld"))
	fmt.Println(hash.HashStringUint64("HelloWorld."))

	bl := bloom.NewBloom()
	bl.Add("HelloWorld.")
	bl.Add("HelloWorld")
	bl.Add("HelloWorld")
	bl.Add("HelloWorld")

	fmt.Println(bl.Contains("HelloWorld"))
	bl.Remove("HelloWorld")
	fmt.Println(bl.Contains("HelloWorld"))
	bl.Remove("HelloWorld")
	fmt.Println(bl.Contains("HelloWorld"))
	bl.Remove("HelloWorld")
	fmt.Println(bl.Contains("HelloWorld"))
	bl.Remove("HelloWorld")
	fmt.Println(bl.Contains("HelloWorld"))
	bl.Remove("HelloWorld")

	bl.Add("HelloWorld")
	fmt.Println(bl.Contains("HelloWorld"))
	bl.Remove("HelloWorld")
	fmt.Println(bl.Contains("HelloWorld."))
}
