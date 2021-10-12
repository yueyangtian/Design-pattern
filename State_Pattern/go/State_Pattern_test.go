package State_Pattern

import "testing"

func TestStatePattern(t *testing.T) {
	v := Votor{}
	v.DoVote()
	v.DoVote()
}
