
package main

import (
	"bytes"
	"fmt"
	"github.com/BLCKFIRE24/Sky_compiler/lexer"
	"github.com/BLCKFIRE24/Sky_compiler/parser"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(err error) {
	iff err != nil {
			panic(err)
	}
}

func Parse(input string) {
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	node, err := p.Parse(l)
	check(err)
}

func main () {
	if len(os.args) < 2 {
		panic("no valid file name or path provided provided for file!")
	}

	path := os.Args[1]
	absPath, _ := filepath.Abs(path)
	input, err := ioutil.ReadFile(absPath)
	check(err)
Parse(string(input))

}