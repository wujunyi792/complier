package print

import (
	"compiler/grammar"
	"compiler/grammarLL1"
	"compiler/lexer"
	"compiler/util/transfer"
	"fmt"
	"github.com/gookit/color"
	"strings"
)

func PrintGrammar(gram []*grammar.Production) {
	start := "E"
	matched := ""
	for i := 0; i < len(gram); i++ {
		if gram[i].Type == "kill" {
			fmt.Printf("规约：%v\n", transfer.Transfer(gram[i].Target))
			start = strings.Replace(start, gram[i].Target, "", 1)
			matched += gram[i].Target
			color.Green.Printf("%s", transfer.Transfer(matched))
			color.Red.Printf(transfer.Transfer(start + "\n"))
			i++
		} else {
			fmt.Printf("推导：%v--->%v\n", transfer.Transfer(gram[i].Origin), transfer.Transfer(gram[i].Next))
			start = strings.Replace(start, gram[i].Origin, gram[i].Next, 1)
			color.Green.Printf("%s", transfer.Transfer(matched))
			color.Red.Printf(transfer.Transfer(start + "\n"))
		}
	}
}

func PrintGrammarLL1(gram []*grammarLL1.Production) {
	start := "E"
	matched := ""
	for i := 0; i < len(gram)-1; i++ {
		if gram[i].Type == "kill" {
			fmt.Printf("规约：%v\n", transfer.Transfer(gram[i].Target))
			start = strings.Replace(start, gram[i].Target, "", 1)
			matched += gram[i].Target
			color.Green.Printf("%s", transfer.Transfer(matched))
			color.Red.Printf(transfer.Transfer(start + "\n"))
		} else {
			fmt.Printf("推导：%v--->%v\n", transfer.Transfer(gram[i].Origin), transfer.Transfer(gram[i].Next))
			start = strings.Replace(start, gram[i].Origin, gram[i].Next, 1)
			color.Green.Printf("%s", transfer.Transfer(matched))
			color.Red.Printf(transfer.Transfer(start + "\n"))
		}
	}
}

func PrintToken(tokens []*lexer.Token) bool {
	err := false
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Typ == lexer.ERROR {
			err = true
		}
		tokens[i].Show()
	}
	fmt.Println("\n ")
	return err
}
