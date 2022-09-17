package main

import "fmt"

func main() {
	s := make([]int, 2)
	s[0] = 0
	s[1] = 1
	printSlice(s) // len=2 cap=2 [0 1]

	//если мы добавляем сверх емкости текущего слайса, элементов меньше удвоенной емкости(cap), то cap увеличится в 2 раза
	s = append(s, 3, 4) //len=4 cap=4 [0 1 3 4]
	printSlice(s)
	//при заполнении слайса на количество элементов более чем удвоенная емкость(cap), то cap увеличивается в 2 раза + блоками по 2
	s = append(s, 5, 6, 7, 8, 9) //len=9 cap=10 [0 1 3 4 5 6 7 8 9]
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
