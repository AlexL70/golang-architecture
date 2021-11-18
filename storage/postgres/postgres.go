package postgres

import architecture "github.com/AlexL70/golang-architecture"

type Db map[int]architecture.Person

func (pg Db) Save(key int, p architecture.Person) {
	pg[key] = p
}

func (pg Db) Retrieve(key int) architecture.Person {
	return pg[key]
}
