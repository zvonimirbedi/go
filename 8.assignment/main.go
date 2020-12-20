package main

import "fmt"

type shape interface {
	getArea() float64
}
type shapeTriangle struct {
	base   float64
	height float64
}
type shapeSquare struct {
	sideLength float64
}

func main() {

	// st := shapeTriangle{base: 2, height: 3}
	st := shapeTriangle{}
	st.base = 2
	st.height = 3

	// ss := shapeSquare{sideLength: 5}
	ss := shapeSquare{}
	ss.sideLength = 5

	printArea(st)
	printArea(ss)

}

func (t shapeTriangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
func (s shapeSquare) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
