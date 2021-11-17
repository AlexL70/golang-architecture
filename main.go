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

func put(a accessor, key int, p person) {
	a.save(key, p)
}

func get(a accessor, key int) person {
	return a.retrieve(key)
}

func main() {
	dbm := mongo{}
	dbp := postg{}

	p1 := person{
		first: "Jenny",
	}

	p2 := person{
		first: "James",
	}

	//	store to mongo
	put(dbm, 1, p1)
	put(dbm, 2, p2)
	//	store to postgress
	put(dbp, 1, p1)
	put(dbp, 2, p2)
	//	retrieve from mongo
	fmt.Println("mongo", get(dbm, 1), get(dbm, 2))
	//	retrieve from postgress
	fmt.Println("postgress", get(dbp, 1), get(dbp, 2))
}
