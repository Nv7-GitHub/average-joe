package chain

import "strings"

var punctuation = []byte{'.', ',', '?', '!', '\'', '*', '`', '~', '_', '\n', ':'}

func simplify(msg string) string {
	msg = strings.ToLower(msg)
	// Remove punctuation
	for _, val := range punctuation {
		msg = strings.ReplaceAll(msg, string(val), "")
	}
	return msg
}
