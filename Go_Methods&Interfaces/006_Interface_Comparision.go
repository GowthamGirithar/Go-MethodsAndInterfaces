package main

import "reflect"

type A interface {

}

type B interface {

}
type Address struct {
	city string
	state string
	country string
}
type Emp struct {
	Name string
	Address
}

func main() {
	var a A
	var b B
	println(a == b) // true

	a= Emp{
		Name:    "Test",
		Address: Address{
			city:    "Test",
			state:   "Test",
			country: "Test",
		},
	}
	b= Emp{
		Name:    "Test",
		Address: Address{
			city:    "Test",
			state:   "Test",
			country: "Test",
		},
	}

	println(a == b)

	var aa interface{}
	var bb interface{}
	aa= []int{1}
	bb= [] int {1}

	println(reflect.DeepEqual(aa,bb))//true  - so use reflect deepequal for complexy values in implementation fields like arrays
	println(aa == bb) // will throw panic at run time. we can compare the interfaces if they are simple and identical
	//not for complex types.same application to struct also

}
