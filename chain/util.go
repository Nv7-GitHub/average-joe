package chain

import "strings"

var punctuation = []byte{'.', ',', '?', '!', '\'', '*', '`', '~', '_', '\n', '"'}
var startendpunctuation = []byte{':'}

func simplify(msg string) string {
	msg = strings.ToLower(msg)
	// Remove punctuation
	for _, val := range punctuation {
		msg = strings.ReplaceAll(msg, string(val), "")
	}

	// Start end punctuation
	containse := false
	for _, val := range startendpunctuation {
		if strings.Contains(msg, string(val)) {
			containse = true
		}
	}
	if containse {
		words := strings.Split(msg, " ")
		for i, word := range words {
			if len(word) == 0 {
				continue
			}

			changed := false
			if word[0] == ':' {
				word = word[1:]
				changed = true
			}
			if word[len(word)-1] == ':' {
				word = word[:len(word)-1]
				changed = true
			}
			if changed {
				words[i] = word
			}
		}
	}

	return msg
}
