package main

import (
	"fmt"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"1", "2", "3", "4"},      // i = 0
		[]string{"5", "6", "7", "8", "9"}, // i = 1
		[]string{"10", "11", "12", "13"},  // i = 2
		[]string{"14", "15", "16", "17"},  // i = 3
	}

	for i := 0; i < len(board); i++ { //i = 0, 1, 2, 3
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(board[i][j], " . ")
		}
		fmt.Println()

		//fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}
}
