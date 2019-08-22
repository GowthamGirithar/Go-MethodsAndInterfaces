package main

type Receipt interface {
	perform(data string)
}

type InvoiceOperation struct {
	Status string
	Receipt Receipt
}

func (inv *InvoiceOperation) perform(data string){
	println(data)
}

func receiptProcess(receipt Receipt, data string){
	receipt.perform(data)
}
type TestRec interface {
	processData(data string)
}

func main() {
    var rec Receipt=&InvoiceOperation{
		Status:  "s",
		Receipt: nil,
	}
	receiptProcess(rec, "test")
}
