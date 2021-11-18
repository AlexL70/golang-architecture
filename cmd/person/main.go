package main

import (
	"fmt"

	architecture "github.com/AlexL70/golang-architecture"
	"github.com/AlexL70/golang-architecture/storage/mongo"
	"github.com/AlexL70/golang-architecture/storage/postgres"
)

func main() {
	p1 := architecture.Person{
		First: "Jenny",
	}

	p2 := architecture.Person{
		First: "James",
	}

	//	store to mongo
	srvM := architecture.NewPersonService(mongo.Db{})
	srvM.Put(1, p1)
	srvM.Put(2, p2)
	//	store to postgress
	srvP := architecture.NewPersonService(postgres.Db{})
	srvP.Put(1, p1)
	srvP.Put(2, p2)
	//	retrieve from mongo
	fmt.Println("mongo")
	fmt.Println(srvM.Get(1))
	fmt.Println(srvM.Get(2))
	fmt.Println(srvM.Get(3))
	//	retrieve from postgress
	fmt.Println("postgress")
	fmt.Println(srvP.Get(1))
	fmt.Println(srvP.Get(2))
	fmt.Println(srvP.Get(3))
}
