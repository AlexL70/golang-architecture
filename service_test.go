package architecture

import (
	// "fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

type Db map[int]Person

func TestPut(t *testing.T) {
	ctl := gomock.NewController(t)
	acc := NewMockAccessor(ctl)
	serv := NewPersonService(acc)
	p := Person{First: "Alex"}
	acc.EXPECT().Save(2, p).MinTimes(1).MaxTimes(1)
	serv.Put(2, p)
	ctl.Finish()
}

// func ExamplePut() {
// 	db := Db{}
// 	serv := NewPersonService(db)
// 	p := Person{First: "Alex"}
// 	serv.Put(2, p)
// 	got := db.Retrieve(2)
// 	fmt.Println(got)
// 	// Output: {Alex}
// }
