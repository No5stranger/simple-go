package base

import (
	"testing"
)

func TestSortMap(t *testing.T) {
	var m map[int64]int64 = map[int64]int64{
		1: 8,
		2: 6,
		4: 7,
	}
	p := sortMapByValue(m)
	t.Log(p)
}
