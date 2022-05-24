package grammarLL1

import (
	"compiler/grammarLL1/SetFollow"
	"compiler/grammarLL1/rule"
	"compiler/grammarLL1/setFirst"
	"fmt"
	"testing"
)

func TestNewRule(t *testing.T) {
	g := rule.NewRules()
	_ = g.AddRules("E->TS\nS->+TS|&\nT->FG\nG->*FG|&\nF->(E)|i")
	//fmt.Println(t)
	firstSet := setFirst.GetFirstSet(g)
	fmt.Println(firstSet.String())
	followSet := SetFollow.GetFollowSet(g, "E", firstSet)
	fmt.Println(followSet.String())
}
