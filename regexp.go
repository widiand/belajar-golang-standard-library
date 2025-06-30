package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`e([a-z])o`)
	
	fmt.Println(regex.MatchString("eko")) //true
	fmt.Println(regex.MatchString("edo")) //true
	fmt.Println(regex.MatchString("eKo")) //false

	fmt.Println(regex.FindAllString("eko edo egi ego e1o eto eKo", 10))
}