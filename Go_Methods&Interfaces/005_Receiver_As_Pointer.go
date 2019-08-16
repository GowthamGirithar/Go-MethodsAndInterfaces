package main

//  interface which have the methods to create the org and update
type OrgService interface {
	 CreateOrg()
	 UpdateOrg()
}
//the Organization
type Organization struct {
	OrgName string
}

//we have till now seen methods with value receiver
//in this we will see method with pointer receiver
func main() {
	// var data OrgService=Organization{OrgName:"TEST"}   -> COMPILE TIME ERROR SAYING CreateOrg has pointer receiver
	d := Organization{OrgName:"test"}
	d.CreateOrg()  // this will work because method with pointer receiver can accept pointer or value

	//But in case of interfaces, if a method has a pointer receiver,
	// then the interface will have a pointer of dynamic type rather than the value of
	// dynamic type. Hence, instead of assigning a value to an interface variable,
	// we need to assign a pointer.
	var data OrgService=&Organization{OrgName:"TEST"}
	data.CreateOrg()
	data.UpdateOrg()

}
//pointer receiver
func (org *Organization) CreateOrg(){
	println("Create Org")
}
//value receiver
func (org Organization) UpdateOrg(){
	println("Update Org")
}



