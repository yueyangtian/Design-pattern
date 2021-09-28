package Abstract_Factory_Pattern

type Worker interface {
	Work() string
}
type WorkerCreate interface {
	Create() Worker
}

//doctor
type Doctor struct{}

func (*Doctor) Work() string {
	return "Doctor do something"
}

type DoctorCreator struct{}

func (*DoctorCreator) Create() Worker {
	return new(Doctor)
}

//programer
type Programer struct{}

func (*Programer) Work() string {
	return "Programer do something"
}

type ProgramerCreator struct{}

func (*ProgramerCreator) Create() Worker {
	return new(Programer)
}
