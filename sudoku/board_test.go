package sudoku

import (
	"strings"
	"testing"
)

// TestNewBoard tests board creation
func TestNewBoard(t *testing.T) {
	board := NewBoard()
	if board == nil {
		t.Fatal("NewBoard() returned nil")
	}

	// Check if board is empty
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if board.GetCell(i, j) != 0 {
				t.Errorf("New board should be empty, but cell (%d, %d) = %d", i, j, board.GetCell(i, j))
			}
		}
	}
}

// TestSetAndGetCell tests cell manipulation
func TestSetAndGetCell(t *testing.T) {
	board := NewBoard()

	// Test valid operations
	if !board.SetCell(0, 0, 5) {
		t.Error("SetCell(0, 0, 5) should succeed")
	}
	if board.GetCell(0, 0) != 5 {
		t.Errorf("GetCell(0, 0) = %d, want 5", board.GetCell(0, 0))
	}

	// Test boundary conditions
	if board.SetCell(-1, 0, 5) {
		t.Error("SetCell(-1, 0, 5) should fail")
	}
	if board.SetCell(0, -1, 5) {
		t.Error("SetCell(0, -1, 5) should fail")
	}
	if board.SetCell(SIZE, 0, 5) {
		t.Error("SetCell(SIZE, 0, 5) should fail")
	}
	if board.SetCell(0, SIZE, 5) {
		t.Error("SetCell(0, SIZE, 5) should fail")
	}

	// Test invalid values
	if board.SetCell(0, 1, -1) {
		t.Error("SetCell(0, 1, -1) should fail")
	}
	if board.SetCell(0, 1, 10) {
		t.Error("SetCell(0, 1, 10) should fail")
	}

	// Test boundary get operations
	if board.GetCell(-1, 0) != -1 {
		t.Error("GetCell(-1, 0) should return -1")
	}
	if board.GetCell(0, -1) != -1 {
		t.Error("GetCell(0, -1) should return -1")
	}
}

// TestIsValid tests the validation logic
func TestIsValid(t *testing.T) {
	board := NewBoard()

	// Set up a test configuration
	testGrid := [SIZE][SIZE]int{
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
	board.SetBoard(testGrid)

	// Test row constraint
	if board.IsValid(0, 2, 5) { // 5 already exists in row 0
		t.Error("Should detect row constraint violation")
	}

	// Test column constraint
	if board.IsValid(2, 0, 5) { // 5 already exists in column 0
		t.Error("Should detect column constraint violation")
	}

	// Test box constraint
	if board.IsValid(1, 1, 3) { // 3 already exists in the same 3x3 box
		t.Error("Should detect box constraint violation")
	}

	// Test valid placement
	if !board.IsValid(0, 2, 4) {
		t.Error("Should allow valid placement")
	}
}

// TestLoadFromString tests string input parsing
func TestLoadFromString(t *testing.T) {
	board := NewBoard()

	validInput := `5 3 . . 7 . . . .
6 . . 1 9 5 . . .
. 9 8 . . . . 6 .
8 . . . 6 . . . 3
4 . . 8 . 3 . . 1
7 . . . 2 . . . 6
. 6 . . . . 2 8 .
. . . 4 1 9 . . 5
. . . . 8 . . 7 9`

	err := board.LoadFromString(validInput)
	if err != nil {
		t.Fatalf("LoadFromString failed: %v", err)
	}

	// Check specific values
	if board.GetCell(0, 0) != 5 {
		t.Errorf("Expected 5 at (0,0), got %d", board.GetCell(0, 0))
	}
	if board.GetCell(0, 2) != 0 {
		t.Errorf("Expected 0 at (0,2), got %d", board.GetCell(0, 2))
	}

	// Test invalid input
	invalidInput := "invalid input"
	err = board.LoadFromString(invalidInput)
	if err == nil {
		t.Error("LoadFromString should fail with invalid input")
	}
}

// TestSolve tests the solving algorithm
func TestSolve(t *testing.T) {
	board := NewBoard()

	// Load a known solvable puzzle
	puzzle := `5 3 . . 7 . . . .
6 . . 1 9 5 . . .
. 9 8 . . . . 6 .
8 . . . 6 . . . 3
4 . . 8 . 3 . . 1
7 . . . 2 . . . 6
. 6 . . . . 2 8 .
. . . 4 1 9 . . 5
. . . . 8 . . 7 9`

	err := board.LoadFromString(puzzle)
	if err != nil {
		t.Fatalf("Failed to load puzzle: %v", err)
	}

	if !board.Solve() {
		t.Fatal("Failed to solve the puzzle")
	}

	if !board.IsSolved() {
		t.Error("Board should be solved but validation failed")
	}
}

// TestToSimpleString tests simple string conversion
func TestToSimpleString(t *testing.T) {
	board := NewBoard()
	testGrid := [SIZE][SIZE]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	board.SetBoard(testGrid)

	result := board.ToSimpleString()
	lines := strings.Split(result, "\n")

	if len(lines) != SIZE {
		t.Errorf("Expected %d lines, got %d", SIZE, len(lines))
	}

	for i, line := range lines {
		if len(line) != SIZE {
			t.Errorf("Line %d has %d characters, expected %d", i, len(line), SIZE)
		}
	}
}

// TestFromSimpleString tests simple string parsing
func TestFromSimpleString(t *testing.T) {
	board := NewBoard()
	input := `53..7....
6..195...
.98....6.
8...6...3
4..8.3..1
7...2...6
.6....28.
...419..5
....8..79`

	err := board.FromSimpleString(input)
	if err != nil {
		t.Fatalf("FromSimpleString failed: %v", err)
	}

	if board.GetCell(0, 0) != 5 {
		t.Errorf("Expected 5 at (0,0), got %d", board.GetCell(0, 0))
	}
	if board.GetCell(0, 2) != 0 {
		t.Errorf("Expected 0 at (0,2), got %d", board.GetCell(0, 2))
	}
}
