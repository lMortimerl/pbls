package parser

import (
	"fmt"
	"pbls/src/ast"
	"pbls/src/lexer"
)

type parser struct {
	tokens  []lexer.Token
	current int
}

func NewParser(tokens []lexer.Token) *parser {
	createTokenLookups()
	createTypeTokenLookups()
	return &parser{
		tokens: tokens,
	}
}

func Parse(tokens []lexer.Token) ast.BlockStmt {
	body := make([]ast.Stmt, 0)
	parser := NewParser(tokens)

	for parser.hasTokens() {
		body = append(body, parse_stmt(parser))
	}

	return ast.BlockStmt{
		Body: body,
	}
}

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.current]
}
func (p *parser) advance() lexer.Token {
	tkn := p.currentToken()
	p.current++
	return tkn
}
func (p *parser) peek() lexer.Token {
	tkn := p.tokens[p.current+2]
	return tkn
}
func (p *parser) hasTokens() bool {
	return p.current < len(p.tokens) && p.currentToken().Kind != lexer.EOF
}
func (p *parser) expectError(expectedKind lexer.TokenKind, err any) lexer.Token {
	token := p.currentToken()
	kind := token.Kind

	if kind != expectedKind {
		if err == nil {
			err = fmt.Sprintf("Expected %s but recieved %s\n", lexer.TokenKindString(expectedKind), lexer.TokenKindString(kind))
		}

		panic(err)
	}

	return p.advance()
}
func (p *parser) expectOneOf(expected ...lexer.TokenKind) lexer.Token {
	token := p.currentToken()
	kind := token.Kind
	var joined string
	for _, expectedKind := range expected {
		joined += lexer.TokenKindString(expectedKind) + " "
		if expectedKind == kind {
			return p.advance()
		}
	}
	panic(fmt.Sprintf("Expected one of '%s' but got '%s'", joined, lexer.TokenKindString(kind)))
}
func (p *parser) expect(expectedKind lexer.TokenKind) lexer.Token {
	return p.expectError(expectedKind, nil)
}
