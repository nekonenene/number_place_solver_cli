package number_place

import (
	"testing"
	"time"
)

// TestSolveWithStrategy はMCVヒューリスティックを使用した解法をテスト
func TestSolveWithStrategy(t *testing.T) {
	board := NewBoard()

	// 既知の解可能パズルを読み込み
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
		t.Fatalf("パズルの読み込みに失敗: %v", err)
	}

	start := time.Now()
	if !board.SolveWithStrategy() {
		t.Fatal("MCVヒューリスティックでの解法に失敗")
	}
	elapsed := time.Since(start)

	if !board.IsSolved() {
		t.Error("盤面は解かれているはずですが検証に失敗")
	}

	t.Logf("MCV解法時間: %v", elapsed)
}

// TestSolveComparison は通常の解法とMCVヒューリスティックの比較テスト
func TestSolveComparison(t *testing.T) {
	if testing.Short() {
		t.Skip("時間のかかるテストのためshortモードでスキップ")
	}

	// より難しいパズルを使用
	hardPuzzle := `. . . . . . . 1 .
4 . . . . 8 . . .
. . . 7 . . 3 . .
. . . . . . . 6 .
. 6 . . 8 . . 4 .
. 5 . . . . . . .
. . 1 . . 6 . . .
. . . 2 . . . . 7
. 8 . . . . . . .`

	// 通常の解法
	board1 := NewBoard()
	err := board1.LoadFromString(hardPuzzle)
	if err != nil {
		t.Fatalf("パズルの読み込みに失敗: %v", err)
	}

	start1 := time.Now()
	solved1 := board1.Solve()
	elapsed1 := time.Since(start1)
	stats1 := board1.GetStats()

	// MCV解法
	board2 := NewBoard()
	err = board2.LoadFromString(hardPuzzle)
	if err != nil {
		t.Fatalf("パズルの読み込みに失敗: %v", err)
	}

	start2 := time.Now()
	solved2 := board2.SolveWithStrategy()
	elapsed2 := time.Since(start2)

	if !solved1 || !solved2 {
		t.Fatal("両方の解法が成功する必要があります")
	}

	t.Logf("通常解法: 時間=%v, セット数=%d, バックトラック数=%d",
		elapsed1, stats1.CellsSet, stats1.BacktrackCount)
	t.Logf("MCV解法: 時間=%v", elapsed2)

	// 両方が正しく解けていることを確認
	if !board1.IsSolved() || !board2.IsSolved() {
		t.Error("両方の解法で正しい解が得られませんでした")
	}

	// 注意：ナンバープレースは複数の解を持つことがあるため、必ずしも同じ解になるとは限らない
	// しかし、両方とも有効な解であることは確認できる
}

// TestHasUniqueSolution は解の一意性チェック機能をテスト
func TestHasUniqueSolution(t *testing.T) {
	if testing.Short() {
		t.Skip("時間のかかるテストのためshortモードでスキップ")
	}

	board := NewBoard()

	// より制約の多い（解が少ない）パズルを使用
	almostSolvedPuzzle := `5 3 4 6 7 8 9 1 2
6 7 2 1 9 5 3 4 8
1 9 8 3 4 2 5 6 7
8 5 9 7 6 1 4 2 3
4 2 6 8 5 3 7 9 1
7 1 3 9 2 4 8 5 6
9 6 1 5 3 7 2 8 4
2 8 7 4 1 9 6 3 5
3 4 5 2 8 6 1 7 .`

	err := board.LoadFromString(almostSolvedPuzzle)
	if err != nil {
		t.Fatalf("パズルの読み込みに失敗: %v", err)
	}

	// ほぼ完成されたパズルは一意解を持つ
	if !board.HasUniqueSolution() {
		t.Error("ほぼ完成されたパズルは一意解を持つはずです")
	}

	// 複数解を持つパズル（意図的に制約を少なくする）
	board2 := NewBoard()
	multiplePuzzle := `1 . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .`

	err = board2.LoadFromString(multiplePuzzle)
	if err != nil {
		t.Fatalf("パズルの読み込みに失敗: %v", err)
	}

	// 空のセルが多いパズルは複数解を持つ（早期終了により）
	if board2.HasUniqueSolution() {
		t.Error("制約の少ないパズルは複数解を持つはずです")
	}
}

// TestCountSolutions は解の数を数える機能をテスト
func TestCountSolutions(t *testing.T) {
	if testing.Short() {
		t.Skip("時間のかかるテストのためshortモードでスキップ")
	}

	board := NewBoard()

	// ほぼ完成されたパズル（計算が早い）
	almostComplete := `5 3 4 6 7 8 9 1 2
6 7 2 1 9 5 3 4 8
1 9 8 3 4 2 5 6 7
8 5 9 7 6 1 4 2 3
4 2 6 8 5 3 7 9 1
7 1 3 9 2 4 8 5 6
9 6 1 5 3 7 2 8 4
2 8 7 4 1 9 6 3 5
3 4 5 2 8 6 1 . .`

	err := board.LoadFromString(almostComplete)
	if err != nil {
		t.Fatalf("パズルの読み込みに失敗: %v", err)
	}

	// 最大3つの解まで数える（計算量を制限）
	solutions := board.countSolutions(3)
	if solutions == 0 {
		t.Error("少なくとも1つの解が存在するはずです")
	}

	t.Logf("見つかった解の数: %d", solutions)
}

// TestSolveStats は解法統計情報の記録をテスト
func TestSolveStats(t *testing.T) {
	board := NewBoard()

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
		t.Fatalf("パズルの読み込みに失敗: %v", err)
	}

	if !board.Solve() {
		t.Fatal("パズルの解法に失敗")
	}

	stats := board.GetStats()

	if stats.CellsSet <= 0 {
		t.Error("セット数は0より大きくなければなりません")
	}

	if stats.BacktrackCount < 0 {
		t.Error("バックトラック数は負の値になってはいけません")
	}

	t.Logf("統計情報 - セット数: %d, バックトラック数: %d",
		stats.CellsSet, stats.BacktrackCount)
}

