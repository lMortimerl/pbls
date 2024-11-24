package lexer_test

import (
	"pbls/src/lexer"
	"testing"
)

func compareTokens(t *testing.T, expected, actual []lexer.Token) {
	if len(actual) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(actual))
	}
	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("at index %d: expected token %v, got %v", i, expected[i], actual[i])
		}
	}
}

func TestSingleTokens(t *testing.T) {
	input := `{ ( ) }`
	expected := []lexer.Token{
		{Kind: lexer.OPEN_CURLY, Value: "{", Line: 1, Column: 1},
		{Kind: lexer.OPEN_PAREN, Value: "(", Line: 1, Column: 3},
		{Kind: lexer.CLOSE_PAREN, Value: ")", Line: 1, Column: 5},
		{Kind: lexer.CLOSE_CURLY, Value: "}", Line: 1, Column: 7},
		{Kind: lexer.EOF, Value: "EOF", Line: 1, Column: 0},
	}

	tokens := lexer.Tokenize([]byte(input))

	compareTokens(t, expected, tokens)
}

func TestReservedKeywords(t *testing.T) {
	input := `if else while true false`
	expected := []lexer.Token{
		{Kind: lexer.IF, Value: "if", Line: 1, Column: 1},
		{Kind: lexer.ELSE, Value: "else", Line: 1, Column: 4},
		{Kind: lexer.WHILE, Value: "while", Line: 1, Column: 9},
		{Kind: lexer.TRUE, Value: "true", Line: 1, Column: 15},
		{Kind: lexer.FALSE, Value: "false", Line: 1, Column: 20},
		{Kind: lexer.EOF, Value: "EOF", Line: 1, Column: 0},
	}

	tokens := lexer.Tokenize([]byte(input))

	compareTokens(t, expected, tokens)
}

func TestIdentifiersAndNumbers(t *testing.T) {
	input := `x1 varName 123 45.67`
	expected := []lexer.Token{
		{Kind: lexer.IDENTIFIER, Value: "x1", Line: 1, Column: 1},
		{Kind: lexer.IDENTIFIER, Value: "varName", Line: 1, Column: 4},
		{Kind: lexer.NUMBER, Value: "123", Line: 1, Column: 12},
		{Kind: lexer.NUMBER, Value: "45.67", Line: 1, Column: 16},
		{Kind: lexer.EOF, Value: "EOF", Line: 1, Column: 0},
	}

	tokens := lexer.Tokenize([]byte(input))

	compareTokens(t, expected, tokens)
}

func TestStrings(t *testing.T) {
	input := `"hello" "world"`
	expected := []lexer.Token{
		{Kind: lexer.STRING, Value: "hello", Line: 1, Column: 1},
		{Kind: lexer.STRING, Value: "world", Line: 1, Column: 9},
		{Kind: lexer.EOF, Value: "EOF", Line: 1, Column: 0},
	}

	tokens := lexer.Tokenize([]byte(input))

	compareTokens(t, expected, tokens)
}

func TestMixedInput(t *testing.T) {
	input := "if (x > 10)\n{ return 'done'; }\r\n"
	expected := []lexer.Token{
		{Kind: lexer.IF, Value: "if", Line: 1, Column: 1},
		{Kind: lexer.OPEN_PAREN, Value: "(", Line: 1, Column: 4},
		{Kind: lexer.IDENTIFIER, Value: "x", Line: 1, Column: 5},
		{Kind: lexer.GREATER, Value: ">", Line: 1, Column: 7},
		{Kind: lexer.NUMBER, Value: "10", Line: 1, Column: 9},
		{Kind: lexer.CLOSE_PAREN, Value: ")", Line: 1, Column: 11},
		{Kind: lexer.NEWLINE, Value: `n`, Line: 2, Column: 1},
		{Kind: lexer.OPEN_CURLY, Value: "{", Line: 2, Column: 2},
		{Kind: lexer.RETURN, Value: "return", Line: 2, Column: 4},
		{Kind: lexer.STRING, Value: "done", Line: 2, Column: 11},
		{Kind: lexer.SEMICOLON, Value: ";", Line: 2, Column: 17},
		{Kind: lexer.CLOSE_CURLY, Value: "}", Line: 2, Column: 19},
		{Kind: lexer.NEWLINE, Value: `rn`, Line: 2, Column: 20},
		{Kind: lexer.EOF, Value: "EOF", Line: 2, Column: 0},
	}

	tokens := lexer.Tokenize([]byte(input))

	compareTokens(t, expected, tokens)
}
