package main

import "fmt"

type Person struct {
	name  string
	age   int
	phone string
}

func main() {
	L := Person{"Lobar", 23, "+89999999"}
	fmt.Println(L)
	L.name += " E"
	fmt.Println(L.name)
	L.age = 66
	fmt.Println(L.age)
	fmt.Println(L)
}
