package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("k([a-z])g")

	fmt.Println(regex.MatchString("kidong"))
	fmt.Println(regex.MatchString("kieong"))
	fmt.Println(regex.MatchString("kideog"))

	search := regex.FindAllString("kieong kideog kidong kieong", -1)
	fmt.Println(search)
}
