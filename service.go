package architecture

import "fmt"

type Person struct {
	First string
}

type Accessor interface {
	Save(int, Person)
	Retrieve(int) Person
}

type PersonService struct {
	a Accessor
}

func NewPersonService(a Accessor) PersonService {
	return PersonService{a: a}
}

func (ps PersonService) Get(key int) (Person, error) {
	p := ps.a.Retrieve(key)
	if p.First == "" {
		return Person{}, fmt.Errorf("Person with ID %d not found", key)
	}
	return p, nil
}

func (ps PersonService) Put(key int, p Person) {
	ps.a.Save(key, p)
}
