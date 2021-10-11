package Iterator_Pattern

type IntIterator interface {
	Next() IntIterator
	Current() int
	Frist() IntIterator
}

type IntArray struct {
	IArr []int
}

type IntIter struct {
	idx int
	I   *IntArray
}

func (this *IntIter) Next() IntIterator {
	if len(this.I.IArr) > this.idx {
		this.idx++
	}
	return this
}

func (this *IntIter) Frist() IntIterator {
	this.idx = 0
	return this
}

func (this *IntIter) Current() int {
	return this.I.IArr[this.idx]
}
