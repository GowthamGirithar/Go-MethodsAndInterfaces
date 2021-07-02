package main

import "fmt"

type Pol1 struct{

}
type Pol2 struct{

}

type Im interface {
	getTags() string
}

func (p Pol1) getTags() string{
	return "hello"
}
func main(){

	var p interface{}=Pol1{}  //p should be of interface type to check this
	if v,ok := p.(Im) ; ok{
		fmt.Print(v.getTags())
	}
	switch c:=p.(type){
	default:
		fmt.Println("hello p ",c) //hello p  {}
	}
	var p2 Pol1
	Test(p2)

	var p1 interface{}=Pol2{}
	if v,ok := p1.(Im) ; ok{
		fmt.Print(v.getTags())
	}

}


func Test(e Im){ //here it should be an interface (not an empty interface)
	switch c:=e.(type){
	case Pol1:
		fmt.Println("Its Pol1") //gets printed
	default:
		fmt.Println("Type is ",c)
	}
}

