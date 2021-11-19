package architecture

import (
	"testing"
)

type Db map[int]Person

func (m Db) Save(key int, p Person) {
	m[key] = p
}

func (m Db) Retrieve(key int) Person {
	return m[key]
}

func TestPutWithMongo(t *testing.T) {
	db := Db{}
	serv := NewPersonService(db)
	p := Person{First: "Alex"}
	serv.Put(2, p)
	got := db.Retrieve(2)
	if got != p {
		t.Fatalf("Want %v, got %v", p, got)
	}
}
