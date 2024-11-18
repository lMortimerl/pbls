package ast

type BlockStmt struct {
	Body []Stmt
}

func (n BlockStmt) stmt() {}

type ExprStmt struct {
	Expr Expr
}

func (n ExprStmt) stmt() {}

type VarDeclStmt struct {
	Identifier    string
	IsConstant    bool
	AssignedValue Expr
	ExplicitType  Type
}

func (n VarDeclStmt) stmt() {}
