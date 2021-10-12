package Menento_Pattern

import "testing"

func TestMenentoPattern(t *testing.T) {
	edit := &TextEdit{}

	edit.Append("hello ")
	edit.Append("world ")
	edit.Show()
	edit.Revert()
	edit.Show()

}
