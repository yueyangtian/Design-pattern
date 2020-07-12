package DesginPattern

import (
	"testing"
)

func TestFilter(t *testing.T) {
	var p PersonList
	p.Add("jauny", "male", 15)
	p.Add("jack", "famale", 20)

	var fliter MaleCriteria

	t.Log(fliter.Criteria(p))

}
