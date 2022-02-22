package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string){
	fmt.Println("Hello", name, ".My name is", customer.Name)
}

func (a Customer) sayHuuu(){
	fmt.Println("Hi from", a.Name)
}

func main() {
	var kidung Customer
	kidung.Name = "Kidung"
	kidung.Address = "Indonesia"
	kidung.Age = 22

	kidung.sayHi("Arya")
	kidung.sayHuuu()

	//fmt.Println(kidung.Name)
	//fmt.Println(kidung.Address)
	//fmt.Println(kidung.Age)
	//
	//arya := Customer{
	//	Name:    "Arya",
	//	Address: "Subang",
	//	Age:     22,
	//}
	//fmt.Println(arya)
	//
	//pikatan := Customer{"Pikatan", "Jakarta", 22}
	//fmt.Println(pikatan)
}
