package main

import (
	"strings"
)

func ModifyArticle(text string) string {

	words := strings.Fields(text)
	vowels := "aeiouhAEIOUH"

	for i := 0; i < len(words); i++ {
		if words[i] == "a" && strings.ContainsAny(vowels, string(words[i+1][0])) {
			words[i] = "an"

		} else if words[i] == "A" && strings.ContainsAny(vowels, string(words[i+1][0])) {
			words[i] = "An"

		} else if words[i] == "an" && !strings.ContainsAny(vowels, string(words[i+1][0])) {
			words[i] = "a"

		} else if words[i] == "An" && !strings.ContainsAny(vowels, string(words[i+1][0])) {
			words[i] = "A"
		}
	}
	return strings.Join(words, " ")

}
