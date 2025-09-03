package frequency

import "sort"

type FreqCompleter struct {
	data map[string]int
}

func (f FreqCompleter) Complete(prefix string) []string {
	matches := make([]string, 0)

	for w := range f.data {
		if len(w) >= len(prefix) && w[:len(prefix)] == prefix {
			matches = append(matches, w)
		}
	}

	sort.Slice(matches, func(i, j int) bool {
		wi, wj := matches[i], matches[j]
		if f.data[wi] == f.data[wj] {
			return wi < wj
		}
		return f.data[wi] > f.data[wj]
	})

	sort.Strings(matches)
	return matches[:10]
}
func New(data map[string]int) FreqCompleter {
	return FreqCompleter{data: data}
}
