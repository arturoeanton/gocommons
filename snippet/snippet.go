package snippet

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type Var struct {
	Key   string
	Value string
	Type  string
}

type Snippet struct {
	Name     string
	Type     string
	Text     string
	Vars     map[string]*Var
	FxEscape func(string) string
}
type Snippets struct {
	snippet map[string]*Snippet
	file    string
}

func LoadFileComment(name string, comment string, fxEscape func(string) string) *Snippets {
	snippet := make(map[string]*Snippet)

	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	key := ""
	varType := "string"
	for scanner.Scan() {
		line := scanner.Text()
		lineTmp := strings.ReplaceAll(line, " ", "")
		if len(lineTmp) > 0 && strings.HasPrefix(lineTmp, comment+"name:") {
			key = strings.TrimPrefix(lineTmp, comment+"name:")
			snippet[key] = &Snippet{Name: key, Type: "", Text: "", Vars: make(map[string]*Var), FxEscape: fxEscape}
			continue
		}
		if key == "" {
			continue
		}
		if len(lineTmp) > 0 && strings.HasPrefix(lineTmp, comment+"type:") {
			snippet[key].Type = strings.TrimPrefix(lineTmp, comment+"type:")
			continue
		}
		if len(lineTmp) > 0 && strings.HasPrefix(lineTmp, comment+"var:") {
			value := strings.Split(strings.TrimPrefix(lineTmp, comment+"var:"), "=")
			if len(value) >= 2 {
				value[1] = strings.Split(strings.TrimPrefix(line, comment+"var:"), value[0]+"=")[1]
			}
			if len(value) == 1 {
				value = append(value, "")
			}
			snippet[key].Vars[value[0]] = &Var{Key: value[0], Value: value[1], Type: varType}
			continue
		}
		if len(lineTmp) > 0 && strings.HasPrefix(lineTmp, comment+"var_type:") {
			varType = strings.TrimPrefix(lineTmp, comment+"var_type:")
			continue
		}
		varType = "string"
		snippet[key].Text += line + "\n"
	}

	//fmt.Println(Snippet)

	return &Snippets{
		snippet: snippet,
		file:    name,
	}
}

func LoadFile(name string) *Snippets {
	return LoadFileComment(name, "--", template.HTMLEscapeString)
}

func (g *Snippets) GetSnippet(name string) Snippet {
	return *g.snippet[name]
}

func (q *Snippet) Escape(f func(string) string) *Snippet {
	q.FxEscape = f
	return q
}

func (q *Snippet) Param(key string, value interface{}) *Snippet {

	str := fmt.Sprint(value)
	if q.Vars[key].Type == "int" {
		if _, err := strconv.Atoi(str); err != nil {
			log.Println(err)
			value = ""
		}
	}
	if q.Vars[key].Type == "float" {
		if _, err := strconv.ParseFloat(str, 64); err != nil {
			log.Println(err)
			value = ""
		}
	}
	if q.Vars[key].Type == "bool" {
		if _, err := strconv.ParseBool(str); err != nil {
			log.Println(err)
			value = ""
		}
	}
	/*
		if q.Vars[key].Type == "string" {

		}
	*/

	q.Vars[key].Value = str
	return q

}

func (q *Snippet) Get() string {
	t := q.Text
	flagEscape := q.FxEscape != nil
	for k, v := range q.Vars {
		if flagEscape {
			t = strings.ReplaceAll(t, "${{"+k+"}}", q.FxEscape(v.Value))
			continue
		}
		t = strings.ReplaceAll(t, "${{"+k+"}}", v.Value)
	}
	return t
}
