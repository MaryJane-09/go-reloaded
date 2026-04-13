package main

import (
	"strconv"
	"strings"
)

func ModifyCases(text string) string {

	words := strings.Fields(text)

	for i := len(words) -1; i >= 0; i-- {
		if strings.HasPrefix(words[i], "(up,") && i+1 < len(words) {
			index := strings.TrimSuffix(words[i+1], ")")
			num, err := strconv.Atoi(index)

			if err != nil {
				continue
			}

			for j := 1; j <= num; j++ {
				if i-j >= 0 {
					words[i-j] = strings.ToUpper(words[i-j])
				}
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}

		if strings.HasPrefix(words[i], "(low,") && i+1 < len(words) {
			index := strings.TrimSuffix(words[i+1], ")")
			num, err := strconv.Atoi(index)

			if err != nil {
				continue
			}

			for j := 1; j <= num; j++ {
				if i-j >= 0 {
					words[i-j] = strings.ToLower(words[i-j])
				}
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}

		if strings.HasPrefix(words[i], "(cap,") && i+1 < len(words) {
			index := strings.TrimSuffix(words[i+1], ")")
			num, err := strconv.Atoi(index)

			if err != nil {
				continue
			}

			for j := 1; j <= num; j++ {
				if i-j >= 0 {
					words[i-j] = strings.ToUpper(string(words[i-j][0])) + strings.ToLower(words[i-j][1:])
				}
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}

		if words[i] == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}

		if words[i] == "(cap)" && i > 0 {
			words[i-1] = strings.ToUpper(string(words[i-1][0])) + strings.ToLower(words[i-1][1:])
			words = append(words[:i], words[i+1:]...)
		}

	}
	return strings.Join(words, " " )
	
}
