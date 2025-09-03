package base

import "sort"

type BaseCompleter struct {
	data map[string]int
}

func (b BaseCompleter) Complete(prefix string) []string {
	matches := make([]string, 0)

	for w := range b.data {
		if len(w) >= len(prefix) && w[:len(prefix)] == prefix {
			matches = append(matches, w)
		}
	}

	sort.Strings(matches)
	return matches[:10]
}
func New(data map[string]int) BaseCompleter {
	return BaseCompleter{data: data}
}
