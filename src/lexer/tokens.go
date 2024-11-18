package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	NUMBER
	STRING
	IDENTIFIER

	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	EQUALS
	NOT
	NOT_EQUALS

	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	DOT
	SEMICOLON
	COLON
	QUESTION
	COMMA
	BACKTICK

	PLUS_PLUS
	MINUS_MINUS
	PLUS_EQUALS
	MINUS_EQUALS
	SLASH_EQUALS
	STAR_EQUALS
	PERCENT_EQUALS

	PLUS
	MINUS
	SLASH
	STAR
	PERCENT

	ALIAS
	AND
	AUTOINSTANCIATE
	CALL
	CASE
	CATCH
	CHOOSE
	CLOSE
	COMMIT
	CONNECT
	CONSTANT
	CONTINUE
	CREATE
	CURSOR
	DECLARE
	DELETE
	DESCRIBE
	DESCRIPTOR
	DESTROY
	DISCONNECT
	DO
	DYNAMIC
	ELSE
	ELSEIF
	END
	ENUMERATED
	EVENT
	EXECUTE
	EXIT
	EXTERNAL
	FALSE
	FETCH
	FINALLY
	FIRST
	FOR
	FORWARD
	FROM
	FUNCTION
	GLOBAL
	GOTO
	HALT
	IF
	IMMEDIATE
	INDIRECT
	INSERT
	INTO
	INTRINSCI
	IS
	LAST
	LIBRARY
	LOOP
	NAMESPACE
	NATIVE
	NEXT
	NOTOF
	ON
	OPEN
	OR
	PARENT
	POST
	PREPARE
	PRIOR
	PRIVATE
	PRIVATEREAD
	PRIVATEWRITE
	PROCEDURE
	PROTECTED
	PROTECTEDREAD
	PROTECTEDWRITE
	PROTOTYPES
	PUBLIC
	READONLY
	REF
	RETURN
	ROLLBACK
	RPCFUNC
	SELECT
	SELECTBLOB
	SHARED
	STATIC
	STEP
	SUBROUTINE
	SUPER
	SYSTEM
	SYSTEMREAD
	SYSTEMWRITE
	THEN
	THIS
	THROW
	THROWS
	TO
	TRIGGER
	TRUE
	TRY
	TYPE
	UNTIL
	UPDATE
	UPDATEBLOB
	USING
	VARIABLES
	WHILE
	WITH
	WITHIN
	XOR
	_DEBUG
)

