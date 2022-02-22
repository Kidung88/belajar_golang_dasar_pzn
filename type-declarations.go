package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpEko NoKTP = "3302250511990002"
	var marriedStatus Married = false
	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)
}
