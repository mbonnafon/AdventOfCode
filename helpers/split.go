package helpers

import "strings"

func SplitKeyValue(s, delimiter string) (string, string) {
	kv := strings.Split(s, delimiter)
	return kv[0], kv[1]
}
