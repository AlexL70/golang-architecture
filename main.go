package main

import "fmt"

type person struct {
	first string
}

type mongo map[int]person

func (m mongo) save(key int, p person) {
	m[key] = p
}

func (m mongo) retrieve(key int) person {
	return m[key]
}

type postg map[int]person

func (pg postg) save(key int, p person) {
	pg[key] = p
}

func (pg postg) retrieve(key int) person {
	return pg[key]
}

type accessor interface {
	save(int, person)
	retrieve(int) person
}

type personService struct {
	a accessor
}

func (ps personService) get(key int) (person, error) {
	p := ps.a.retrieve(key)
	if p.first == "" {
		return person{}, fmt.Errorf("Person with ID %d not found", key)
	}
	return p, nil
}

func (ps personService) put(key int, p person) {
	ps.a.save(key, p)
}

func main() {
	p1 := person{
		first: "Jenny",
	}

	p2 := person{
		first: "James",
	}

	//	store to mongo
	srvM := personService{a: mongo{}}
	srvM.put(1, p1)
	srvM.put(2, p2)
	//	store to postgress
	srvP := personService{a: postg{}}
	srvP.put(1, p1)
	srvP.put(2, p2)
	//	retrieve from mongo
	fmt.Println("mongo")
	fmt.Println(srvM.get(1))
	fmt.Println(srvM.get(2))
	fmt.Println(srvM.get(3))
	//	retrieve from postgress
	fmt.Println("postgress")
	fmt.Println(srvP.get(1))
	fmt.Println(srvP.get(2))
	fmt.Println(srvP.get(3))
}
