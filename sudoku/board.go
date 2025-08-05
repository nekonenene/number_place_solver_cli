package sudoku

import "fmt"

const SIZE = 9 // サイズ定数

// Board represents a sudoku board with validation and solving capabilities
type Board struct {
	grid [SIZE][SIZE]int
}

// NewBoard creates a new empty sudoku board
func NewBoard() *Board {
	return &Board{}
}

// SetBoard sets the board with the provided 9x9 grid
func (b *Board) SetBoard(grid [SIZE][SIZE]int) {
	b.grid = grid
}

// GetBoard returns a copy of the current board state
func (b *Board) GetBoard() [SIZE][SIZE]int {
	return b.grid
}

// SetCell sets a value at the specified position
func (b *Board) SetCell(row, col, value int) bool {
	if row < 0 || row >= SIZE || col < 0 || col >= SIZE {
		return false
	}
	if value < 0 || value > 9 {
		return false
	}
	b.grid[row][col] = value
	return true
}

// GetCell returns the value at the specified position
func (b *Board) GetCell(row, col int) int {
	if row < 0 || row >= SIZE || col < 0 || col >= SIZE {
		return -1
	}
	return b.grid[row][col]
}

// IsEmpty checks if a cell is empty (contains 0)
func (b *Board) IsEmpty(row, col int) bool {
	return b.GetCell(row, col) == 0
}

// Print displays the current board state in a formatted way
func (b *Board) Print() {
	for i := 0; i < SIZE; i++ {
		if i%3 == 0 {
			fmt.Println("+-------+-------+-------+")
		}
		for j := 0; j < SIZE; j++ {
			if j%3 == 0 {
				fmt.Print("| ")
			}
			if b.grid[i][j] == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", b.grid[i][j])
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+-------+-------+-------+")
}

// IsValid checks if placing a number at the given position is valid
func (b *Board) IsValid(row, col, num int) bool {
	// Check row constraint
	for j := 0; j < SIZE; j++ {
		if b.grid[row][j] == num {
			return false
		}
	}

	// Check column constraint
	for i := 0; i < SIZE; i++ {
		if b.grid[i][col] == num {
			return false
		}
	}

	// Check 3x3 box constraint
	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3
	for i := boxRow; i < boxRow+3; i++ {
		for j := boxCol; j < boxCol+3; j++ {
			if b.grid[i][j] == num {
				return false
			}
		}
	}

	return true
}

// IsComplete checks if the board is completely filled
func (b *Board) IsComplete() bool {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if b.grid[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

// IsSolved checks if the board is both complete and valid
func (b *Board) IsSolved() bool {
	if !b.IsComplete() {
		return false
	}

	// Check all constraints
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			num := b.grid[i][j]
			b.grid[i][j] = 0 // Temporarily remove to check validity
			if !b.IsValid(i, j, num) {
				b.grid[i][j] = num // Restore
				return false
			}
			b.grid[i][j] = num // Restore
		}
	}
	return true
}
