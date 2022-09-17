package main

import "fmt"

func main() {
	arrayStings := [6]string{"zero", "one", "twe", "tree", "four", "five"}
	fmt.Println(arrayStings)

	sliceOne := arrayStings[:4]
	sliceTwo := arrayStings[3:]
	fmt.Println(sliceOne)
	fmt.Println(sliceTwo)

	sliceTwo[0] = "---"
	fmt.Println(sliceOne)
	fmt.Println(sliceTwo)

	fmt.Println(arrayStings)
}
