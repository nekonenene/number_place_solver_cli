package sudoku

import (
	"strings"
	"testing"
)

// TestNewBoard は盤面作成をテスト
func TestNewBoard(t *testing.T) {
	board := NewBoard()
	if board == nil {
		t.Fatal("NewBoard() returned nil")
	}

	// 盤面が空かどうかをチェック
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if board.GetCell(i, j) != 0 {
				t.Errorf("New board should be empty, but cell (%d, %d) = %d", i, j, board.GetCell(i, j))
			}
		}
	}
}

// TestSetAndGetCell はセル操作をテスト
func TestSetAndGetCell(t *testing.T) {
	board := NewBoard()

	// 有効な操作をテスト
	if !board.SetCell(0, 0, 5) {
		t.Error("SetCell(0, 0, 5) should succeed")
	}
	if board.GetCell(0, 0) != 5 {
		t.Errorf("GetCell(0, 0) = %d, want 5", board.GetCell(0, 0))
	}

	// 境界条件をテスト
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

	// 無効な値をテスト
	if board.SetCell(0, 1, -1) {
		t.Error("SetCell(0, 1, -1) should fail")
	}
	if board.SetCell(0, 1, 10) {
		t.Error("SetCell(0, 1, 10) should fail")
	}

	// 境界取得操作をテスト
	if board.GetCell(-1, 0) != -1 {
		t.Error("GetCell(-1, 0) should return -1")
	}
	if board.GetCell(0, -1) != -1 {
		t.Error("GetCell(0, -1) should return -1")
	}
}

// TestIsValid は検証ロジックをテスト
func TestIsValid(t *testing.T) {
	board := NewBoard()

	// テスト設定をセットアップ
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

	// 行の制約をテスト
	if board.IsValid(0, 2, 5) { // 5 already exists in row 0
		t.Error("Should detect row constraint violation")
	}

	// 列の制約をテスト
	if board.IsValid(2, 0, 5) { // 5 already exists in column 0
		t.Error("Should detect column constraint violation")
	}

	// ボックスの制約をテスト
	if board.IsValid(1, 1, 3) { // 3 already exists in the same 3x3 box
		t.Error("Should detect box constraint violation")
	}

	// 有効な配置をテスト
	if !board.IsValid(0, 2, 4) {
		t.Error("Should allow valid placement")
	}
}

// TestLoadFromString は文字列入力の解析をテスト
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

	// 特定の値をチェック
	if board.GetCell(0, 0) != 5 {
		t.Errorf("Expected 5 at (0,0), got %d", board.GetCell(0, 0))
	}
	if board.GetCell(0, 2) != 0 {
		t.Errorf("Expected 0 at (0,2), got %d", board.GetCell(0, 2))
	}

	// 無効な入力をテスト
	invalidInput := "invalid input"
	err = board.LoadFromString(invalidInput)
	if err == nil {
		t.Error("LoadFromString should fail with invalid input")
	}
}

// TestSolve は解法アルゴリズムをテスト
func TestSolve(t *testing.T) {
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
		t.Fatalf("Failed to load puzzle: %v", err)
	}

	if !board.Solve() {
		t.Fatal("Failed to solve the puzzle")
	}

	if !board.IsSolved() {
		t.Error("Board should be solved but validation failed")
	}
}

// TestToSimpleString はシンプル文字列変換をテスト
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

// TestFromSimpleString はシンプル文字列解析をテスト
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

// TestLoadFromStringWithSpaces はスペース区切り入力をテスト
func TestLoadFromStringWithSpaces(t *testing.T) {
	board := NewBoard()

	// スペース区切り形式でのテスト
	spaceInput := `5 3 . . 7 . . . .
6 . . 1 9 5 . . .
. 9 8 . . . . 6 .
8 . . . 6 . . . 3
4 . . 8 . 3 . . 1
7 . . . 2 . . . 6
. 6 . . . . 2 8 .
. . . 4 1 9 . . 5
. . . . 8 . . 7 9`

	err := board.LoadFromString(spaceInput)
	if err != nil {
		t.Fatalf("LoadFromString with spaces failed: %v", err)
	}

	// 検証
	if board.GetCell(0, 0) != 5 {
		t.Errorf("Expected 5 at (0,0), got %d", board.GetCell(0, 0))
	}
	if board.GetCell(0, 2) != 0 {
		t.Errorf("Expected 0 at (0,2), got %d", board.GetCell(0, 2))
	}
}

