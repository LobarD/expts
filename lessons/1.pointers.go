package main

import "fmt"

func main() {
	st := "book"
	fmt.Println(st)

	p := &st
	fmt.Println(p)
	fmt.Println(*p) // *p указывает на st

	//*p = "lb" // *p  тоже самое что и st
	*p += " notebook"
	fmt.Println(st)
}
