package main

import "compiler/lexer"

func MakeToken() {
	tokens := lexer.Analyse("for(int i=1;i<=10;i++) begin\na_b++;#zszszszszs\nb_C--;#zszszszszs\nB123:=1234567;\na=@;\n123a=0;\na.b;\nend")
	for i := 0; i < len(tokens); i++ {
		tokens[i].Show()
	}
}

func main() {

}
