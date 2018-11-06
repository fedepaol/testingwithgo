package testingwithgo

import (
	"strings"
)

// START OMIT
func Reverse(s string) string {
	b := strings.Builder{}
	for i := range s {
		b.WriteByte(s[len(s)-1-i])
	}
	return b.String()
}

// END OMIT
