package lexer

import (
	"fmt"
	"regexp"
)

type regexHandler func(lex *Lexer, regex *regexp.Regexp)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type Lexer struct {
	patterns []regexPattern
	Tokens   []Token
	source   string
	current  int
	line     int
	column   int
}

func (l *Lexer) advanceN(n int) {
	l.current += n
	l.column += n
	if l.current < len(l.source) {
		curRune := rune(l.source[l.current])
		if curRune == '\n' {
			l.line++
			l.column = 1
		}
	}
}
func (l *Lexer) push(t Token) {
	l.Tokens = append(l.Tokens, t)
}
func (l *Lexer) at() string {
	return string(l.source[l.current])
}
func (l *Lexer) remainder() string {
	return l.source[l.current:]
}
func (l *Lexer) atEOF() bool {
	return l.current >= len(l.source)
}
func Tokenize(source []byte) []Token {
	lex := NewLexer(source)

	for !lex.atEOF() {
		matched := false

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(string(lex.remainder()))

			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		if !matched {
			panic(fmt.Sprintf("Lexer::Error -> unrecoginized token at line %d column %d near %s\n", lex.line, lex.column, string(lex.remainder())))
		}
	}
	lex.push(NewToken(EOF, "EOF", lex.line, 0))

	return lex.Tokens
}

func numberHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindString(string(lex.remainder()))
	lex.push(NewToken(NUMBER, match, lex.line, lex.column))
	lex.advanceN(len(match))
}
func skipHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(string(lex.remainder()))
	lex.advanceN(match[1])
}
func stringHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	stringLiteral := lex.remainder()[match[0]+1 : match[1]-1]

	lex.push(NewToken(STRING, stringLiteral, lex.line, lex.column))
	lex.advanceN(len(stringLiteral) + 2)
}
func symbolHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())

	if kind, exists := reserved_lu[match]; exists {
		lex.push(NewToken(kind, match, lex.line, lex.column))
	} else {
		lex.push(NewToken(IDENTIFIER, match, lex.line, lex.column))
	}

	lex.advanceN(len(match))
}
func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *Lexer, regex *regexp.Regexp) {
		lex.push(NewToken(kind, value, lex.line, lex.column))
		lex.advanceN(len(value))
	}
}
func newLineHandler(kind TokenKind, value string) regexHandler {
	return func(lex *Lexer, regex *regexp.Regexp) {
		lex.push(NewToken(kind, value, lex.line, lex.column))
		lex.advanceN(len(value))
		lex.line++
		lex.column = 1
	}
}
func NewLexer(source []byte) *Lexer {
	return &Lexer{
		source: string(source),
		line:   1,
		column: 1,
		patterns: []regexPattern{
			{regexp.MustCompile(`string`), defaultHandler(IDENTIFIER_TYPE, "string")},
			{regexp.MustCompile(`long`), defaultHandler(IDENTIFIER_TYPE, "long")},
			{regexp.MustCompile(`int`), defaultHandler(IDENTIFIER_TYPE, "int")},
			{regexp.MustCompile(`char`), defaultHandler(IDENTIFIER_TYPE, "char")},

			{regexp.MustCompile(`[a-zA-Z_][a-zA-Z_0-9]*`), symbolHandler},
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			{regexp.MustCompile(`"[^"]*"`), stringHandler},
			{regexp.MustCompile(`'[^']*'`), stringHandler},
			{regexp.MustCompile(`\/\/.*`), skipHandler},
			{regexp.MustCompile(`\/\*[\s\S]*?\*\/`), skipHandler},
			{regexp.MustCompile(`[ \t]+`), skipHandler},
			{regexp.MustCompile(`\n`), defaultHandler(NEWLINE, "n")},
			{regexp.MustCompile(`\r\n`), defaultHandler(NEWLINE, "rn")},

			{regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			{regexp.MustCompile(`\]`), defaultHandler(CLOSE_BRACKET, "]")},
			{regexp.MustCompile(`\{`), defaultHandler(OPEN_CURLY, "{")},
			{regexp.MustCompile(`\}`), defaultHandler(CLOSE_CURLY, "}")},
			{regexp.MustCompile(`\(`), defaultHandler(OPEN_PAREN, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(CLOSE_PAREN, ")")},
			{regexp.MustCompile(`=`), defaultHandler(EQUALS, "=")},
			{regexp.MustCompile(`!=`), defaultHandler(NOT_EQUALS, "!=")},
			{regexp.MustCompile(`!`), defaultHandler(NOT, "!")},
			{regexp.MustCompile(`>=`), defaultHandler(GREATER_EQUAL, ">=")},
			{regexp.MustCompile(`>`), defaultHandler(GREATER, ">")},
			{regexp.MustCompile(`<=`), defaultHandler(LESS_EQUAL, "<=")},
			{regexp.MustCompile(`<`), defaultHandler(LESS, "<")},
			{regexp.MustCompile(`\.`), defaultHandler(DOT, ".")},
			{regexp.MustCompile(`;`), defaultHandler(SEMICOLON, ";")},
			{regexp.MustCompile(`:`), defaultHandler(COLON, ":")},
			{regexp.MustCompile(`\?`), defaultHandler(QUESTION, "?")},
			{regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
			{regexp.MustCompile("`"), defaultHandler(BACKTICK, "`")},
			{regexp.MustCompile(`\+\+`), defaultHandler(PLUS_PLUS, "++")},
			{regexp.MustCompile(`--`), defaultHandler(MINUS_MINUS, "--")},
			{regexp.MustCompile(`\+=`), defaultHandler(PLUS_EQUALS, "+=")},
			{regexp.MustCompile(`-=`), defaultHandler(MINUS_EQUALS, "-=")},
			{regexp.MustCompile(`/=`), defaultHandler(SLASH_EQUALS, "/=")},
			{regexp.MustCompile(`\*=`), defaultHandler(STAR_EQUALS, "*=")},
			{regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`-`), defaultHandler(MINUS, "-")},
			{regexp.MustCompile(`/`), defaultHandler(SLASH, "/")},
			{regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
			{regexp.MustCompile(`%`), defaultHandler(PERCENT, "%")},
		},
	}
}