// TestLoadFromStringWithMixedSpaces は半角スペースを空欄として使用する入力をテスト
func TestLoadFromStringWithMixedSpaces(t *testing.T) {
	board := NewBoard()

	// 半角スペースを空欄として使用（9文字固定）
	mixedInput := `53  7....
6  195...
 98    6.
8   6   3
4  8 3  1
7   2   6
 6    28.
   419  5
    8  79`

	err := board.LoadFromString(mixedInput)
	if err != nil {
		t.Fatalf("LoadFromString with mixed spaces failed: %v", err)
	}

	// 検証
	if board.GetCell(0, 0) != 5 {
		t.Errorf("Expected 5 at (0,0), got %d", board.GetCell(0, 0))
	}
	if board.GetCell(0, 2) != 0 {
		t.Errorf("Expected 0 at (0,2), got %d", board.GetCell(0, 2))
	}
	if board.GetCell(1, 1) != 0 {
		t.Errorf("Expected 0 at (1,1), got %d", board.GetCell(1, 1))
	}
}

// TestBoardCopy は盤面のコピー機能をテスト
func TestBoardCopy(t *testing.T) {
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

	// コピーを取得
	copied := board.GetBoard()

	// コピーを変更
	copied[0][0] = 9

	// 元の盤面が変更されていないことを確認
	if board.GetCell(0, 0) != 5 {
		t.Error("元の盤面がコピーの変更により影響を受けました")
	}
}

// TestIsEmptyEdgeCases はIsEmptyのエッジケースをテスト
func TestIsEmptyEdgeCases(t *testing.T) {
	board := NewBoard()

	// 境界外のセルをテスト
	if board.IsEmpty(-1, 0) {
		t.Error("境界外のセルに対してtrueを返すべきではありません")
	}
	if board.IsEmpty(0, -1) {
		t.Error("境界外のセルに対してtrueを返すべきではありません")
	}
	if board.IsEmpty(SIZE, 0) {
		t.Error("境界外のセルに対してtrueを返すべきではありません")
	}
	if board.IsEmpty(0, SIZE) {
		t.Error("境界外のセルに対してtrueを返すべきではありません")
	}

	// 正常なセルをテスト
	if !board.IsEmpty(0, 0) {
		t.Error("新しい盤面のセルは空であることを返すべきです")
	}

	// セルに値を設定してテスト
	board.SetCell(0, 0, 5)
	if board.IsEmpty(0, 0) {
		t.Error("値が設定されたセルは空ではありません")
	}
}

// TestIsCompleteEdgeCases はIsCompleteのエッジケースをテスト
func TestIsCompleteEdgeCases(t *testing.T) {
	board := NewBoard()

	// 空の盤面は完成していない
	if board.IsComplete() {
		t.Error("空の盤面は完成していません")
	}

	// 部分的に埋められた盤面
	for i := 0; i < SIZE-1; i++ {
		for j := 0; j < SIZE; j++ {
			board.SetCell(i, j, 1)
		}
	}

	if board.IsComplete() {
		t.Error("部分的に埋められた盤面は完成していません")
	}

	// 最後の行の残りのセルを埋める
	for j := 0; j < SIZE-1; j++ {
		board.SetCell(SIZE-1, j, 1)
	}
	board.SetCell(SIZE-1, SIZE-1, 1)

	if !board.IsComplete() {
		t.Error("全てのセルが埋められた盤面は完成しています")
	}
}

// TestIsSolvedInvalidBoard は無効な盤面でのIsSolvedをテスト
func TestIsSolvedInvalidBoard(t *testing.T) {
	board := NewBoard()

	// 無効な盤面を作成（同じ行に同じ数字）
	invalidGrid := [SIZE][SIZE]int{
		{1, 1, 3, 4, 5, 6, 7, 8, 9}, // 同じ行に1が2個
		{2, 3, 4, 5, 6, 7, 8, 9, 1},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{5, 6, 7, 8, 9, 1, 2, 3, 4},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{9, 1, 2, 3, 4, 5, 6, 7, 8},
	}
	board.SetBoard(invalidGrid)

	if board.IsSolved() {
		t.Error("無効な盤面は解決されていません")
	}
}

// TestGetStatsInitialState は初期状態での統計情報をテスト
func TestGetStatsInitialState(t *testing.T) {
	board := NewBoard()

	stats := board.GetStats()
	if stats.BacktrackCount != 0 {
		t.Errorf("初期バックトラック数は0であるべきですが、%d でした", stats.BacktrackCount)
	}
	if stats.CellsSet != 0 {
		t.Errorf("初期セット数は0であるべきですが、%d でした", stats.CellsSet)
	}
}

