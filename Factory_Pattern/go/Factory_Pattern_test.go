package Factory_Pattern

import (
	"testing"
)

func TestCircle(t *testing.T) {
	shap, err := Create("Cricle")
	if err != nil {
		t.Log("return val err:", err.Error())
	}
	t.Log("Shap:", shap.Draw())
}

func TestSquare(t *testing.T) {
	shap, err := Create("Square")
	if err != nil {
		t.Log("return val err:", err.Error())
	}
	t.Log("Shap:", shap.Draw())
}

func TestRectangle(t *testing.T) {
	shap, err := Create("Rectangle")
	if err != nil {
		t.Log("return val err:", err.Error())
	}
	t.Log("Shap:", shap.Draw())
}

func TestOther(t *testing.T) {
	shap, err := Create("Other")
	if err != nil {
		t.Log("return val err:", err.Error())
		return
	}
	t.Log("Shap:", shap.Draw())
}
