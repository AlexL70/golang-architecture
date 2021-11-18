package mongo

import architecture "github.com/AlexL70/golang-architecture"

type Db map[int]architecture.Person

func (m Db) Save(key int, p architecture.Person) {
	m[key] = p
}

func (m Db) Retrieve(key int) architecture.Person {
	return m[key]
}
