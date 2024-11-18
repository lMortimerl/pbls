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
		panic(fmt.Sprintf("1-NUD HANDLER EXPECTED FOR TOKEN %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)
	for bp_lu[p.currentToken().Kind] > bp {
		tokenKind := p.currentToken().Kind
		led_fn, exists := led_lu[tokenKind]
		if !exists {
			panic(fmt.Sprintf("2-LED HANDLER EXPECTED FOR TOKEN %s\n", lexer.TokenKindString(tokenKind)))
		}
		left = led_fn(p, left, bp_lu[p.currentToken().Kind])
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
func parse_assignment_expr(p *parser, left ast.Expr, bp BindingPower) ast.Expr {
	operator := p.advance()
	rhs := parse_expr(p, bp)
	return ast.AssignmentExpr{
		Assigne:  left,
		Operator: operator,
		Value:    rhs,
	}
}
func parse_prefix_expr(p *parser) ast.Expr {
	operatorToken := p.advance()
	rhs := parse_expr(p, default_bp)

	return ast.PrefixExpr{
		Operator: operatorToken,
		Value:    rhs,
	}
}
func parse_grouping_expr(p *parser) ast.Expr {
	p.advance()
	expr := parse_expr(p, default_bp)
	p.expect(lexer.CLOSE_PAREN)
	return expr
}

// TODO: ComputedExpression => array[index] | array[expr]
// TODO: XX => const foo = [1, 2, 3]
