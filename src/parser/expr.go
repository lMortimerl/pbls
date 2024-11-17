package parser

import (
	"fmt"
	"pbls/src/ast"
	"pbls/src/lexer"
	"strconv"
)

func parse_expr(p *parser, bp BindingPower) ast.Expr {
	tokenKind := p.currentToken().Kind
	nud_fn, exists := nud_lu[tokenKind]
	if !exists {
		panic(fmt.Sprintf("NUD HANDLER EXPECTED FOR TOKEN %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)
	for bp_lu[p.currentToken().Kind] > bp {
		tokenKind := p.currentToken().Kind
		led_fn, exists := led_lu[tokenKind]
		if !exists {
			panic(fmt.Sprintf("LED HANDLER EXPECTED FOR TOKEN %s\n", lexer.TokenKindString(tokenKind)))
		}
		left = led_fn(p, left, bp)
	}

	return left
}

func parse_binary_expr(p *parser, left ast.Expr, bp BindingPower) ast.Expr {
	operatorToken := p.advance()
	right := parse_expr(p, bp)
	return ast.BinaryExpr{
		Left:     left,
		Operator: operatorToken,
		Right:    right,
	}
}
func parse_primary_expr(p *parser) ast.Expr {
	switch p.currentToken().Kind {
	case lexer.NUMBER:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.NumberExpr{
			Value: number,
		}
	case lexer.STRING:
		return ast.StringExpr{
			Value: p.advance().Value,
		}
	case lexer.IDENTIFIER:
		return ast.SymbolExpr{
			Value: p.advance().Value,
		}
	default:
		panic(fmt.Sprintf("Cannot create primary_expression from %s\n", lexer.TokenKindString(p.currentToken().Kind)))
	}
}
