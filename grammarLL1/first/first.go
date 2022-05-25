package first

import (
	"compiler/grammarLL1/rule"
	"compiler/grammarLL1/util"
	"compiler/util/transfer"
	"fmt"
	"strings"
)

type FirstSet map[string]map[string]struct{}

func GetFirstSet(rules *rule.Rule) FirstSet {
	firstSet := make(FirstSet)
	var changed bool
	for {
		changed = false
		for key, r := range rules.Rules {
			// key 左值 r推导值
			if firstSet[key] == nil {
				firstSet[key] = make(map[string]struct{})
			}
			for _, v := range r {
				// 遍历产生式
				// 第一个是终结符,直接将终结符加进first集
				if util.IsTerminal(v[0]) {
					if mergeSet(firstSet[key], map[string]struct{}{string(v[0]): {}}) != 0 {
						changed = true
					}
					continue
				} else {
					// 第一个是非终结符
					if removeEmptyAndMergeSet(firstSet[key], firstSet[string(v[0])]) != 0 {
						changed = true
					}
				}
			}
		}
		if !changed {
			break
		}
	}
	return firstSet
}

func removeEmptyAndMergeSet(a map[string]struct{}, b map[string]struct{}) int {
	flag := false
	if _, flag = b["&"]; flag {
		delete(b, "&")
	}
	count := 0
	for key, value := range b {
		if _, ok := a[key]; !ok {
			count++
		}
		a[key] = value
	}
	if flag {
		b["&"] = struct{}{}
	}
	return count
}

func mergeSet(a map[string]struct{}, b map[string]struct{}) int {
	count := 0
	for key, value := range b {
		if _, ok := a[key]; !ok {
			count++
		}
		a[key] = value
	}
	return count
}

func (f FirstSet) String() string {
	var build strings.Builder
	for key, value := range f {
		build.WriteString(fmt.Sprintf("FIRST(%s) = { ", transfer.Transfer(key)))
		for item := range value {
			build.WriteString(fmt.Sprintf("%s ", transfer.Transfer(item)))
		}
		build.WriteString("}\n")
	}
	return build.String()
}

func (f FirstSet) haveEmpty(first string) bool {
	_, ok := f[first]["&"]
	return ok
}

func (f FirstSet) IsInFirstSet(first string, target string) bool {
	for key := range f[first] {
		if key == target {
			return true
		}
	}
	return false
}
