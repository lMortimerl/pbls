package parser_test

import (
	"fmt"
	"os"
	"pbls/src/ast"
	"pbls/src/lexer"
	"pbls/src/parser"
	"reflect"
	"testing"
)

func readFile(filename string) []byte {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Could not read the input file!\n%s", err))
	}
	return content
}

func compareAst(t *testing.T, errMsg string, expected, actual ast.BlockStmt) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Error: %s", errMsg)
	}
}

func parse(statement string) ast.BlockStmt {
	return parser.Parse(lexer.Tokenize([]byte(statement)))
}

func TestSimpleVariableDeclaration(t *testing.T) {
	expected := ast.BlockStmt{
		Body: []ast.Stmt{ast.VarDeclStmt{
			Identifier:    "ls_string",
			IsConstant:    false,
			AssignedValue: ast.StringExpr{Value: "A B C"},
			ExplicitType: ast.SymbolType{
				Name: "string",
			},
		}},
	}
	actual := parse("string ls_string = \"A B C\";")
	compareAst(t, "Cannot parse simple string declaration!", expected, actual)
}
func TestMultiVariableDeclaration(t *testing.T) {
	expected := ast.BlockStmt{
		Body: []ast.Stmt{
			ast.MultiVarDeclStmt{
				Stmts: []ast.VarDeclStmt{
					{
						Identifier:    "ls_string1",
						IsConstant:    false,
						AssignedValue: nil,
						ExplicitType: ast.SymbolType{
							Name: "string",
						},
					},
					{
						Identifier:    "ls_string2",
						IsConstant:    false,
						AssignedValue: nil,
						ExplicitType: ast.SymbolType{
							Name: "string",
						},
					},
				},
			},
		},
	}
	actual := parse("string ls_string1, ls_string2;")
	compareAst(t, "Cannot parse multi Variable declaration!", expected, actual)
}
func TestVarDeclExprStmt(t *testing.T) {
	expected := ast.BlockStmt{
		Body: []ast.Stmt{
			ast.VarDeclStmt{
				Identifier: "ls_str",
				IsConstant: false,
				AssignedValue: ast.BinaryExpr{
					Left:     ast.StringExpr{Value: "A"},
					Operator: lexer.Token{Kind: lexer.PLUS, Value: "+", Line: 1, Column: 20},
					Right:    ast.StringExpr{Value: "B"},
				},
				ExplicitType: ast.SymbolType{Name: "string"},
			},
		},
	}
	actual := parse("string ls_str = \"A\"+\"B\";")
	compareAst(t, "Cannot parse Variable declaration with simple expression!", expected, actual)
}
