package Prototype_Pattern

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	f1 := File{name: "file1"}
	fc := Floder{name: "floderchild"}

	ff := Floder{name: "floderfather"}

	ff.childen = append(ff.childen, &fc)
	ff.childen = append(ff.childen, &f1)

	f1.Print()
	fc.Print()
	ff.Print()

	fcopy := ff.Clone()

	fcopy.Print()
}
