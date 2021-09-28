package Abstract_Factory_Pattern

import (
	"testing"
)

func TestDoctor(t *testing.T) {
	var c WorkerCreate
	c = new(DoctorCreator)
	p := c.Create()
	t.Log(p.Work())
}

func TestProgramer(t *testing.T) {
	var c WorkerCreate
	c = new(ProgramerCreator)
	p := c.Create()
	t.Log(p.Work())
}
