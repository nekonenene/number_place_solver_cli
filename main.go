package main

import (
	"fmt"
	"os"
	"sudoku_solver/sudoku"
)

// main serves as the entry point of the sudoku solver application
func main() {
	fmt.Println("Sudoku Solver")
	fmt.Println("=============")

	// Create a sample sudoku puzzle for testing
	puzzle := sudoku.NewBoard()

	// Load a test puzzle
	testPuzzle := [9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	puzzle.SetBoard(testPuzzle)

	fmt.Println("Original puzzle:")
	puzzle.Print()

	// Solve the puzzle
	if puzzle.Solve() {
		fmt.Println("\nSolved puzzle:")
		puzzle.Print()
	} else {
		fmt.Println("No solution found!")
		os.Exit(1)
	}
}
