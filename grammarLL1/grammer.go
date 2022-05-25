package grammarLL1

import (
	"bytes"
	"compiler/grammarLL1/analysisTable"
	"compiler/grammarLL1/first"
	"compiler/grammarLL1/follow"
	"compiler/grammarLL1/rule"
	util2 "compiler/grammarLL1/util"
	"compiler/lexer"
	"compiler/util"
	"fmt"
	"github.com/gookit/color"
	"io"
)

const EndToken = "#"

type Grammar struct {
	Stream   *util.Stream
	table    analysisTable.SymbolTable
	endToken string
	tokens   *util.Queue
}

func Analyze(raw []*lexer.Token, rules string, start string) ([]*Production, bool) {
	g := rule.NewRules()
	_ = g.AddRules(rules)
	firstSet := first.GetFirstSet(g)
	fmt.Println(firstSet.String())
	followSet := follow.GetFollowSet(g, start, firstSet)
	fmt.Println(followSet.String())
	table := analysisTable.GetAnalyzeTable(firstSet, followSet, g)
	res := table.String()
	fmt.Println(res)

	grm := NewGrammar(raw, bytes.NewBufferString(start), EndToken, table)
	prod := grm.Analyze()
	return prod, true
}

func NewGrammar(token []*lexer.Token, r io.Reader, et string, rule analysisTable.SymbolTable) *Grammar {
	s := util.NewStream(r, EndToken)
	q := util.New()
	for i := 0; i < len(token); i++ {
		q.PushBack(token[i])
	}
	q.PushBack(&lexer.Token{
		Typ:   lexer.END,
		Value: "#",
	})
	return &Grammar{Stream: s, endToken: et, table: rule, tokens: q}
}

func (g *Grammar) Analyze() (res []*Production) {
	for {
		if g.tokens.Front().(*lexer.Token).Typ == lexer.END && g.Stream.Peek() == "#" {
			res = append(res, &Production{
				Type:   "kill",
				Target: "#",
			})
			break
		}
		if g.tokens.Front().(*lexer.Token).Typ != lexer.END && g.Stream.Peek() == "#" {
			color.Redln("Wrong grammar")
			break
		}

		ProcessC := g.Stream.Next()
		TargetC := g.tokens.Front().(*lexer.Token)
		if ProcessC == "&" {
			res = append(res, &Production{
				Type:   "kill",
				Target: ProcessC,
			})
			continue
		}
		if util2.IsTerminal(ProcessC[0]) {
			if ProcessC == TargetC.Value || (ProcessC == "i" && TargetC.Typ == lexer.VARIABLE) {
				res = append(res, &Production{
					Type:   "kill",
					Target: TargetC.Value,
				})
				g.tokens.Pop()
				continue
			} else {
				color.Redln("Wrong grammar")
				return
			}
		} else {
			proc := g.table[ProcessC][TargetC.Value]
			if proc == nil {
				color.Redln("Wrong grammar")
				return
			}
			for i := 0; i < len(proc.Right); i++ {
				g.Stream.PutBack(string(proc.Right[len(proc.Right)-1-i]))
			}
			res = append(res, &Production{
				Type:   "Continue",
				Origin: proc.Left,
				Next:   proc.Right,
			})
		}
	}
	return
}
