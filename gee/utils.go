package gee

import (
	"strings"
)

func SplitTrieNodePath(pattern string) []string {
	result := make([]string, 0)

	for _, item := range strings.Split(pattern, "/") {
		if item == "" {
			continue
		}
		result = append(result, item)
		if item[0] == '*' {
			break
		}
	}

	return result
}
