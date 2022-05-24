package setFirst

import (
	"compiler/grammarLL1/rule"
	"compiler/grammarLL1/util"
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
					if MergeSet(firstSet[key], map[string]struct{}{string(v[0]): {}}) != 0 {
						changed = true
					}
					continue
				} else {
					// 第一个是非终结符
					if RemoveEmptyAndMergeSet(firstSet[key], firstSet[string(v[0])]) != 0 {
						changed = true
					}
				}

				// 查看后面是不是终结符
				//if rule.AllIsTer(v) {
				//	var j = 0
				//	for j < len(v) {
				//		// 含有empty表达式把其first合到当前first集
				//		if rules.HaveEmptyFormula(string(v[j])) {
				//			if RemoveEmptyAndMergeSet(firstSet[key], firstSet[string(v[j])]) != 0 {
				//				changed = true
				//			}
				//			j++
				//		} else {
				//			break
				//		}
				//	}
				//	// 全部都含有empty表达式
				//	if j == len(v) {
				//		if MergeSet(firstSet[key], map[string]struct{}{"@": {}}) != 0 {
				//			changed = true
				//		}
				//	}
				//}
			}
		}
		if !changed {
			break
		}
	}
	return firstSet
}

func RemoveEmptyAndMergeSet(a map[string]struct{}, b map[string]struct{}) int {
	flag := false
	if _, flag = b["@"]; flag {
		delete(b, "@")
	}
	count := 0
	for key, value := range b {
		if _, ok := a[key]; !ok {
			count++
		}
		a[key] = value
	}
	if flag {
		b["@"] = struct{}{}
	}
	return count
}

func MergeSet(a map[string]struct{}, b map[string]struct{}) int {
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
		build.WriteString(fmt.Sprintf("FIRST(%s) = { ", key))
		for item := range value {
			build.WriteString(fmt.Sprintf("%s ", item))
		}
		build.WriteString("}\n")
	}
	return build.String()
}

func (f FirstSet) Strings() []string {
	var build strings.Builder
	var ans []string
	for key, value := range f {
		build.WriteString(fmt.Sprintf("FIRST(%s)={ ", key))
		for item := range value {
			build.WriteString(fmt.Sprintf("%s ", item))
		}
		build.WriteString("}")
		ans = append(ans, build.String())
		build.Reset()
	}
	return ans
}

func (f FirstSet) HaveEmpty(first string) bool {
	_, ok := f[first]["@"]
	return ok
}
