package main

import (
	"compiler/grammar"
	"compiler/lexer"
	"fmt"
	"github.com/gookit/color"
	"strings"
)

func MakeToken(code string) []*lexer.Token {
	//tokens := lexer.Analyse("while(a==b)begin\na:=a+1;#zhushi\nb:=b-1;\nc=c*d;\nd=c/d;\nif(a>b) then  c=C else C=c;\nend")
	//tokens := lexer.Analyse("for(int i=1;i<=10;i++) begin\na_b++;#zszszszszs\nb_C--;#zszszszszs\nB123:=1234567;\na=@;\n123a=0;\na.b;\nend")
	tokens := lexer.Analyse(code)
	//tokens := lexer.Analyse("i*i**")
	for i := 0; i < len(tokens); i++ {
		tokens[i].Show()
	}
	return tokens
}

func Grammar(tokens []*lexer.Token) {
	gram, correct := grammar.Analyse(tokens)
	if !correct {
		return
	}
	start := "E"
	matched := ""
	for _, g := range gram {
		if g.Type == "kill" {
			fmt.Printf("规约：%v\n", g.Target)
			start = strings.Replace(start, g.Target, "", 1)
			matched += g.Target
			color.Green.Printf("%s", matched)
			fmt.Println(start)
		} else {
			fmt.Printf("推导：%v--->%v\n", g.Origin, g.Next)
			start = strings.Replace(start, g.Origin, g.Next, 1)
			color.Green.Printf("%s", matched)
			fmt.Println(start)
		}
	}
}

func main() {
	tokens := MakeToken("i*i")
	Grammar(tokens)
}
