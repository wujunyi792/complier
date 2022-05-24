package rule

import (
	"errors"
	"strings"
)

type Rule struct {
	Rules map[string][]string
}

type Formula struct {
	Left  string
	Right string
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

// 是否有空产生式
func (r *Rule) HaveEmptySet(first string) bool {
	for _, value := range r.Rules[first] {
		if value == "&" {
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

func (r *Rule) GetProcessMethod(first, end string) *Formula {
	for _, value := range r.Rules[first] {
		if value[0] == end[0] {
			return &Formula{
				Left:  first,
				Right: value,
			}
		}
	}
	return nil
}
