package sudoku

import (
	"testing"
)

// ベンチマーク用のパズルデータ
var (
	// 簡単なパズル
	easyPuzzle = `5 3 4 6 7 8 9 1 .
6 7 2 1 9 5 3 4 8
1 9 8 3 4 2 5 6 7
8 5 9 7 6 1 4 2 3
4 2 6 8 5 3 7 9 1
7 1 3 9 2 4 8 5 6
9 6 1 5 3 7 2 8 4
2 8 7 4 1 9 6 3 5
3 4 5 2 8 6 1 7 .`

	// 中程度のパズル
	mediumPuzzle = `5 3 . . 7 . . . .
6 . . 1 9 5 . . .
. 9 8 . . . . 6 .
8 . . . 6 . . . 3
4 . . 8 . 3 . . 1
7 . . . 2 . . . 6
. 6 . . . . 2 8 .
. . . 4 1 9 . . 5
. . . . 8 . . 7 9`

	// 難しいパズル
	hardPuzzle = `8 . . . . . . . .
. . 3 6 . . . . .
. 7 . . 9 . 2 . .
. 5 . . . 7 . . .
. . . . 4 5 7 . .
. . . 1 . . . 3 .
. . 1 . . . . 6 8
. . 8 5 . . . 1 .
. 9 . . . . 4 . .`

	// 非常に難しいパズル
	expertPuzzle = `. . . . . . . 1 .
4 . . . . 8 . . .
. . . 7 . . 3 . .
. . . . . . . 6 .
. 6 . . 8 . . 4 .
. 5 . . . . . . .
. . 1 . . 6 . . .
. . . 2 . . . . 7
. 8 . . . . . . .`
)

// BenchmarkSolveEasy は簡単なパズルの解法をベンチマーク
func BenchmarkSolveEasy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		board := NewBoard()
		board.LoadFromString(easyPuzzle)
		board.Solve()
	}
}

// BenchmarkSolveMedium は中程度のパズルの解法をベンチマーク
func BenchmarkSolveMedium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		board := NewBoard()
		board.LoadFromString(mediumPuzzle)
		board.Solve()
	}
}

// BenchmarkSolveHard は難しいパズルの解法をベンチマーク
func BenchmarkSolveHard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		board := NewBoard()
		board.LoadFromString(hardPuzzle)
		board.Solve()
	}
}

// BenchmarkSolveExpert は非常に難しいパズルの解法をベンチマーク
func BenchmarkSolveExpert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		board := NewBoard()
		board.LoadFromString(expertPuzzle)
		board.Solve()
	}
}

// BenchmarkSolveWithStrategyEasy はMCVヒューリスティックでの簡単なパズル解法をベンチマーク
func BenchmarkSolveWithStrategyEasy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		board := NewBoard()
		board.LoadFromString(easyPuzzle)
		board.SolveWithStrategy()
	}
}

// BenchmarkSolveWithStrategyMedium はMCVヒューリスティックでの中程度パズル解法をベンチマーク
func BenchmarkSolveWithStrategyMedium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		board := NewBoard()
		board.LoadFromString(mediumPuzzle)
		board.SolveWithStrategy()
	}
}

// BenchmarkSolveWithStrategyHard はMCVヒューリスティックでの難しいパズル解法をベンチマーク
func BenchmarkSolveWithStrategyHard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		board := NewBoard()
		board.LoadFromString(hardPuzzle)
		board.SolveWithStrategy()
	}
}

// BenchmarkSolveWithStrategyExpert はMCVヒューリスティックでの非常に難しいパズル解法をベンチマーク
func BenchmarkSolveWithStrategyExpert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		board := NewBoard()
		board.LoadFromString(expertPuzzle)
		board.SolveWithStrategy()
	}
}

// BenchmarkIsValid は制約チェック機能をベンチマーク
func BenchmarkIsValid(b *testing.B) {
	board := NewBoard()
	board.LoadFromString(mediumPuzzle)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.IsValid(0, 2, 4)
	}
}

// BenchmarkCountPossibilities は可能性カウント機能をベンチマーク
func BenchmarkCountPossibilities(b *testing.B) {
	board := NewBoard()
	board.LoadFromString(mediumPuzzle)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.countPossibilities(0, 2)
	}
}