var reserved_lu map[string]TokenKind = map[string]TokenKind{
	"alias":           ALIAS,
	"and":             AND,
	"autoinstanciate": AUTOINSTANCIATE,
	"call":            CALL,
	"case":            CASE,
	"catch":           CATCH,
	"choose":          CHOOSE,
	"close":           CLOSE,
	"commit":          COMMIT,
	"connect":         CONNECT,
	"constant":        CONSTANT,
	"continue":        CONTINUE,
	"create":          CREATE,
	"cursor":          CURSOR,
	"declare":         DECLARE,
	"delete":          DELETE,
	"describe":        DESCRIBE,
	"descriptor":      DESCRIPTOR,
	"destroy":         DESTROY,
	"disconnect":      DISCONNECT,
	"do":              DO,
	"dynamic":         DYNAMIC,
	"else":            ELSE,
	"elseif":          ELSEIF,
	"end":             END,
	"enumerated":      ENUMERATED,
	"event":           EVENT,
	"execute":         EXECUTE,
	"exit":            EXIT,
	"external":        EXTERNAL,
	"false":           FALSE,
	"fetch":           FETCH,
	"finally":         FINALLY,
	"first":           FIRST,
	"for":             FOR,
	"forward":         FORWARD,
	"from":            FROM,
	"function":        FUNCTION,
	"global":          GLOBAL,
	"goto":            GOTO,
	"halt":            HALT,
	"if":              IF,
	"immediate":       IMMEDIATE,
	"indirect":        INDIRECT,
	"insert":          INSERT,
	"into":            INTO,
	"intrinsci":       INTRINSCI,
	"is":              IS,
	"last":            LAST,
	"library":         LIBRARY,
	"loop":            LOOP,
	"namespace":       NAMESPACE,
	"native":          NATIVE,
	"next":            NEXT,
	"notof":           NOTOF,
	"on":              ON,
	"open":            OPEN,
	"or":              OR,
	"parent":          PARENT,
	"post":            POST,
	"prepare":         PREPARE,
	"prior":           PRIOR,
	"private":         PRIVATE,
	"privateread":     PRIVATEREAD,
	"privatewrite":    PRIVATEWRITE,
	"procedure":       PROCEDURE,
	"protected":       PROTECTED,
	"protectedread":   PROTECTEDREAD,
	"protectedwrite":  PROTECTEDWRITE,
	"prototypes":      PROTOTYPES,
	"public":          PUBLIC,
	"readonly":        READONLY,
	"ref":             REF,
	"return":          RETURN,
	"rollback":        ROLLBACK,
	"rpcfunc":         RPCFUNC,
	"select":          SELECT,
	"selectblob":      SELECTBLOB,
	"shared":          SHARED,
	"static":          STATIC,
	"step":            STEP,
	"subroutine":      SUBROUTINE,
	"super":           SUPER,
	"system":          SYSTEM,
	"systemread":      SYSTEMREAD,
	"systemwrite":     SYSTEMWRITE,
	"then":            THEN,
	"this":            THIS,
	"throw":           THROW,
	"throws":          THROWS,
	"to":              TO,
	"trigger":         TRIGGER,
	"true":            TRUE,
	"try":             TRY,
	"type":            TYPE,
	"until":           UNTIL,
	"update":          UPDATE,
	"updateblob":      UPDATEBLOB,
	"using":           USING,
	"variables":       VARIABLES,
	"while":           WHILE,
	"with":            WITH,
	"within":          WITHIN,
	"xor":             XOR,
	"_debug":          _DEBUG,
}

type Token struct {
	Kind   TokenKind
	Value  string
	Line   int
	Column int
}

