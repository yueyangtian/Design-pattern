package Bridge_Pattern

import (
	"testing"
)

func TestBrige(t *testing.T) {
	hpPrinter := &hp_windows{}
	suf_w := &suf_windows{}

	win := &windows{}
	lins := &linux{}

	lins.setPrinter(hpPrinter)
	lins.print()

	win.setPrinter(suf_w)
	win.print()
}
