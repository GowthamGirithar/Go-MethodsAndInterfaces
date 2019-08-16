package main

import "math"

//shape have the method Area
type ShapeI interface {
	Area() float32
}

//rect and circle are the types of struct
type Rectangle struct {
	width  float32
	height float32
}

type Circ struct {
	radius float32
}


func main() {
	var shape ShapeI= Rectangle{
		width:  10,
		height: 10,
	}

	// i.(Type) -> i is the interface and Type is the implementation
	data , ok := shape.(Rectangle)
	println("the type assertion is ok " , ok) // true becuase we have assigned the circle with values to shape
	if ok{
		println("the area of rectangle is " , data.Area())
	}

	data1 , ok1 := shape.(Circ)  // here interface shapeI dont have the circle concrete values , so it will give panic
	println("the type assertion of circle is  ok " , ok1)  // false
	if ok{
		println("the area of rectangle is " , data1.Area())
	}
	shape= Circ{radius:10}
	data2 , ok2 := shape.(Circ)  // here interface shapeI dont have the circle concrete values , so it will give panic
	println("the type assertion of circle after assigning is  ok " , ok2)  // true
	if ok{
		println("the area of circle after assigning  is " , data2.Area())
	}



}
// Area calculates the height * width
// Rect have implemented the Shape
func (rect Rectangle) Area() float32 {
	return rect.height * rect.width
}
// circle calculates the 2*pi*r
// circle have implemented the Shape
func (circ Circ) Area() float32 {
	return 2 * math.Pi * circ.radius
}

/**
In type assertion syntax i.(Type), if Type does not implement the interface (the type of) i then
Go compiler will throw an error. But if Type implements the interface but i does not have a
concrete value of Type then Go will panic in runtime. Luckily, there is another variant of type assertion syntax, which is
value, ok := i.(Type)
 */


