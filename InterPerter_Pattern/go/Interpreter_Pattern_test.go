package Interpreter_Pattern

import "testing"

func TestInterpreterPattern(t *testing.T) {
	t.Log(Parser("6 100 11 + * 10 *"))
}
