package main

type Organization struct {
	Exist      bool
	ProperName bool
	HaveNumber bool
}

type FinalOrg func(org *Organization) (bool, error)

func main() {

	a := func(org *Organization) {
		org.HaveNumber = true
	}
	b := func(org *Organization) {
		org.Exist = true
	}
	c := func(org *Organization) {
		org.ProperName = true
	}

	//i can call a like
	orgg := Organization{}
	a(&orgg)

	UpdateOrganization(a, b, c)

	org := Organization{}
	println(org.Exist)
	//CALL THE ORGANIZATION EXIST AND RETURN
	f := org.OrganizationExist()
	f(&org)
	println(org.Exist)
	// function type implements
	f.OrganizationReq(&org)
	println(org.Exist)
}

//here functions are passed and each one will be executed if we use it
//function as arguement
func UpdateOrganization(functions ...func(organization *Organization)) bool {
	org := Organization{}
	for _, f := range functions {
		f(&org)
	}
	println(org.ProperName, org.Exist, org.HaveNumber)
	return org.HaveNumber
}

//return type as function type FinalOrg type
// retyurn is function
func (org Organization) OrganizationExist() FinalOrg {
	return func(org *Organization) (bool, error) {
		org.Exist = !org.Exist
		return org.Exist, nil
	}
}

//function type implements
func (fo FinalOrg) OrganizationReq(organization *Organization) {
	// here executing the function with org
	fo(organization)
}

