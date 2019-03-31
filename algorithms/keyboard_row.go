package algorithms

import (
	"strings"
	"fmt"
)

func findWords(words []string) []string {
	rows := []string{
		"qwertyuiop",
		"asdfghjkl",
		"zxcvbnm",
	}

	var rs []string

	for _, word := range words {

		for _, row := range rows {
			matched := compWord(word, row)
			if matched {
				rs = append(rs, word)
				break
			}
		}

	}

	return rs
}

func compWord(word, row string) bool {
	for _, wordC := range word {
		var matched bool

		for _, rowC := range row {
			wr := fmt.Sprintf("%c", wordC)
			rr := fmt.Sprintf("%c", rowC)
			if wr == rr || wr == strings.ToUpper(rr) {
				matched = true
			}
		}

		if !matched {
			return false
		}
	}

	return true
}