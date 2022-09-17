package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_", "_"}, // i = 0
		[]string{"_", "_", "_", "_"}, // i = 1
		[]string{"_", "_", "_", "_"}, // i = 2
		[]string{"_", "_", "_", "_"}, // i = 3
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "="
	board[1][2] = "2"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ { //i = 0, 1, 2, 3
		fmt.Println(strings.Join(board[i], ".."))

		//fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}
}
