package main

import (
	"fmt"
	"os"
	"pbls/src/lexer"
)

func main() {
	var sourceFilename string = ".\\examples\\00.lang"
	content, err := os.ReadFile(sourceFilename)
	if err != nil {
		fmt.Printf("Could not read the input file!\n%s", err)
	}
	fmt.Printf("Read %d bytes from %s\n", len(content), sourceFilename)
	tokens := lexer.Tokenize(content)
	for _, tkn := range tokens {
		tkn.Debug()
	}
}