// TestFindMostConstrainedCell はMCV選択の動作をテスト
func TestFindMostConstrainedCell(t *testing.T) {
	board := NewBoard()

	// 制約の異なるセルを持つパズル
	puzzle := `5 3 4 . 7 8 9 1 2
6 7 2 1 9 5 3 4 8
1 9 8 3 4 2 5 6 7
8 5 9 7 6 1 4 2 3
4 2 6 8 5 3 7 9 1
7 1 3 9 2 4 8 5 6
9 6 1 5 3 7 2 8 4
2 8 7 4 1 9 6 3 5
3 4 5 2 8 6 1 7 .`

	err := board.LoadFromString(puzzle)
	if err != nil {
		t.Fatalf("パズルの読み込みに失敗: %v", err)
	}

	row, col := board.findMostConstrainedCell()

	// 最も制約されたセルを見つけたことを確認
	if row == -1 || col == -1 {
		t.Error("空のセルが見つからなかった")
		return
	}

	// 見つかったセルの可能性の数をテスト
	possibilities := board.countPossibilities(row, col)
	t.Logf("最も制約されたセル (%d,%d) の可能性の数: %d", row, col, possibilities)

	// 他のセルよりも制約されていることを確認
	found := false
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if board.IsEmpty(i, j) && (i != row || j != col) {
				otherPossibilities := board.countPossibilities(i, j)
				if otherPossibilities < possibilities {
					t.Errorf("セル (%d,%d) の可能性数 %d がセル (%d,%d) の可能性数 %d より少ない",
						i, j, otherPossibilities, row, col, possibilities)
				}
				found = true
			}
		}
	}

	if !found {
		// 唯一の空のセルだった
		if possibilities <= 0 || possibilities > 9 {
			t.Errorf("可能性の数が無効: %d", possibilities)
		}
	}
}

// TestCountPossibilities は可能性カウント機能をテスト
func TestCountPossibilities(t *testing.T) {
	board := NewBoard()

	// 特定の制約パターンを持つパズル
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
		t.Fatalf("パズルの読み込みに失敗: %v", err)
	}

	// セル (0,2) の可能性をテスト
	count := board.countPossibilities(0, 2)
	if count <= 0 || count > 9 {
		t.Errorf("可能性の数は1-9の範囲にあるはずですが、%d でした", count)
	}

	// 具体的に確認
	validNums := make([]int, 0)
	for num := 1; num <= 9; num++ {
		if board.IsValid(0, 2, num) {
			validNums = append(validNums, num)
		}
	}

	if len(validNums) != count {
		t.Errorf("IsValidで確認した可能性の数 %d とcountPossibilitiesの結果 %d が一致しません",
			len(validNums), count)
	}

	t.Logf("セル (0,2) の可能な数字: %v (計%d個)", validNums, count)
}

// TestUnsolvablePuzzle は解けないパズルのテスト
func TestUnsolvablePuzzle(t *testing.T) {
	board := NewBoard()

	// 明らかに解けないパズル（同じ行に同じ数字）
	unsolvablePuzzle := `1 1 . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .`

	// 無効なパズルは読み込み時点でエラーになるはず
	err := board.LoadFromString(unsolvablePuzzle)
	if err == nil {
		t.Error("無効なパズルは読み込み時点でエラーになるべきです")
	} else {
		t.Logf("期待通りエラーが発生: %v", err)
	}

	// 別の無効パズル（同じ列に同じ数字）
	board2 := NewBoard()
	invalidColumnPuzzle := `1 . . . . . . . .
1 . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .`

	err2 := board2.LoadFromString(invalidColumnPuzzle)
	if err2 == nil {
		t.Error("無効なパズル（列重複）は読み込み時点でエラーになるべきです")
	} else {
		t.Logf("期待通りエラーが発生: %v", err2)
	}

	// 3x3ボックス内の重複
	board3 := NewBoard()
	invalidBoxPuzzle := `1 . . . . . . . .
. 1 . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .`

	err3 := board3.LoadFromString(invalidBoxPuzzle)
	if err3 == nil {
		t.Error("無効なパズル（ボックス重複）は読み込み時点でエラーになるべきです")
	} else {
		t.Logf("期待通りエラーが発生: %v", err3)
	}
}

// TestComplexPuzzle は複雑なパズルの解法テスト
func TestComplexPuzzle(t *testing.T) {
	if testing.Short() {
		t.Skip("時間のかかるテストのためshortモードでスキップ")
	}

	board := NewBoard()

	// より複雑で難しいパズル
	complexPuzzle := `8 . . . . . . . .
. . 3 6 . . . . .
. 7 . . 9 . 2 . .
. 5 . . . 7 . . .
. . . . 4 5 7 . .
. . . 1 . . . 3 .
. . 1 . . . . 6 8
. . 8 5 . . . 1 .
. 9 . . . . 4 . .`

	err := board.LoadFromString(complexPuzzle)
	if err != nil {
		t.Fatalf("パズルの読み込みに失敗: %v", err)
	}

	start := time.Now()
	solved := board.Solve()
	elapsed := time.Since(start)

	if !solved {
		t.Fatal("複雑なパズルの解法に失敗")
	}

	if !board.IsSolved() {
		t.Error("解は正しくありません")
	}

	stats := board.GetStats()
	t.Logf("複雑なパズル - 時間: %v, セット数: %d, バックトラック数: %d",
		elapsed, stats.CellsSet, stats.BacktrackCount)
}
