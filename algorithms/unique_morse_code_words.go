package algorithms

import "fmt"

var morseMaps = map[string]string{
	"a": ".-",
	"b": "-...",
	"c": "-.-.",
	"d": "-..",
	"e": ".",
	"f": "..-.",
	"g": "--.",
	"h": "....",
	"i": "..",
	"j": ".---",
	"k": "-.-",
	"l": ".-..",
	"m": "--",
	"n": "-.",
	"o": "---",
	"p": ".--.",
	"q": "--.-",
	"r": ".-.",
	"s": "...",
	"t": "-",
	"u": "..-",
	"v": "...-",
	"w": ".--",
	"x": "-..-",
	"y": "-.--",
	"z": "--..",
}

func uniqueMorseRepresentations(words []string) int {
	var morseList []string
	var distincts = make(map[string]string)

	for _, word := range words {
		var morseConverted string
		for _, char := range word {
			k := fmt.Sprintf("%c", char)
			morseConverted = fmt.Sprintf("%s%s", morseConverted, morseMaps[k])
		}
		morseList = append(morseList, morseConverted)
		morseConverted = ""
	}

	for _, v := range morseList {
		if distincts[v] == "" {
			distincts[v] = v
		}
	}
	return len(distincts)
}