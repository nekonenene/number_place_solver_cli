package sudoku

// Solve attempts to solve the sudoku puzzle using backtracking algorithm
func (b *Board) Solve() bool {
	return b.solveRecursive()
}

// solveRecursive implements the backtracking algorithm recursively
func (b *Board) solveRecursive() bool {
	// Find the first empty cell
	row, col := b.findEmptyCell()
	if row == -1 {
		// No empty cell found, puzzle is solved
		return true
	}

	// Try digits 1-9
	for num := 1; num <= 9; num++ {
		if b.IsValid(row, col, num) {
			// Place the number
			b.grid[row][col] = num

			// Recursively solve the rest
			if b.solveRecursive() {
				return true
			}

			// Backtrack: remove the number if it doesn't lead to a solution
			b.grid[row][col] = 0
		}
	}

	// No valid number found for this cell
	return false
}

// findEmptyCell finds the first empty cell (containing 0) in the board
func (b *Board) findEmptyCell() (int, int) {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if b.grid[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1 // No empty cell found
}

// SolveWithStrategy solves using the most constrained variable heuristic
func (b *Board) SolveWithStrategy() bool {
	return b.solveWithMCV()
}

// solveWithMCV implements backtracking with Most Constrained Variable heuristic
func (b *Board) solveWithMCV() bool {
	// Find the empty cell with the fewest possibilities
	row, col := b.findMostConstrainedCell()
	if row == -1 {
		// No empty cell found, puzzle is solved
		return true
	}

	// Try digits 1-9
	for num := 1; num <= 9; num++ {
		if b.IsValid(row, col, num) {
			// Place the number
			b.grid[row][col] = num

			// Recursively solve the rest
			if b.solveWithMCV() {
				return true
			}

			// Backtrack: remove the number if it doesn't lead to a solution
			b.grid[row][col] = 0
		}
	}

	// No valid number found for this cell
	return false
}

// findMostConstrainedCell finds the empty cell with the fewest valid possibilities
func (b *Board) findMostConstrainedCell() (int, int) {
	bestRow, bestCol := -1, -1
	minPossibilities := 10 // Maximum possibilities is 9

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if b.grid[i][j] == 0 {
				possibilities := b.countPossibilities(i, j)
				if possibilities < minPossibilities {
					minPossibilities = possibilities
					bestRow, bestCol = i, j
				}
			}
		}
	}

	return bestRow, bestCol
}

// countPossibilities counts how many valid numbers can be placed in a cell
func (b *Board) countPossibilities(row, col int) int {
	count := 0
	for num := 1; num <= 9; num++ {
		if b.IsValid(row, col, num) {
			count++
		}
	}
	return count
}

// HasUniqueSolution checks if the puzzle has exactly one solution
func (b *Board) HasUniqueSolution() bool {
	solutions := b.countSolutions(0)
	return solutions == 1
}

// countSolutions counts the number of possible solutions (up to maxSolutions)
func (b *Board) countSolutions(maxSolutions int) int {
	// Make a copy of the board
	originalGrid := b.grid
	defer func() { b.grid = originalGrid }()

	return b.countSolutionsRecursive(maxSolutions)
}

// countSolutionsRecursive recursively counts solutions
func (b *Board) countSolutionsRecursive(maxSolutions int) int {
	// Find the first empty cell
	row, col := b.findEmptyCell()
	if row == -1 {
		// No empty cell found, found one solution
		return 1
	}

	solutionCount := 0
	// Try digits 1-9
	for num := 1; num <= 9; num++ {
		if b.IsValid(row, col, num) {
			// Place the number
			b.grid[row][col] = num

			// Recursively count solutions
			solutionCount += b.countSolutionsRecursive(maxSolutions)

			// Early exit if we've found enough solutions
			if maxSolutions > 0 && solutionCount >= maxSolutions {
				b.grid[row][col] = 0
				return solutionCount
			}

			// Backtrack
			b.grid[row][col] = 0
		}
	}

	return solutionCount
}
