package rule

import (
	"compiler/grammarLL1/util"
	"errors"
	"strings"
)

type Rule struct {
	Rules map[string][]string
}

func NewRules() *Rule {
	return &Rule{Rules: make(map[string][]string)}
}

func (r *Rule) AddRules(s string) error {
	lineRule := strings.Split(s, "\n")
	for _, t := range lineRule {
		if strings.TrimSpace(t) == "" {
			continue
		}
		c := strings.Split(t, "->")
		if len(c) != 2 || len(c[0]) != 1 {
			return errors.New("invalid arg")
		}
		right := strings.Split(strings.ReplaceAll(c[1], " ", ""), "|")
		for i := range right {
			r.Rules[c[0]] = append(r.Rules[c[0]], right[i])
		}
	}
	return nil
}

//func (r Rules) String() string {
//	var builder strings.Builder
//	for fist, second := range r {
//		for i := range second {
//			builder.WriteString(fmt.Sprintf("%c->%s\n", fist, second[i]))
//		}
//	}
//	return builder.String()
//}
//
func AllIsTer(s string) bool {
	for _, k := range s {
		if util.IsTerminal(byte(k)) {
			return false
		}
	}
	return true
}

//
// 判断是否存在有 X->@
func (r *Rule) HaveEmptyFormula(first string) bool {
	for _, value := range r.Rules[first] {
		if value == "@" {
			return true
		}
	}
	return false
}

func (r *Rule) TheFirstItemIs(first, item string) string {
	for _, value := range r.Rules[first] {
		if value == item {
			return value
		}
	}
	return ""
}

//
//func (r Rules) Dfs(first, terminal byte) bool {
//	for i := range r[first] {
//		if util.IsTerminal(r[first][i][0]) && r[first][i][0] == terminal {
//			return true
//		} else if !util.IsTerminal(r[first][i][0]) {
//			if ok := r.Dfs(r[first][i][0], terminal); ok {
//				return true
//			}
//		}
//	}
//	return false
//}
