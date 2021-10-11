package Iterator_Pattern

import "testing"

func TestIteratorPattern(t *testing.T) {
	a := IntArray{IArr: []int{1, 2, 3, 5}}
	var it IntIterator
	it = &IntIter{idx: 0, I: &a}
	it = it.Next()
	t.Logf("it next %d", it.Current())
}
