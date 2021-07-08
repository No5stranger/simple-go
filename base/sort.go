package base

import (
	"sort"
)

type Pair struct {
	K int64
	V int64
}

type Pairs []Pair

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	return p[i].V < p[j].V
}

func sortMapByValue(m map[int64]int64) Pairs {
	var p Pairs
	for k, v := range m {
		p = append(p, Pair{K: k, V: v})
	}
	sort.Sort(p)
	return p
}
