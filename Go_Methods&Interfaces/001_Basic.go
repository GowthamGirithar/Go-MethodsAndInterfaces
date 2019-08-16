package main

type InvoiceDto struct {
	InvoideGuid string
}

// empty interface
type Interface1 interface {
}

//interface with two methods
type Interface2 interface {
	Create()
	Update()
}

//interface with one method
type Interface3 interface {
	Delete()
}

//composite interface
type Interface4 interface {
	Interface2
	Interface3
}

func main() {
	var data Interface2 = InvoiceDto{InvoideGuid: "12345"} // this is not valid unless InvoiceDto implement
	// you cannot assign to interface unless it implements
	// like how in java , we have to implement all the interfaces
	println(data) // print address
	data.Create()
	data.Update()

	var dataInterface2 Interface2
	println(dataInterface2 == nil) // true

}

func (invoice InvoiceDto) Create() {
	println("Create Invoice method")
}

func (invoice InvoiceDto) Update() {
	println("Update Invoice method")
}

//learning
// you cannot assign to interface unless it implements
// like how in java , we have to implement all the interfaces

