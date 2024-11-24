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
	p.expectOneOf(lexer.NEWLINE, lexer.SEMICOLON)
	return ast.ExprStmt{
		Expr: expression,
	}
}

func parse_comma_separated_declaration(p *parser, t ast.Type, varList []ast.VarDeclStmt) ast.Stmt {
	for {
		currTkn := p.advance()
		if currTkn.Kind == lexer.COMMA {
			continue
		} else if currTkn.Kind == lexer.NEWLINE || currTkn.Kind == lexer.SEMICOLON {
			break
		}
		varList = append(varList, ast.VarDeclStmt{
			Identifier:    currTkn.Value,
			IsConstant:    false,
			AssignedValue: nil,
			ExplicitType:  t,
		})
	}
	p.advance()
	return ast.MultiVarDeclStmt{
		Stmts: varList,
	}
}

func parse_var_decl_stmt(p *parser) ast.Stmt {
	var isConstant bool = false
	var currTkn lexer.Token = p.currentToken()
	var varType ast.Type

	if currTkn.Kind == lexer.CONSTANT {
		isConstant = true
		currTkn = p.advance()
	}
	varType = parse_type(p, default_bp)
	varName := p.advance()
	var declaration ast.VarDeclStmt = ast.VarDeclStmt{
		Identifier:    varName.Value,
		IsConstant:    isConstant,
		AssignedValue: nil,
		ExplicitType:  varType,
	}
	currTkn = p.currentToken()
	if currTkn.Kind == lexer.NEWLINE || currTkn.Kind == lexer.SEMICOLON {
		return declaration
	} else if currTkn.Kind == lexer.COMMA {
		return parse_comma_separated_declaration(p, varType, []ast.VarDeclStmt{declaration})
	} else if currTkn.Kind == lexer.EQUALS {
		p.advance()
		declaration.AssignedValue = parse_expr(p, default_bp)
	}
	p.expectOneOf(lexer.NEWLINE, lexer.SEMICOLON)

	return declaration
}

func _parse_var_decl_stmt(p *parser) ast.Stmt {
	var isConstant bool = false
	var currTkn lexer.Token = p.currentToken()
	var varType ast.Type
	var varValue ast.Expr

	if currTkn.Kind == lexer.CONSTANT {
		isConstant = true
		currTkn = p.advance()
	}
	varType = parse_type(p, default_bp)
	var varName = p.advance()

	var declaration ast.VarDeclStmt = ast.VarDeclStmt{
		Identifier:    varName.Value,
		IsConstant:    isConstant,
		AssignedValue: nil,
		ExplicitType:  varType,
	}
	currTkn = p.advance()
	if currTkn.Kind == lexer.COMMA {
		return parse_comma_separated_declaration(p, varType, []ast.VarDeclStmt{declaration})
	} else if currTkn.Kind == lexer.EQUALS {
		varValue = parse_expr(p, default_bp)
	} else if isConstant && p.currentToken().Kind == lexer.OPEN_BRACKET {
		panic("PowerBuilder does not allow the use of constant arrays!")
	} else if isConstant {
		panic("Constants need to be initialized!")
	}
	p.expectOneOf(lexer.SEMICOLON, lexer.NEWLINE)
	declaration.AssignedValue = varValue
	return declaration
}
