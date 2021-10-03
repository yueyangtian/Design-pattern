package Chain_of_Responsibility

import "testing"

func TestChainofResponsibilityPattern(t *testing.T) {
	l := GetLogger()
	l.Log(1, "hello")
}
