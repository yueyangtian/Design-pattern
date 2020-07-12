package DesginPattern

type PersonCriteria struct {
	Name string
	Age  int
	Sex  string
}

type PersonList struct {
	List []PersonCriteria
}

func (m *PersonList) Add(name, sex string, age int) {
	m.List = append(m.List, PersonCriteria{Name: name, Age: age, Sex: sex})
}

type Criteria interface {
	Criteria() []string
}

type MaleCriteria struct{}

func (s *MaleCriteria) Criteria(m PersonList) []string {
	res := []string{}
	for _, person := range m.List {
		if person.Sex == "male" {
			res = append(res, person.Name)
		}
	}
	return res
}
