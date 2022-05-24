package grammarLL1

import (
	"compiler/grammarLL1/analysisTable"
	"compiler/grammarLL1/first"
	"compiler/grammarLL1/follow"
	"compiler/grammarLL1/rule"
	"fmt"
	"testing"
)

func TestNewRule(t *testing.T) {
	g := rule.NewRules()
	_ = g.AddRules("E->TS\nS->+TS|&\nT->FG\nG->*FG|&\nF->(E)|i")
	//fmt.Println(t)
	firstSet := first.GetFirstSet(g)
	fmt.Println(firstSet.String())
	followSet := follow.GetFollowSet(g, "E", firstSet)
	fmt.Println(followSet.String())
	table := analysisTable.GetAnalyzeTable(firstSet, followSet, g)
	fmt.Println(table)
}
