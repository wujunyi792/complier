package analysisTable

import (
	"compiler/grammarLL1/first"
	"compiler/grammarLL1/follow"
	"compiler/grammarLL1/rule"
	"compiler/util/transfer"
	"fmt"
	"github.com/liushuochen/gotable"
)

type SymbolTable map[string]map[string]*rule.Formula

func GetAnalyzeTable(firstSet first.FirstSet, followSet follow.FollowSet, rules *rule.Rule) SymbolTable {
	symbolTable := make(SymbolTable)
	endSymbol := make(map[string]struct{})
	for key, firstSets := range firstSet {
		if _, ok := symbolTable[key]; !ok {
			symbolTable[key] = make(map[string]*rule.Formula)
		}
		for set := range firstSets {
			if set == "&" {
				continue
			}
			endSymbol[set] = struct{}{}
		}
	}
	for _, followSets := range followSet {
		for set := range followSets {
			endSymbol[set] = struct{}{}
		}
	}

	for left := range firstSet {
		for key, _ := range endSymbol {
			symbolTable[left][key] = nil
		}
	}
	// 到此表结构组装完成

	changed := false
	for {
		changed = false
		for left, sets := range firstSet {
			for set, _ := range sets {
				if symbolTable[left][set] != nil {
					continue
				}

				// first集有空集
				if set == "&" {
					for fl := range followSet[left] {
						if symbolTable[left][fl] == nil {
							symbolTable[left][fl] = &rule.Formula{
								Left:  left,
								Right: set,
							}
							changed = true

						}
					}
					continue
				}
				// 无空集
				symbolTable[left][set] = rules.GetProcessMethod(left, set)
				if symbolTable[left][set] != nil {
					changed = true
				}
				// 没匹配到
				if symbolTable[left][set] == nil {
					for _, out := range rules.Rules[left] {
						if firstSet.IsInFirstSet(string(out[0]), set) {
							symbolTable[left][set] = &rule.Formula{
								Left:  left,
								Right: out,
							}
							changed = true
						}
					}
				}
			}
		}
		if !changed {
			break
		}
	}
	return symbolTable
}

func (s SymbolTable) String() string {
	column := []string{" "}
	for rowKey := range s {
		for colKey := range s[rowKey] {
			column = append(column, colKey)
		}
		break
	}
	table, err := gotable.Create(column...)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	for rowKey := range s {
		row := make(map[string]string)
		row[" "] = rowKey
		for col, formula := range s[rowKey] {
			if formula == nil {
				row[col] = ""
			} else {
				row[col] = fmt.Sprintf("%s->%s", transfer.Transfer(formula.Left), transfer.Transfer(formula.Right))
			}
		}
		err = table.AddRow(row)
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}
	}
	return table.String()
}
