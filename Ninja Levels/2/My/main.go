//	Task: define interface and use it
package main

import (
	"fmt"
	"math"
	"strconv"
)

type Figure interface {
	ToString() string
	Area() float64
}

type rectancle struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (r rectancle) ToString() string {
	return fmt.Sprintf("rectangle of length %s and width %s", strconv.FormatFloat(r.length, 'f', -1, 64), strconv.FormatFloat(r.width, 'f', -1, 64))
}

func (r rectancle) Area() float64 {
	return r.length * r.width
}

func (c circle) ToString() string {
	return fmt.Sprintf("circle of radius %s", strconv.FormatFloat(c.radius, 'f', -1, 64))
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func PrintArea(f Figure) {
	fmt.Println("A", f.ToString(), "has area equal to", f.Area())
}

func main() {
	PrintArea(rectancle{length: 4, width: 3})
	PrintArea(circle{radius: 2.5})
}
