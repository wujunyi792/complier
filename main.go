package main

import (
	"compiler/grammar"
	"compiler/lexer"
	servicePrint "compiler/print"
)

func MakeToken(code string) []*lexer.Token {
	//tokens := lexer.Analyse("while(a==b)begin\na:=a+1;#zhushi\nb:=b-1;\nc=c*d;\nd=c/d;\nif(a>b) then  c=C else C=c;\nend")
	//tokens := lexer.Analyse("for(int i=1;i<=10;i++) begin\na_b++;#zszszszszs\nb_C--;#zszszszszs\nB123:=1234567;\na=@;\n123a=0;\na.b;\nend")
	tokens := lexer.Analyse(code)
	//tokens := lexer.Analyse("i*i**")
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
	tokens := MakeToken("i*i")
	Grammar(tokens)
}
