package print

import (
	"compiler/grammar"
	"compiler/lexer"
	"fmt"
	"github.com/gookit/color"
	"strings"
)

func transfer(str string) string {
	str = strings.ReplaceAll(str, "G", "E'")
	str = strings.ReplaceAll(str, "S", "T'")
	return str
}

func PrintGrammar(gram []*grammar.Production) {
	start := "E"
	matched := ""
	for i := 0; i < len(gram); i++ {
		if gram[i].Type == "kill" {
			fmt.Printf("规约：%v\n", transfer(gram[i].Target))
			start = strings.Replace(start, gram[i].Target, "", 1)
			matched += gram[i].Target
			color.Green.Printf("%s", transfer(matched))
			fmt.Println(transfer(start))
			i++
		} else {
			fmt.Printf("推导：%v--->%v\n", transfer(gram[i].Origin), transfer(gram[i].Next))
			start = strings.Replace(start, gram[i].Origin, gram[i].Next, 1)
			color.Green.Printf("%s", transfer(matched))
			fmt.Println(transfer(start))
		}
	}
}

func PrintToken(tokens []*lexer.Token) {
	for i := 0; i < len(tokens); i++ {
		tokens[i].Show()
	}
}
