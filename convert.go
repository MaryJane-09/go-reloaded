package main

import (
	"strconv"
	"strings"
)

func ConvertToDecimal(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			num, err := strconv.ParseInt(words[i-1], 16, 64)

			if err == nil {
				words[i-1] = strconv.FormatInt(num, 10)
			}
			words = append(words[:i], words[i+1:]...)
		}

		if words[i] == "(bin)" && i > 0 {
			num, err := strconv.ParseInt(words[i-1], 2, 64)

			if err == nil {
				words[i-1] = strconv.FormatInt(num, 10)
			}
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")

}
