package testingwithgo

import "strings"

func Reverse(s string) string {
	b := strings.Builder{}
	for i := range s {
		b.WriteByte(s[len(s)-1-i])
	}
	return b.String()
}
