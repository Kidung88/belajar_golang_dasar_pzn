package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Kidung")
	// helper.sayGoodbye("Kidung") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}
