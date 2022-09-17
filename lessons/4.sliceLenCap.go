package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	s = append(s, 6)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