func NewToken(kind TokenKind, value string, line int, column int) Token {
	return Token{
		Kind:   kind,
		Value:  value,
		Line:   line,
		Column: column,
	}
}
func (t *Token) isOneOfMany(expected ...TokenKind) bool {
	for _, tkn := range expected {
		if tkn == t.Kind {
			return true
		}
	}
	return false
}
func (t *Token) Debug() {
	if t.isOneOfMany(IDENTIFIER, NUMBER, STRING) {
		fmt.Printf("%s (%s)\n", TokenKindString(t.Kind), t.Value)
	} else {
		fmt.Printf("%s ()\n", TokenKindString(t.Kind))
	}
}
func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "--END OF FILE--"
	case NUMBER:
		return "number"
	case STRING:
		return "string"
	case IDENTIFIER:
		return "identifier"
	case OPEN_BRACKET:
		return "["
	case CLOSE_BRACKET:
		return "]"
	case OPEN_CURLY:
		return "{"
	case CLOSE_CURLY:
		return "}"
	case OPEN_PAREN:
		return "("
	case CLOSE_PAREN:
		return ")"
	case EQUALS:
		return "="
	case NOT:
		return "!"
	case NOT_EQUALS:
		return "!="

	case GREATER:
		return ">"
	case GREATER_EQUAL:
		return ">="
	case LESS:
		return "<"
	case LESS_EQUAL:
		return "<="

	case DOT:
		return "."
	case SEMICOLON:
		return ";"
	case COLON:
		return ":"
	case QUESTION:
		return "?"
	case COMMA:
		return ","
	case BACKTICK:
		return "´"

	case PLUS_PLUS:
		return "++"
	case MINUS_MINUS:
		return "--"
	case PLUS_EQUALS:
		return "+="
	case MINUS_EQUALS:
		return "-="
	case SLASH_EQUALS:
		return "/="
	case STAR_EQUALS:
		return "*="

	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case SLASH:
		return "/"
	case STAR:
		return "*"
	case PERCENT:
		return "%"

	case ALIAS:
		return "alias"
	case AND:
		return "and"
	case AUTOINSTANCIATE:
		return "autoinstanciate"
	case CALL:
		return "call"
	case CASE:
		return "case"
	case CATCH:
		return "catch"
	case CHOOSE:
		return "choose"
	case CLOSE:
		return "close"
	case COMMIT:
		return "commit"
	case CONNECT:
		return "connect"
	case CONSTANT:
		return "constant"
	case CONTINUE:
		return "continue"
	case CREATE:
		return "create"
	case CURSOR:
		return "cursor"
	case DECLARE:
		return "declare"
	case DELETE:
		return "delete"
	case DESCRIBE:
		return "describe"
	case DESCRIPTOR:
		return "descriptor"
	case DESTROY:
		return "destroy"
	case DISCONNECT:
		return "disconnec"
	case DO:
		return "do"
	case DYNAMIC:
		return "dynamic"
	case ELSE:
		return "else"
	case ELSEIF:
		return "elseif"
	case END:
		return "end"
	case ENUMERATED:
		return "enumerated"
	case EVENT:
		return "event"
	case EXECUTE:
		return "execute"
	case EXIT:
		return "exit"
	case EXTERNAL:
		return "external"
	case FALSE:
		return "false"
	case FETCH:
		return "fetch"
	case FINALLY:
		return "finally"
	case FIRST:
		return "first"
	case FOR:
		return "for"
	case FORWARD:
		return "forward"
	case FROM:
		return "from"
	case FUNCTION:
		return "function"
	case GLOBAL:
		return "global"
	case GOTO:
		return "goto"
	case HALT:
		return "halt"
	case IF:
		return "if"
	case IMMEDIATE:
		return "immediate"
	case INDIRECT:
		return "indirect"
	case INSERT:
		return "insert"
	case INTO:
		return "into"
	case INTRINSCI:
		return "intrinsci"
	case IS:
		return "is"
	case LAST:
		return "last"
	case LIBRARY:
		return "library"
	case LOOP:
		return "loop"
	case NAMESPACE:
		return "namespace"
	case NATIVE:
		return "native"
	case NEXT:
		return "next"
	case NOTOF:
		return "notof"
	case ON:
		return "on"
	case OPEN:
		return "open"
	case OR:
		return "or"
	case PARENT:
		return "parent"
	case POST:
		return "post"
	case PREPARE:
		return "prepare"
	case PRIOR:
		return "prior"
	case PRIVATE:
		return "private"
	case PRIVATEREAD:
		return "privateread"
	case PRIVATEWRITE:
		return "privatewrite"
	case PROCEDURE:
		return "procedure"
	case PROTECTED:
		return "protected"
	case PROTECTEDREAD:
		return "protectedread"
	case PROTECTEDWRITE:
		return "protectedwrite"
	case PROTOTYPES:
		return "prototypes"
	case PUBLIC:
		return "public"
	case READONLY:
		return "readonly"
	case REF:
		return "ref"
	case RETURN:
		return "return"
	case ROLLBACK:
		return "rollback"
	case RPCFUNC:
		return "rpcfunc"
	case SELECT:
		return "select"
	case SELECTBLOB:
		return "selectblob"
	case SHARED:
		return "shared"
	case STATIC:
		return "static"
	case STEP:
		return "step"
	case SUBROUTINE:
		return "subroutine"
	case SUPER:
		return "super"
	case SYSTEM:
		return "system"
	case SYSTEMREAD:
		return "systemread"
	case SYSTEMWRITE:
		return "systemwrite"
	case THEN:
		return "then"
	case THIS:
		return "this"
	case THROW:
		return "throw"
	case THROWS:
		return "throws"
	case TO:
		return "to"
	case TRIGGER:
		return "trigger"
	case TRUE:
		return "true"
	case TRY:
		return "try"
	case TYPE:
		return "type"
	case UNTIL:
		return "until"
	case UPDATE:
		return "update"
	case UPDATEBLOB:
		return "updateblob"
	case USING:
		return "using"
	case VARIABLES:
		return "variables"
	case WHILE:
		return "while"
	case WITH:
		return "with"
	case WITHIN:
		return "within"
	case XOR:
		return "xor"
	case _DEBUG:
		return "_debug"
	default:
		fmt.Printf("Unexpected Token %d", kind)
	}
	return ""
}
