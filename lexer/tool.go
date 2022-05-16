package lexer

import "regexp"

var (
	ptnLetter   = regexp.MustCompile("^[_a-zA-Z]$")
	ptnNumber   = regexp.MustCompile("^[0-9]$")
	ptnLiteral  = regexp.MustCompile("^[_a-zA-Z0-9]$")
	ptnOperator = regexp.MustCompile("^[+\\-*<>=!&|^%/;:]$")
	ptnBracket  = regexp.MustCompile("^[()]$")
)

func IsLetter(c string) bool {
	return ptnLetter.MatchString(c)
}

func IsNumber(c string) bool {
	return ptnNumber.MatchString(c)
}

func IsLiteral(c string) bool {
	return ptnLiteral.MatchString(c)
}
func IsBracket(c string) bool {
	return ptnBracket.MatchString(c)
}

func IsOperator(c string) bool {
	return ptnOperator.MatchString(c)
}

var KeyWords = map[string]bool{
	"begin": true,
	"end":   true,
	"if":    true,
	"then":  true,
	"else":  true,
	"for":   true,
	"while": true,
	"do":    true,
	"and":   true,
	"or":    true,
	"not":   true,
}

func IsKeyword(key string) bool {
	return KeyWords[key]
}

var KeyTypes = map[string]bool{
	"int":    true,
	"float":  true,
	"double": true,
	"byte":   true,
	"string": true,
}

func IsKeyTypes(key string) bool {
	return KeyTypes[key]
}
