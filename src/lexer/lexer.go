package lexer

import (
	"fmt"
	"regexp"
)

type regexHandler func(lex *Lexer, regex *regexp.Regexp)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type Lexer struct {
	patterns []regexPattern
	Tokens   []Token
	source   string
	current  int
	line     int
	column   int
}

func (l *Lexer) advanceN(n int) {
	l.current += n
	l.column += n
	if l.current < len(l.source) {
		curRune := rune(l.source[l.current])
		if curRune == '\n' {
			l.line++
			l.column = 1
		}
	}
}
func (l *Lexer) push(t Token) {
	l.Tokens = append(l.Tokens, t)
}
func (l *Lexer) at() string {
	return string(l.source[l.current])
}
func (l *Lexer) remainder() string {
	return l.source[l.current:]
}
func (l *Lexer) atEOF() bool {
	return l.current >= len(l.source)
}
func Tokenize(source []byte) []Token {
	lex := NewLexer(source)

	for !lex.atEOF() {
		matched := false

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(string(lex.remainder()))

			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		if !matched {
			panic(fmt.Sprintf("Lexer::Error -> unrecoginized token at line %d column %d near %s\n", lex.line, lex.column, string(lex.remainder())))
		}
	}

	lex.push(NewToken(EOF, "EOF", lex.line, lex.column))

	return lex.Tokens
}

func numberHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindString(string(lex.remainder()))
	lex.push(NewToken(NUMBER, match, lex.line, lex.column))
	lex.advanceN(len(match))
}
func skipHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(string(lex.remainder()))
	lex.advanceN(match[1])
}
func stringHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	stringLiteral := lex.remainder()[match[0]+1:match[1]-1]

	lex.push(NewToken(STRING, stringLiteral, lex.line, lex.column))
	lex.advanceN(len(stringLiteral)+2)
}
func symbolHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())

	if kind, exists := reserved_lu[match]; exists {
		lex.push(NewToken(kind, match, lex.line, lex.column))
	} else {
		lex.push(NewToken(IDENTIFIER, match, lex.line, lex.column))
	}

	lex.advanceN(len(match))
}
func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *Lexer, regex *regexp.Regexp) {
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value, lex.line, lex.column))
	}
}
func NewLexer(source []byte) *Lexer {
	return &Lexer{
		source: string(source),
		line:   1,
		patterns: []regexPattern{
			{regexp.MustCompile(`[a-zA-Z_][a-zA-Z_0-9]*`), symbolHandler},
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			{regexp.MustCompile(`"[^"]*"`), stringHandler},
			{regexp.MustCompile(`'[^"]*'`), stringHandler},
			{regexp.MustCompile(`\/\/.*`), skipHandler},
			{regexp.MustCompile(`\/\*[^*\/]*\/`), skipHandler},
			{regexp.MustCompile(`\s+`), skipHandler},

			{regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			{regexp.MustCompile(`\]`), defaultHandler(CLOSE_BRACKET, "]")},
			{regexp.MustCompile(`\{`), defaultHandler(OPEN_CURLY, "{")},
			{regexp.MustCompile(`\}`), defaultHandler(CLOSE_CURLY, "}")},
			{regexp.MustCompile(`\(`), defaultHandler(OPEN_PAREN, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(CLOSE_PAREN, ")")},
			{regexp.MustCompile(`=`), defaultHandler(EQUALS, "=")},
			{regexp.MustCompile(`!=`), defaultHandler(NOT_EQUALS, "!=")},
			{regexp.MustCompile(`!`), defaultHandler(NOT, "!")},
			{regexp.MustCompile(`>=`), defaultHandler(GREATER_EQUAL, ">=")},
			{regexp.MustCompile(`>`), defaultHandler(GREATER, ">")},
			{regexp.MustCompile(`<=`), defaultHandler(LESS_EQUAL, "<=")},
			{regexp.MustCompile(`<`), defaultHandler(LESS, "<")},
			{regexp.MustCompile(`\.`), defaultHandler(DOT, ".")},
			{regexp.MustCompile(`;`), defaultHandler(SEMICOLON, ";")},
			{regexp.MustCompile(`:`), defaultHandler(COLON, ":")},
			{regexp.MustCompile(`\?`), defaultHandler(QUESTION, "?")},
			{regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
			{regexp.MustCompile("`"), defaultHandler(BACKTICK, "`")},
			{regexp.MustCompile(`\+\+`), defaultHandler(PLUS_PLUS, "++")},
			{regexp.MustCompile(`--`), defaultHandler(MINUS_MINUS, "--")},
			{regexp.MustCompile(`\+=`), defaultHandler(PLUS_EQUALS, "+=")},
			{regexp.MustCompile(`-=`), defaultHandler(MINUS_EQUALS, "-=")},
			{regexp.MustCompile(`/=`), defaultHandler(SLASH_EQUALS, "/=")},
			{regexp.MustCompile(`\*=`), defaultHandler(STAR_EQUALS, "*=")},
			{regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`-`), defaultHandler(MINUS, "-")},
			{regexp.MustCompile(`/`), defaultHandler(SLASH, "/")},
			{regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
			{regexp.MustCompile(`%`), defaultHandler(PERCENT, "%")},
			/*
				{regexp.MustCompile(`alias`), defaultHandler(ALIAS, "alias")},
				{regexp.MustCompile(`and`), defaultHandler(AND, "and")},
				{regexp.MustCompile(`autoinstanciate`), defaultHandler(AUTOINSTANCIATE, "autoinstanciate")},
				{regexp.MustCompile(`call`), defaultHandler(CALL, "call")},
				{regexp.MustCompile(`case`), defaultHandler(CASE, "case")},
				{regexp.MustCompile(`catch`), defaultHandler(CATCH, "catch")},
				{regexp.MustCompile(`choose`), defaultHandler(CHOOSE, "choose")},
				{regexp.MustCompile(`close`), defaultHandler(CLOSE, "close")},
				{regexp.MustCompile(`commit`), defaultHandler(COMMIT, "commit")},
				{regexp.MustCompile(`connect`), defaultHandler(CONNECT, "connect")},
				{regexp.MustCompile(`constant`), defaultHandler(CONSTANT, "constant")},
				{regexp.MustCompile(`continue`), defaultHandler(CONTINUE, "continue")},
				{regexp.MustCompile(`create`), defaultHandler(CREATE, "create")},
				{regexp.MustCompile(`cursor`), defaultHandler(CURSOR, "cursor")},
				{regexp.MustCompile(`declare`), defaultHandler(DECLARE, "declare")},
				{regexp.MustCompile(`delete`), defaultHandler(DELETE, "delete")},
				{regexp.MustCompile(`describe`), defaultHandler(DESCRIBE, "describe")},
				{regexp.MustCompile(`descriptor`), defaultHandler(DESCRIPTOR, "descriptor")},
				{regexp.MustCompile(`destroy`), defaultHandler(DESTROY, "destroy")},
				{regexp.MustCompile(`disconnect`), defaultHandler(DISCONNECT, "disconnect")},
				{regexp.MustCompile(`do`), defaultHandler(DO, "do")},
				{regexp.MustCompile(`dynamic`), defaultHandler(DYNAMIC, "dynamic")},
				{regexp.MustCompile(`else`), defaultHandler(ELSE, "else")},
				{regexp.MustCompile(`elseif`), defaultHandler(ELSEIF, "elseif")},
				{regexp.MustCompile(`end`), defaultHandler(END, "end")},
				{regexp.MustCompile(`enumerated`), defaultHandler(ENUMERATED, "enumerated")},
				{regexp.MustCompile(`event`), defaultHandler(EVENT, "event")},
				{regexp.MustCompile(`execute`), defaultHandler(EXECUTE, "execute")},
				{regexp.MustCompile(`exit`), defaultHandler(EXIT, "exit")},
				{regexp.MustCompile(`external`), defaultHandler(EXTERNAL, "external")},
				{regexp.MustCompile(`false`), defaultHandler(FALSE, "false")},
				{regexp.MustCompile(`fetch`), defaultHandler(FETCH, "fetch")},
				{regexp.MustCompile(`finally`), defaultHandler(FINALLY, "finally")},
				{regexp.MustCompile(`first`), defaultHandler(FIRST, "first")},
				{regexp.MustCompile(`for`), defaultHandler(FOR, "for")},
				{regexp.MustCompile(`forward`), defaultHandler(FORWARD, "forward")},
				{regexp.MustCompile(`from`), defaultHandler(FROM, "from")},
				{regexp.MustCompile(`function`), defaultHandler(FUNCTION, "function")},
				{regexp.MustCompile(`global`), defaultHandler(GLOBAL, "global")},
				{regexp.MustCompile(`goto`), defaultHandler(GOTO, "goto")},
				{regexp.MustCompile(`halt`), defaultHandler(HALT, "halt")},
				{regexp.MustCompile(`if`), defaultHandler(IF, "if")},
				{regexp.MustCompile(`immediate`), defaultHandler(IMMEDIATE, "immediate")},
				{regexp.MustCompile(`indirect`), defaultHandler(INDIRECT, "indirect")},
				{regexp.MustCompile(`insert`), defaultHandler(INSERT, "insert")},
				{regexp.MustCompile(`into`), defaultHandler(INTO, "into")},
				{regexp.MustCompile(`intrinsci`), defaultHandler(INTRINSCI, "intrinsci")},
				{regexp.MustCompile(`is`), defaultHandler(IS, "is")},
				{regexp.MustCompile(`last`), defaultHandler(LAST, "last")},
				{regexp.MustCompile(`library`), defaultHandler(LIBRARY, "library")},
				{regexp.MustCompile(`loop`), defaultHandler(LOOP, "loop")},
				{regexp.MustCompile(`namespace`), defaultHandler(NAMESPACE, "namespace")},
				{regexp.MustCompile(`native`), defaultHandler(NATIVE, "native")},
				{regexp.MustCompile(`next`), defaultHandler(NEXT, "next")},
				{regexp.MustCompile(`notoff`), defaultHandler(NOTOF, "notoff")},
				{regexp.MustCompile(`on`), defaultHandler(ON, "on")},
				{regexp.MustCompile(`open`), defaultHandler(OPEN, "open")},
				{regexp.MustCompile(`or`), defaultHandler(OR, "or")},
				{regexp.MustCompile(`parent`), defaultHandler(PARENT, "parent")},
				{regexp.MustCompile(`post`), defaultHandler(POST, "post")},
				{regexp.MustCompile(`prepare`), defaultHandler(PREPARE, "prepare")},
				{regexp.MustCompile(`prior`), defaultHandler(PRIOR, "prior")},
				{regexp.MustCompile(`private`), defaultHandler(PRIVATE, "private")},
				{regexp.MustCompile(`privateread`), defaultHandler(PRIVATEREAD, "privateread")},
				{regexp.MustCompile(`privatewrite`), defaultHandler(PRIVATEWRITE, "privatewrite")},
				{regexp.MustCompile(`procedure`), defaultHandler(PROCEDURE, "procedure")},
				{regexp.MustCompile(`protected`), defaultHandler(PROTECTED, "protected")},
				{regexp.MustCompile(`protectedread`), defaultHandler(PROTECTEDREAD, "protectedread")},
				{regexp.MustCompile(`protectedwrite`), defaultHandler(PROTECTEDWRITE, "protectedwrite")},
				{regexp.MustCompile(`prototypes`), defaultHandler(PROTOTYPES, "prototypes")},
				{regexp.MustCompile(`public`), defaultHandler(PUBLIC, "public")},
				{regexp.MustCompile(`readonly`), defaultHandler(READONLY, "readonly")},
				{regexp.MustCompile(`ref`), defaultHandler(REF, "ref")},
				{regexp.MustCompile(`return`), defaultHandler(RETURN, "return")},
				{regexp.MustCompile(`rollback`), defaultHandler(ROLLBACK, "rollback")},
				{regexp.MustCompile(`rpcfunc`), defaultHandler(RPCFUNC, "rpcfunc")},
				{regexp.MustCompile(`select`), defaultHandler(SELECT, "select")},
				{regexp.MustCompile(`selectblob`), defaultHandler(SELECTBLOB, "selectblob")},
				{regexp.MustCompile(`shared`), defaultHandler(SHARED, "shared")},
				{regexp.MustCompile(`static`), defaultHandler(STATIC, "static")},
				{regexp.MustCompile(`step`), defaultHandler(STEP, "step")},
				{regexp.MustCompile(`subroutine`), defaultHandler(SUBROUTINE, "subroutine")},
				{regexp.MustCompile(`super`), defaultHandler(SUPER, "super")},
				{regexp.MustCompile(`system`), defaultHandler(SYSTEM, "system")},
				{regexp.MustCompile(`systemread`), defaultHandler(SYSTEMREAD, "systemread")},
				{regexp.MustCompile(`systemwrite`), defaultHandler(SYSTEMWRITE, "systemwrite")},
				{regexp.MustCompile(`then`), defaultHandler(THEN, "then")},
				{regexp.MustCompile(`this`), defaultHandler(THIS, "this")},
				{regexp.MustCompile(`throw`), defaultHandler(THROW, "throw")},
				{regexp.MustCompile(`throws`), defaultHandler(THROWS, "throws")},
				{regexp.MustCompile(`to`), defaultHandler(TO, "to")},
				{regexp.MustCompile(`trigger`), defaultHandler(TRIGGER, "trigger")},
				{regexp.MustCompile(`true`), defaultHandler(TRUE, "true")},
				{regexp.MustCompile(`try`), defaultHandler(TRY, "try")},
				{regexp.MustCompile(`type`), defaultHandler(TYPE, "type")},
				{regexp.MustCompile(`until`), defaultHandler(UNTIL, "until")},
				{regexp.MustCompile(`update`), defaultHandler(UPDATE, "update")},
				{regexp.MustCompile(`updateblob`), defaultHandler(UPDATEBLOB, "updateblob")},
				{regexp.MustCompile(`using`), defaultHandler(USING, "using")},
				{regexp.MustCompile(`variables`), defaultHandler(VARIABLES, "variables")},
				{regexp.MustCompile(`while`), defaultHandler(WHILE, "while")},
				{regexp.MustCompile(`with`), defaultHandler(WITH, "with")},
				{regexp.MustCompile(`within`), defaultHandler(WITHIN, "within")},
				{regexp.MustCompile(`xor`), defaultHandler(XOR, "xor")},
				{regexp.MustCompile(`_debug`), defaultHandler(_DEBUG, "_debug")},
			*/
		},
	}
}
