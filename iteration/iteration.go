package iteration

import "strings"

const repeatCount = 5

func Repeat(character string, num int) string {
	var repeated strings.Builder
	for i := 0; i < num; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
