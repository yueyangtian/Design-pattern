package DesginPattern

import (
	"testing"
)

func TestBrige(t *testing.T) {
	hpPrinter := &hp{}
	suf := &suf{}

	win := &windows{}
	lins := &linux{}

	lins.setPrinter(hpPrinter)
	lins.print()

	win.setPrinter(suf)
	win.print()
}
