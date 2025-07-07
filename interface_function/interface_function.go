package interface_function

import "fmt"

type Shape interface {
	Area() float64
}
type Circle struct {
	radius float64
}
type Rectangle struct {
	width  float64
	height float64
}

func (r Circle) Area() float64 {
	area := 3.14 * r.radius * r.radius
	fmt.Println("Circle Area", area)
	return area
}
func (rect Rectangle) Area() float64 {
	area := rect.height * rect.width
	fmt.Println("Rectangle Area", area)
	return area
}
func printArea(s Shape) {
	fmt.Println("Calculated Area:", s.Area())
}
func InterfaceAll() {
	s := []Shape{
		Circle{radius: 2.5},
		Rectangle{height: 10.0, width: 25.0},
	}
	for _, shape := range s {
		printArea(shape)
	}
}

// Or Do this
func InterfaceFunction() {
	var s Shape
	cr := Circle{radius: 2.5}
	s = cr
	s.Area()
	rect := Rectangle{height: 2.5, width: 3.5}
	s = rect
	s.Area()

}
