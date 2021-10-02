package Composite_Pattern

import (
	"testing"
)

func TestCompositePattern(t *testing.T) {
	l1 := &ComLeaf{name: "l1"}
	l2 := &ComLeaf{name: "l2"}
	l3 := &ComLeaf{name: "l3"}

	n0 := &ComFloder{name: "n0"}

	n0.Add(l1)
	n0.Add(l2)
	n0.Add(l3)

	n1 := &ComFloder{name: "n1"}

	l11 := &ComLeaf{name: "l11"}
	l12 := &ComLeaf{name: "l12"}
	l13 := &ComLeaf{name: "l13"}

	n1.Add(l11)
	n1.Add(l12)
	n1.Add(l13)

	n0.Add(n1)

	n0.search("l11")
	n0.search("l1")
	n0.search("l4")
	n0.search("l14")
}
