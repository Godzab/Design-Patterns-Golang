package main

import (
	"fmt"
	"strings"
)

const (
	indent = 2
)

type HtmlElement struct{
	name, text string
	elements []HtmlElement
}

type HtmlBuilder struct{
	root HtmlElement
	rootName string
}

func NewHtmlBuilder(rootName string)*HtmlBuilder{
	return &HtmlBuilder{rootName: rootName, root: HtmlElement{}}
}

func (e *HtmlElement) String(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indent * indent)
	sb.WriteString(fmt.Sprintf("%d<%s>\n",
		i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ",
			indent * (indent + 1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.String(indent+1))
	}
	return sb.String()
}


func (b *HtmlBuilder) String() string{
	return b.root.String(indent)
}

//AddChildFluent is a fluent function which allows chaining
func (b *HtmlBuilder) AddChildFluent(childName, childText string)*HtmlBuilder{
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

func main(){
	//Complex implementation

	hello := "Hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("<p>")
	fmt.Println(sb.String())

	words := []string{"hello", "world"}
	sb.Reset()
	sb.WriteString("<ul>")
	for _, v := range words{
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	// Simpler implementation

	b := NewHtmlBuilder("ul")
	b.AddChildFluent("li", "hello")
	b.AddChildFluent("li", "world")
	fmt.Println(b.String())
}
