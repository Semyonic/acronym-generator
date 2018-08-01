package acronym

import (
	"fmt"
	"strings"
)

func Abbreviate(s string) string {
	var words []string
	// Split the given string without condition
	words = strings.Fields(s)
	var result string

	for i := 0; i < len(words); i++ {
		if strings.IndexAny(string(words[i]), "-") > 0 {
			// Split the string if it matches "-" char, convert them to upperCase
			split := strings.Split(strings.ToUpper(string(words[i])), "-")
			for j := 0; j < len(split); j++ {
				result += string(strings.ToUpper(string(split[j][0])))
			}
		} else {
			result += string(strings.ToUpper(string(words[i][0])))
		}
		fmt.Println(result)
	}
	return result
}

const (
	NAME = "Complementary metal-oxide semiconductor"
)

func acronym() {
	Abbreviate(NAME)
}
