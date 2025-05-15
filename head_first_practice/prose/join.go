package prose

import (
	"fmt"
	"strings"
)

func JoinWithCommas(phrases []string) string {
	if len(phrases) == 0 {
		return ""
	} else if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	}
	res := strings.Join(phrases[:len(phrases)-1], ", ")
	res += ", and "
	res += phrases[len(phrases)-1]
	return res
}

func Main() {
	fmt.Println(JoinWithCommas([]string{"a", "b", "c", "d"}))
	fmt.Println(JoinWithCommas([]string{"a"}))
	fmt.Println(JoinWithCommas([]string{"a", "b"}))
	// fmt.Println(strings.Join([]string{"a", "b", "c"}, ","))
}
