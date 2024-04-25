package dataProcessingUtils

import "unicode"

func RemoveNonInteger(strings ...*string) {
	if strings != nil {
		for i, str := range strings {
			strRunes := []rune(*str)

			newValue := ""
			for _, r := range strRunes {
				if unicode.IsDigit(r) {
					newValue = newValue + string(r)
				}
			}

			*strings[i] = newValue
		}
	}
}
