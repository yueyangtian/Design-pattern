package Strategy_Pattern

import "testing"

func TestStrategy(t *testing.T) {
	add := Context{Option: &AddOperation{}}
	t.Logf("1 + 2 = %d", add.Exec(1, 2))
	mutil := Context{Option: &MutilOperation{}}
	t.Logf("1 * 2 = %d", mutil.Exec(1, 2))
}