// TestLargeNumberValidation は大きな数値の検証をテスト
func TestLargeNumberValidation(t *testing.T) {
	board := NewBoard()

	// 10以上の値を設定しようとする
	if board.SetCell(0, 0, 10) {
		t.Error("10は無効な値ですが設定できてしまいました")
	}

	// 負の値を設定しようとする
	if board.SetCell(0, 0, -1) {
		t.Error("-1は無効な値ですが設定できてしまいました")
	}

	// 境界値をテスト
	if !board.SetCell(0, 0, 0) {
		t.Error("0は有効な値です")
	}
	if !board.SetCell(0, 1, 9) {
		t.Error("9は有効な値です")
	}
}

// TestComplexConstraintValidation は複雑な制約検証をテスト
func TestComplexConstraintValidation(t *testing.T) {
	board := NewBoard()

	// 知られている有効なパズルを使用
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
		t.Fatalf("有効なパズルの読み込みに失敗: %v", err)
	}

	// セル(0,2)に対する制約チェック
	// この位置に4は配置可能であるべき
	if !board.IsValid(0, 2, 4) {
		t.Error("セル(0,2)に4を配置できるはずです")
	}

	// この位置に5は配置不可能であるべき（同じ行に既に5がある）
	if board.IsValid(0, 2, 5) {
		t.Error("セル(0,2)に5は配置できません（行制約）")
	}

	// 可能な数字をカウント
	validCount := 0
	for num := 1; num <= 9; num++ {
		if board.IsValid(0, 2, num) {
			validCount++
		}
	}

	if validCount == 0 {
		t.Error("セル(0,2)に配置可能な数字が見つかりません")
	} else {
		t.Logf("セル(0,2)に配置可能な数字の数: %d", validCount)
	}
}

// TestStressTestBoardOperations は盤面操作のストレステスト
func TestStressTestBoardOperations(t *testing.T) {
	board := NewBoard()

	// 大量のセット/ゲット操作
	for iteration := 0; iteration < 100; iteration++ {
		for i := 0; i < SIZE; i++ {
			for j := 0; j < SIZE; j++ {
				value := ((i*SIZE + j + iteration) % 9) + 1
				if !board.SetCell(i, j, value) {
					t.Fatalf("SetCell失敗 (%d, %d, %d) at iteration %d", i, j, value, iteration)
				}
				if board.GetCell(i, j) != value {
					t.Fatalf("GetCell不一致 (%d, %d): 期待値=%d, 実際=%d", i, j, value, board.GetCell(i, j))
				}
			}
		}

		// 盤面をクリア
		for i := 0; i < SIZE; i++ {
			for j := 0; j < SIZE; j++ {
				board.SetCell(i, j, 0)
			}
		}
	}
}

// TestIsValidPuzzle は初期盤面の有効性チェックをテスト
func TestIsValidPuzzle(t *testing.T) {
	// 有効なパズル
	board := NewBoard()
	validGrid := [SIZE][SIZE]int{
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
	board.SetBoard(validGrid)

	if !board.IsValidPuzzle() {
		t.Error("有効なパズルが無効と判定されました")
	}

	// 無効なパズル（行に重複）
	board2 := NewBoard()
	invalidGrid := [SIZE][SIZE]int{
		{1, 1, 0, 0, 0, 0, 0, 0, 0}, // 同じ行に1が2個
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	board2.SetBoard(invalidGrid)

	if board2.IsValidPuzzle() {
		t.Error("無効なパズル（行重複）が有効と判定されました")
	}

	// 無効なパズル（列に重複）
	board3 := NewBoard()
	invalidGrid2 := [SIZE][SIZE]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0}, // 同じ列に1が2個
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	board3.SetBoard(invalidGrid2)

	if board3.IsValidPuzzle() {
		t.Error("無効なパズル（列重複）が有効と判定されました")
	}

	// 無効なパズル（3x3ボックス内に重複）
	board4 := NewBoard()
	invalidGrid3 := [SIZE][SIZE]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0}, // 同じボックス内に1が2個
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	board4.SetBoard(invalidGrid3)

	if board4.IsValidPuzzle() {
		t.Error("無効なパズル（ボックス重複）が有効と判定されました")
	}
}
