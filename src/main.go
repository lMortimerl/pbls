package main

import (
	"fmt"
	"os"
	"pbls/src/lexer"
	"pbls/src/parser"

	"github.com/sanity-io/litter"
)

func main() {
	var sourceFilename string = ".\\examples\\06.lang"
	content, err := os.ReadFile(sourceFilename)
	if err != nil {
		fmt.Printf("Could not read the input file!\n%s", err)
	}
	fmt.Printf("Read %d bytes from %s\n", len(content), sourceFilename)
	tokens := lexer.Tokenize(content)
	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
