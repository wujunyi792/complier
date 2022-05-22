package main

import (
	"compiler/grammar"
	"compiler/lexer"
	servicePrint "compiler/print"
)

func MakeToken(code string) []*lexer.Token {
	tokens := lexer.Analyse(code)
	servicePrint.PrintToken(tokens)
	return tokens
}

func Grammar(tokens []*lexer.Token) {
	gram, correct := grammar.Analyse(tokens)
	if !correct {
		panic("语法推导失败")
	}
	servicePrint.PrintGrammar(gram)
}

func main() {
	//tokens := MakeToken("i*(i-i)/(i+i)*((i+i-i*i/i))")
	//tokens := MakeToken("i*i**")
	//tokens := MakeToken("i+i*i(")
	//tokens := MakeToken("i+i*i/i-i)")
	//tokens := MakeToken("i+)i-i(")
	tokens := MakeToken("(i-i)(i/i)")
	Grammar(tokens)
}
