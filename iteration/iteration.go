package iteration

import (
	"strings"
)

const repeatCount = 1000000

func Repeat(character string) string {

	var repeated strings.Builder

	for range repeatCount {
		repeated.WriteString(character)
	}

	//go log.Printf("repeated %q %d times", character, repeatCount)
	return repeated.String()

}
