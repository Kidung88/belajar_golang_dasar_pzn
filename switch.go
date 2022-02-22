package main

import "fmt"

func main() {

	name := "Arya"

	switch name {
	case "Kidung":
		fmt.Println("Hello Kidung")
		fmt.Println("Hello Kidung")
	case "Arya":
		fmt.Println("Hello Arya")
		fmt.Println("Hello Arya")
	default:
		fmt.Println("Kenalan Donk")
		fmt.Println("Kenalan Donk")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah benar")
	//}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
