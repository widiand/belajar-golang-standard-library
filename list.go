package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("khannedy")

	head := data.Front()
	fmt.Println(head.Value) //eko
	
	next := head.Next()
	fmt.Println(next.Value) //Kurniawan
	
	next = next.Next() 
	fmt.Println(next.Value) //Khannedey

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}