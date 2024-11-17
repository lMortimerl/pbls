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

func parse_var_decl_stmt(p *parser) ast.Stmt {
	isConstant := p.advance().Kind == lexer.CONSTANT
	varName := p.expectError(lexer.IDENTIFIER, "Expected varname after constant").Value

	p.expectError(lexer.EQUALS, "Expected = after varname")
	assignedValue := parse_expr(p, assignment)
	p.expectError(lexer.SEMICOLON, "Expected Semicolon!")

	return ast.VarDeclStmt{
		Identifier:    varName,
		IsConstant:    isConstant,
		AssignedValue: assignedValue,
	}
}
