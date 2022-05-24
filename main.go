package main

import (
	"compiler/grammar"
	"compiler/lexer"
	servicePrint "compiler/print"
	"github.com/gookit/color"
	"os"
)

func MakeToken(code string) []*lexer.Token {
	tokens := lexer.Analyse(code)
	err := servicePrint.PrintToken(tokens)
	if err {
		color.Redln("Lexer ERR, please check")
		os.Exit(-1)
	}
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
	tokens := MakeToken("i*(i-i)/(i+i)*((i+i-i*i/i))")
	//tokens := MakeToken("i*i**")
	//tokens := MakeToken("i+i*i(")
	//tokens := MakeToken("i+i*i/i-i)")
	//tokens := MakeToken("i+)i-i(")
	//tokens := MakeToken("(i-i)(i/i)")
	Grammar(tokens)
}
