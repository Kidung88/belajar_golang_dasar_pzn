package main

import "fmt"

func main() {
	var name = "Pikatan"

	if name == "Kidung" {
		fmt.Println("Hello Kidung")
	} else if name == "Arya" {
		fmt.Println("Hello Arya")
	} else if name == "Pikatan" {
		fmt.Println("Hello Pikatan")
	} else {
		fmt.Println("Hi, kenalan donk")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