// BenchmarkFindMostConstrainedCell は最も制約の多いセル検索をベンチマーク
func BenchmarkFindMostConstrainedCell(b *testing.B) {
	board := NewBoard()
	board.LoadFromString(mediumPuzzle)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.findMostConstrainedCell()
	}
}

// BenchmarkLoadFromString は文字列からの読み込みをベンチマーク
func BenchmarkLoadFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		board := NewBoard()
		board.LoadFromString(mediumPuzzle)
	}
}

// BenchmarkFromSimpleString はシンプル文字列からの読み込みをベンチマーク
func BenchmarkFromSimpleString(b *testing.B) {
	simplePuzzle := `53..7....
6..195...
.98....6.
8...6...3
4..8.3..1
7...2...6
.6....28.
...419..5
....8..79`

	for i := 0; i < b.N; i++ {
		board := NewBoard()
		board.FromSimpleString(simplePuzzle)
	}
}

// BenchmarkToSimpleString はシンプル文字列への変換をベンチマーク
func BenchmarkToSimpleString(b *testing.B) {
	board := NewBoard()
	board.LoadFromString(mediumPuzzle)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.ToSimpleString()
	}
}

// BenchmarkSaveToString は文字列への保存をベンチマーク
func BenchmarkSaveToString(b *testing.B) {
	board := NewBoard()
	board.LoadFromString(mediumPuzzle)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.SaveToString()
	}
}

// BenchmarkIsSolved は解答検証をベンチマーク
func BenchmarkIsSolved(b *testing.B) {
	board := NewBoard()
	board.LoadFromString(mediumPuzzle)
	board.Solve() // 解いておく

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.IsSolved()
	}
}

// BenchmarkIsComplete は完成チェックをベンチマーク
func BenchmarkIsComplete(b *testing.B) {
	board := NewBoard()
	board.LoadFromString(mediumPuzzle)
	board.Solve() // 解いておく

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.IsComplete()
	}
}

// BenchmarkSetAndGetCell はセル操作をベンチマーク
func BenchmarkSetAndGetCell(b *testing.B) {
	board := NewBoard()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		row := i % SIZE
		col := (i / SIZE) % SIZE
		value := (i % 9) + 1
		board.SetCell(row, col, value)
		board.GetCell(row, col)
	}
}

// BenchmarkMemoryAllocation は新しい盤面作成のメモリ割り当てをベンチマーク
func BenchmarkMemoryAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		board := NewBoard()
		_ = board // 使用していることを示す
	}
}

// BenchmarkCopyBoard は盤面のコピー操作をベンチマーク
func BenchmarkCopyBoard(b *testing.B) {
	board := NewBoard()
	board.LoadFromString(mediumPuzzle)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copied := board.GetBoard()
		_ = copied // 使用していることを示す
	}
}

// BenchmarkHasUniqueSolution は解の一意性チェックをベンチマーク
func BenchmarkHasUniqueSolution(b *testing.B) {
	board := NewBoard()
	board.LoadFromString(mediumPuzzle)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.HasUniqueSolution()
	}
}

// BenchmarkFullSolveProcess は完全な解法プロセスをベンチマーク
func BenchmarkFullSolveProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		board := NewBoard()
		board.LoadFromString(mediumPuzzle)
		board.Solve()
		board.IsSolved()
		board.ToSimpleString()
	}
}

// BenchmarkMultiplePuzzleSolving は複数のパズルを連続で解くベンチマーク
func BenchmarkMultiplePuzzleSolving(b *testing.B) {
	puzzles := []string{easyPuzzle, mediumPuzzle, hardPuzzle}

	for i := 0; i < b.N; i++ {
		puzzle := puzzles[i%len(puzzles)]
		board := NewBoard()
		board.LoadFromString(puzzle)
		board.Solve()
	}
}

// BenchmarkValidationIntensive は制約チェックを集約的に行うベンチマーク
func BenchmarkValidationIntensive(b *testing.B) {
	board := NewBoard()
	board.LoadFromString(mediumPuzzle)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for row := 0; row < SIZE; row++ {
			for col := 0; col < SIZE; col++ {
				if board.IsEmpty(row, col) {
					for num := 1; num <= 9; num++ {
						board.IsValid(row, col, num)
					}
				}
			}
		}
	}
}
