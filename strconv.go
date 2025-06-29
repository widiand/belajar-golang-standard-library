package main

import (
	"fmt"
	"strconv"
)

func main() {
	resultBool, err := strconv.ParseBool("true") //string ke bool
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultBool)
	}

	resultInt, err := strconv.Atoi("10") //string ke int (A = string, i = int, a to i)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}
	fmt.Println(strconv.Itoa(999)) // int ke string (i to a)

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)
}
