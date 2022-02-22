package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Kidung Arya", "Kidung"))
	fmt.Println(strings.Contains("Kidung Pikatan", "Reza"))

	fmt.Println(strings.Split("Kidung Arya Pikatan", " "))

	fmt.Println(strings.ToLower("Kidung Arya Pikatan"))
	fmt.Println(strings.ToUpper("Kidung Arya Pikatan"))
	fmt.Println(strings.ToTitle("Kidung Arya Pikatan"))

	fmt.Println(strings.Trim("      Kidung Arya Pikatan     ", " "))
	fmt.Println(strings.ReplaceAll("Kidung Arya Kidung", "Kidung", "Reza"))
}
