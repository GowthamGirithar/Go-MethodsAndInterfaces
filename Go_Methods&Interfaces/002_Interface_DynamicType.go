package main

import "math"

//shape have the method Area
type Shape interface {
	Area() float32
}

//rect and circle are the types of struct
type Rect struct {
	width  float32
	height float32
}
type Circle struct {
	radius float32
}

func main() {
	var data Shape=Rect{
		width:  10,
		height: 10,
	}
	println("the area or rectangle is ",data.Area())
	data = Circle{radius:10}
	println("the area of circle  is ",data.Area())

	// so in runtime it dynamically determine which to call based on the type assigned
}
// Area calculates the height * width
// Rect have implemented the Shape
func (rect Rect) Area() float32 {
	return rect.height * rect.width
}
// circle calculates the 2*pi*r
// circle have implemented the Shape
func (circ Circle) Area() float32 {
	return 2 * math.Pi * circ.radius
}
