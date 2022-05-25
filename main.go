package main

import (
	"compiler/grammar"
	"compiler/grammarLL1"
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

func GrammarLL1(tokens []*lexer.Token) {
	gram, correct := grammarLL1.Analyze(tokens, "E->TG\nG->ATG|&\nT->FS\nS->MFS|&\nF->(E)|i\nA->+|-\nM->*|/", "E")
	if !correct {
		panic("语法推导失败")
	}
	servicePrint.PrintGrammarLL1(gram)
}

func main() {
	//tokens := MakeToken("i*(i-i)/(i+i)*((i+i-i*i/i))")
	//tokens := MakeToken("i+i")
	tokens := MakeToken("i*i**")
	//tokens := MakeToken("i+i*i(")
	//tokens := MakeToken("i+i*i/i-i)")
	//tokens := MakeToken("i+)i-i(")
	//tokens := MakeToken("(i-i)(i/i)")
	//Grammar(tokens)
	GrammarLL1(tokens)
}
