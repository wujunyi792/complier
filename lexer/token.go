package lexer

import "fmt"

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
		return "error "
	}

	panic("unexpected token type")
}

type Token struct {
	Typ   TokenType
	Value string
	ID    int
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

func (t *Token) String() string {
	return fmt.Sprintf("< %v, %v, %v >", t.Typ, t.Value, t.ID)
}

func (t *Token) IsValue() bool {
	return t.IsVariable() || t.IsScalar()
}

func (t *Token) IsType() bool {
	return t.Typ == TYPE
}
