package parser

import (
	"pbls/src/ast"
	"pbls/src/lexer"
)

func parse_stmt(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentToken().Kind]

	if exists {
		return stmt_fn(p)
	}
	expression := parse_expr(p, default_bp)
	p.expect(lexer.SEMICOLON)
	return ast.ExprStmt{
		Expr: expression,
	}
}
