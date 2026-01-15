package iteration

import "strings"

func Repeat(s string, c int) string {
	var repeated strings.Builder
	for range c {
		repeated.WriteString((s))
	}
	return repeated.String()
}
