package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type ByAge []User

func (s ByAge) Len() int {
	return len(s)
}

func (s ByAge) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s ByAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Eko", 30},
		{"Budi", 35},
		{"Joko", 25},
		{"Adit", 23},
	}

	sort.Sort(ByAge(users))
	fmt.Println(users)

	sort.Reverse(ByAge(users))
	fmt.Println(users)
}