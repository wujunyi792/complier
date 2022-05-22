package grammar

import (
	"bytes"
	"compiler/lexer"
	"compiler/util"
	"fmt"
	"io"
	"strings"
)

const EndToken = "$"

type Grammar struct {
	*util.Stream
	productions map[string][]string
	endToken    string
	tokens      *util.Queue
}

func makeProductions() map[string][]string {
	res := make(map[string][]string)
	res["E"] = []string{"TG"}
	res["G"] = []string{"ATG", "&"}
	res["T"] = []string{"FS"}
	res["S"] = []string{"MFS", "&"}
	res["F"] = []string{"(E)", "i"}
	res["A"] = []string{"+", "-"}
	res["M"] = []string{"/", "*"}
	return res
}

func Analyse(raw []*lexer.Token) ([]*Production, bool) {
	grm := NewGrammar(raw, bytes.NewBufferString("E"), EndToken)
	productions, success := grm.Analyse(1)
	if !success {
		return nil, false
	}
	e := grm.tokens.FrontRaw()
	if e != nil {
		if e.Value.(*lexer.Token).Value == "" && e.Next() == nil {
			return productions, true
		}
		return nil, false
	}

	return productions, true
}

func NewGrammar(token []*lexer.Token, r io.Reader, et string) *Grammar {
	s := util.NewStream(r, EndToken)
	q := util.New()
	for i := 0; i < len(token); i++ {
		q.PushBack(token[i])
	}
	return &Grammar{Stream: s, endToken: et, productions: makeProductions(), tokens: q}
}

func (g *Grammar) frontMatch() bool {
	return true
}

func (g *Grammar) GetNextOrigin() string {
	c := g.Stream.Next()
	lookahead := g.Stream.Peek()
	if lookahead == "'" {
		return c + g.Stream.Next()
	}
	return c
}

func (g *Grammar) isEndType(r string) bool {
	dic := map[string]bool{
		"i": true,
		"+": true,
		"-": true,
		"*": true,
		"/": true,
		"(": true,
		")": true,
		"&": true,
	}
	return dic[r]
}

func (g *Grammar) Analyse(count int) (res []*Production, canKill bool) {

	// 当前推导式头部
	origin := g.GetNextOrigin()
	if origin == "&" {
		res = append(res, &Production{
			Type:   "kill",
			Target: "&",
		})
		g.tokens.PushFront(&lexer.Token{})
		return res, true
	}
	if g.isEndType(origin) {
		// 当前产生式首位为终结符
		c := g.tokens.Front()
		if c == nil {
			return res, false
		}
		currentToken := c.(*lexer.Token)
		switch origin {

		case "i":
			{
				if !currentToken.IsValue() {
					g.Stream.ClearFronts(count - 1)
					return res, false
				}
				res = append(res, &Production{
					Type:   "kill",
					Target: "i",
				})
				return res, true
			}
		case "+":
			{
				if !currentToken.IsOperator() || currentToken.Value != "+" {
					g.Stream.ClearFronts(count - 1)
					return res, false
				}
				res = append(res, &Production{
					Type:   "kill",
					Target: "+",
				})
				return res, true
			}
		case "-":
			{
				if !currentToken.IsOperator() || currentToken.Value != "-" {
					g.Stream.ClearFronts(count - 1)
					return res, false
				}
				res = append(res, &Production{
					Type:   "kill",
					Target: "-",
				})
				return res, true
			}
		case "*":
			{
				if !currentToken.IsOperator() || currentToken.Value != "*" {
					g.Stream.ClearFronts(count - 1)
					return res, false
				}
				res = append(res, &Production{
					Type:   "kill",
					Target: "*",
				})
				return res, true
			}
		case "/":
			{
				if !currentToken.IsOperator() || currentToken.Value != "/" {
					g.Stream.ClearFronts(count - 1)
					return res, false
				}
				res = append(res, &Production{
					Type:   "kill",
					Target: "/",
				})
				return res, true
			}
		case "(":
			{
				if !currentToken.IsBracket() || currentToken.Value != "(" {
					g.Stream.ClearFronts(count - 1)
					return res, false
				}
				res = append(res, &Production{
					Type:   "kill",
					Target: "(",
				})
				return res, true
			}
		case ")":
			{
				if !currentToken.IsBracket() || currentToken.Value != ")" {
					g.Stream.ClearFronts(count - 1)
					return res, false
				}
				res = append(res, &Production{
					Type:   "kill",
					Target: ")",
				})
				return res, true
			}
		}
	}
	canUse := g.productions[origin]
	var killToken []*lexer.Token
	for i := 0; i < len(canUse); i++ {
		for _, s := range reverse(strings.Split(canUse[i], "")) {
			g.Stream.PutBack(s)
		}
		//g.Stream.Print()
		//g.PrintToken()
		match := false
		if canUse[i] == "" {
			match = true
		}
		for j := 0; j < len(canUse[i]); j++ {
			next := len(canUse[i]) - j
			ps, kill := g.Analyse(next)
			if kill {
				res = append(res, &Production{
					Origin: origin,
					Next:   canUse[i],
					Type:   "Continue",
				})
				res = append(res, ps...)
				if j != len(canUse[i])-1 {
					p := g.tokens.Pop()
					if p != nil {
						killToken = append(killToken, p.(*lexer.Token))
					}
				}
				//g.Stream.Print()
				//g.PrintToken()
			} else {
				for _, s := range reverseAny(killToken) {
					g.tokens.PushFront(s)
				}
				break
			}
			if j == len(canUse[i])-1 {
				match = true
			}
		}
		if !match && i == len(canUse)-1 {
			break
		}
		if match {
			if canUse[i] == "" {
				return res, true
			}
			return res, true
		}

	}

	g.Stream.ClearFronts(count - 1)
	return res, false
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func reverseAny(s []*lexer.Token) []*lexer.Token {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (g *Grammar) PrintToken() {
	e := g.tokens.FrontRaw()
	for {
		if e != nil {
			fmt.Print(e.Value.(*lexer.Token).Value)
			e = e.Next()
		} else {
			break
		}
	}
	fmt.Println("$\n ")
}
