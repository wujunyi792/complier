package transfer

import "strings"

func Transfer(str string) string {
	//str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "G", "E'")
	str = strings.ReplaceAll(str, "S", "T'")
	str = strings.ReplaceAll(str, "&", "ε")
	return str
}
