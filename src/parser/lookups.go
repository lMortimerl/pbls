package parser

import (
	"pbls/src/ast"
	"pbls/src/lexer"
)

type BindingPower int

const (
	default_bp BindingPower = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type stmt_handler func(p *parser) ast.Stmt
type nud_handler func(p *parser) ast.Expr
type led_handler func(p *parser, left ast.Expr, bp BindingPower) ast.Expr

type stmt_lookup map[lexer.TokenKind]stmt_handler
type nud_lookup map[lexer.TokenKind]nud_handler
type led_lookup map[lexer.TokenKind]led_handler
type bp_lookup map[lexer.TokenKind]BindingPower

var bp_lu = bp_lookup{}
var nud_lu = nud_lookup{}
var led_lu = led_lookup{}
var stmt_lu = stmt_lookup{}

func led(kind lexer.TokenKind, bp BindingPower, led_fn led_handler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}
func nud(kind lexer.TokenKind, bp BindingPower, nud_fn nud_handler) {
	bp_lu[kind] = bp
	nud_lu[kind] = nud_fn
}
func stmt(kind lexer.TokenKind, bp BindingPower, stmt_fn stmt_handler) {
	bp_lu[kind] = bp
	stmt_lu[kind] = stmt_fn
}
func createTokenLookups() {
	// Logical
	led(lexer.AND, logical, parse_binary_expr)
	led(lexer.OR, logical, parse_binary_expr)

	// Relational
	led(lexer.LESS, relational, parse_binary_expr)
	led(lexer.GREATER, relational, parse_binary_expr)
	led(lexer.LESS_EQUAL, relational, parse_binary_expr)
	led(lexer.GREATER_EQUAL, relational, parse_binary_expr)
	led(lexer.EQUALS, relational, parse_binary_expr)
	led(lexer.NOT_EQUALS, relational, parse_binary_expr)

	// Additive
	led(lexer.PLUS, additive, parse_binary_expr)
	led(lexer.MINUS, additive, parse_binary_expr)

	// Multiplicative
	led(lexer.STAR, multiplicative, parse_binary_expr)
	led(lexer.SLASH, multiplicative, parse_binary_expr)
	led(lexer.PERCENT, multiplicative, parse_binary_expr)

	// Literals & Symbols
	nud(lexer.NUMBER, primary, parse_primary_expr)
	nud(lexer.STRING, primary, parse_primary_expr)
	nud(lexer.IDENTIFIER, primary, parse_primary_expr)
}