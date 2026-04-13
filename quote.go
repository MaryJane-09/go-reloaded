package main

import (
	"strings"
)

func FixQuotes(text string) string {
	words := strings.Split(text, "'")

	for i := range words {
		if i%2 == 1 {
			words[i] = strings.TrimSpace(words[i])
		}

	}
	return strings.Join(words, "'")
}
