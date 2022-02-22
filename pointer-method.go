package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	kidung := Man{"Kidung"}
	kidung.Married()

	fmt.Println(kidung.Name)
}
