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
	var isConstant bool = false
	var varValue ast.Expr
	var possibleType ast.Type
	if p.currentToken().Kind == lexer.CONSTANT {
		p.advance()
		isConstant = true
	}
	possibleType = parse_type(p, default_bp)
	varName := p.advance()
	if p.advance().Kind == lexer.EQUALS {
		varValue = parse_expr(p, default_bp)
	}
	p.expectError(lexer.SEMICOLON, "Expected Semicolon")

	return ast.VarDeclStmt{
		Identifier:    varName.Value,
		IsConstant:    isConstant,
		AssignedValue: varValue,
		ExplicitType:  possibleType,
	}
	/*
	   var explicitType ast.Type
	   var assignedValue ast.Expr
	   isConstant := p.advance().Kind == lexer.CONSTANT
	   varName := p.expectError(lexer.IDENTIFIER, "Expected varname after constant").Value

	   if p.currentToken().Kind == lexer.COLON {
	       p.advance()
	       explicitType = parse_type(p, default_bp)
	   }

	   if p.currentToken().Kind != lexer.SEMICOLON {
	       p.expectError(lexer.EQUALS, "Expected = after varname")
	       assignedValue = parse_expr(p, assignment)
	   } else if explicitType == nil {
	       panic("Missing either righthand side in var declaration or explicit type.")
	   }

	   p.expectError(lexer.SEMICOLON, "Expected Semicolon!")

	   if isConstant && assignedValue == nil {
	       panic("Cannot define Constant without providing value")
	   }

	   return ast.VarDeclStmt{
	       Identifier:    varName,
	       IsConstant:    isConstant,
	       AssignedValue: assignedValue,
	       ExplicitType:  explicitType,
	   }
	*/
}
