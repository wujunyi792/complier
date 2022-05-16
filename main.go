package main

import (
	"compiler/lexer"
	"fmt"
)

func main() {
	tokens := lexer.Analyse("for(int i=1.;i<=10;i++) begin\na_b++;#zszszszszs\nb_C--;#zszszszszs\nB123:=1234567;\na=@;\n123a=0;\na.b;\nend")
	for i := 0; i < len(tokens); i++ {
		fmt.Println(tokens[i].String())
	}
}
