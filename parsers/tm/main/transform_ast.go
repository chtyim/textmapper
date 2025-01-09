package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/inspirer/textmapper/parsers/tm"
	"github.com/inspirer/textmapper/parsers/tm/ast"
	"github.com/inspirer/textmapper/parsers/tm/selector"
)

func main() {
	//var s TokenStream
	//var p Parser
	//
	ctx := context.Background()
	//
	path := "/Users/terencey/Works/bqparser/tm/zetasql_ts.tm"
	source, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	tree, err := ast.Parse(ctx, path, string(source), func(err tm.SyntaxError) bool {
		fmt.Printf("Error %+v\n", err)
		return false
	})
	if err != nil {
		panic(err)
	}

	children := tree.Root().Children(selector.Any)
	for _, c := range children {
		if c.Type() == tm.ParserSection {
			fmt.Printf(":: parser\n\n")
			for _, p := range c.Children(selector.Any) {
				text := p.Text()
				if p.Type() == tm.Nonterm {
					if reportClause := p.Child(selector.ReportClause); reportClause == nil {
						idx := strings.Index(text, ":")
						//nodeName := ToJavaClassName()
						firstChild := p.Child(selector.Any)
						if firstChild.Type() != tm.Identifier {
							panic(errors.New("Nonterm should starts with identifier"))
						}
						nodeName := ToJavaClassName(firstChild.Text())
						text = fmt.Sprintf("%v -> %v%v", text[:idx], nodeName, text[idx:])
					}
				}
				fmt.Printf("%v\n", text)
			}
		} else {
			fmt.Printf("%v\n", c.Text())
		}
	}
}

func ToJavaClassName(s string) string {
	// 1. Remove invalid characters:
	reg := regexp.MustCompile("[^a-zA-Z0-9]+") // Keep only letters, numbers, and underscores
	s = reg.ReplaceAllString(s, " ")           // Replace invalid characters with spaces

	// 2. Trim leading/trailing spaces and convert to title case:
	s = strings.TrimSpace(s)
	s = strings.Title(s) // Capitalize the first letter of each word

	// 3. Remove spaces to create UpperCamelCase:
	s = strings.ReplaceAll(s, " ", "")

	// 4. Handle cases where the first character is not a letter:
	if len(s) > 0 && !unicode.IsLetter(rune(s[0])) {
		s = "_" + s // Prepend an underscore if it starts with a number
	}

	// 5. Handle edge cases where the string is empty after processing:
	if len(s) == 0 {
		return "DefaultClassName" // Provide a default name (you can customize this)
	}

	return s
}
