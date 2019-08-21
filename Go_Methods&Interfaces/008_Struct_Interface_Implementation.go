package main

type Invoice struct {
	Create Create
	Update Update
}

type Create func(invoice *Invoice) (bool, error)
type Update func(invoice *Invoice, data string) (bool, error)

func NewInvoiceCreate(invoice *Invoice) Create {
	return func(invoice *Invoice) (b bool, e error) {
		println("Inside the Invoice function")
		return false, nil
	}
}

type API struct {
	Invoice Invoice
	Dispute Dispute
}

type Dispute struct {
}

func main() {
	inv := Invoice{}
	apiClient := &API{
		Invoice: Invoice{
			Create: NewInvoiceCreate(&inv),
		},
	}
	//way to call the implementation
	f1 :=apiClient.Invoice.Create  // will not execute the function
	res, _ := f1(&inv)
	println(res)
	//alternate way
	f , _:= apiClient.Invoice.Create(&inv) // will execute function f(&inv)
	println(f1 , f)

}
