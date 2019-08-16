package main

type InformationService interface {
	CreateInfo()
}


type AddressService interface {
	CreateAddress()
}
type EmployeeService interface {
	InformationService
	AddressService
}
type EmployeeDTO struct {
	EmpName string
}
func (e EmployeeDTO) CreateInfo() {

}

func (e EmployeeDTO) CreateAddress() {

}
//here EmployeeDto should implement both the methods of InformationService and AddressService
//because EmployeeService is embedded interface
func main() {
	var d EmployeeService = EmployeeDTO{EmpName:"Gowtham"}
	println(d)
}
