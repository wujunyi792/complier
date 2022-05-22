package lexer

import (
	"fmt"
	"github.com/gookit/color"
)

type TokenType int

var variableMap map[string]int

func init() {
	variableMap = make(map[string]int)
}

const (
	KEYWORD  TokenType = 1
	TYPE     TokenType = 2
	VARIABLE TokenType = 3
	OPERATOR TokenType = 4
	BRACKET  TokenType = 5
	STRING   TokenType = 6
	FLOAT    TokenType = 7
	BOOLEAN  TokenType = 8
	INTEGER  TokenType = 9
	COMMENT  TokenType = 10
	ERROR    TokenType = -1
)

func (tt TokenType) String() string {
	switch tt {

	case KEYWORD:
		return "keyword "
	case VARIABLE:
		return "variable"
	case TYPE:
		return "type    "
	case OPERATOR:
		return "operator"
	case BRACKET:
		return "bracket "
	case STRING:
		return "string  "
	case FLOAT:
		return "float   "
	case BOOLEAN:
		return "boolean "
	case INTEGER:
		return "integer "
	case COMMENT:
		return "comment "
	case ERROR:
		return "error"
	}

	panic("unexpected token type")
}

type Token struct {
	Typ    TokenType
	Value  string
	ID     int
	Column int
	Row    int
}

func NewToken(t TokenType, v string) *Token {
	ID := 0
	if t == 3 {
		if id, ok := variableMap[v]; !ok {
			variableMap[v] = len(variableMap) + 1
			ID = variableMap[v]
		} else {
			ID = id
		}
	}
	return &Token{Typ: t, Value: v, ID: ID}
}

func NewTokenWithLocation(t TokenType, v string, row int, column int) *Token {
	ID := 0
	if t == 3 {
		if id, ok := variableMap[v]; !ok {
			variableMap[v] = len(variableMap) + 1
			ID = variableMap[v]
		} else {
			ID = id
		}
	}
	return &Token{Typ: t, Value: v, ID: ID, Row: row, Column: column}
}

func (t *Token) IsVariable() bool {
	return t.Typ == VARIABLE
}

func (t *Token) IsScalar() bool {
	return t.Typ == FLOAT || t.Typ == BOOLEAN || t.Typ == INTEGER || t.Typ == STRING
}

func (t *Token) IsNumber() bool {
	return t.Typ == INTEGER || t.Typ == FLOAT
}

func (t *Token) IsOperator() bool {
	return t.Typ == OPERATOR
}

func (t *Token) IsBracket() bool {
	return t.Typ == BRACKET
}

func (t *Token) Show() {
	if t.Typ == ERROR {
		color.Red.Printf("【%d:%d】ERROR: You Have an Lexical error near \"%s\" at line %d column %d.\n", t.Row, t.Column, t.Value, t.Row, t.Column)
	} else if t.Typ == VARIABLE {
		color.Green.Printf("< %v, %v, id=%v >\n", t.Typ, t.Value, t.ID)
	} else {
		fmt.Printf("< %v, %v >\n", t.Typ, t.Value)
	}
}

func (t *Token) IsValue() bool {
	return t.IsVariable() || t.IsScalar()
}

func (t *Token) IsType() bool {
	return t.Typ == TYPE
}
