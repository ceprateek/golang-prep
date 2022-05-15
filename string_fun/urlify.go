package string_fun

import (
	"fmt"
	"strings"
)

func urlify(s string) string {
	var result strings.Builder
	for _, v := range s {
		if v == ' '{
			result.WriteString("%20")
		} else {
			result.WriteRune(v)
		}
	}
	return result.String()
}

func PlayUrlify() {
	s := "Mr John Smith"
	result := urlify(s)
	fmt.Println(result)
}
