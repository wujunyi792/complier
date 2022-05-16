package main

import (
	"compiler/lexer"
)

func main() {
	tokens := lexer.Analyse("for(int i=1..1;i<=10;i++) begin\n\na_b++;#zszszszszs\nb_C--;#zszszszszs\nB123:=1234567;\na=@;\n123ab_=0;\na.b;\nend")
	for i := 0; i < len(tokens); i++ {
		tokens[i].Show()
	}
}
