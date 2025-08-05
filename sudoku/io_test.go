package sudoku

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// TestLoadFromReader はio.Readerからの読み込みをテスト
func TestLoadFromReader(t *testing.T) {
	board := NewBoard()

	input := `5 3 . . 7 . . . .
6 . . 1 9 5 . . .
. 9 8 . . . . 6 .
8 . . . 6 . . . 3
4 . . 8 . 3 . . 1
7 . . . 2 . . . 6
. 6 . . . . 2 8 .
. . . 4 1 9 . . 5
. . . . 8 . . 7 9`

	reader := strings.NewReader(input)
	err := board.LoadFromReader(reader)
	if err != nil {
		t.Fatalf("Readerからの読み込みに失敗: %v", err)
	}

	// 特定の値をチェック
	if board.GetCell(0, 0) != 5 {
		t.Errorf("セル(0,0)は5であるべきですが、%d でした", board.GetCell(0, 0))
	}
	if board.GetCell(0, 2) != 0 {
		t.Errorf("セル(0,2)は0であるべきですが、%d でした", board.GetCell(0, 2))
	}
}

// TestLoadFromReaderWithComments はコメント行を含む入力のテスト
func TestLoadFromReaderWithComments(t *testing.T) {
	board := NewBoard()

	inputWithComments := `# これはコメント行です
5 3 . . 7 . . . .
# 2行目のコメント
6 . . 1 9 5 . . .
. 9 8 . . . . 6 .

8 . . . 6 . . . 3
4 . . 8 . 3 . . 1
7 . . . 2 . . . 6
. 6 . . . . 2 8 .
. . . 4 1 9 . . 5
. . . . 8 . . 7 9`

	reader := strings.NewReader(inputWithComments)
	err := board.LoadFromReader(reader)
	if err != nil {
		t.Fatalf("コメント付き入力の読み込みに失敗: %v", err)
	}

	// 正しく読み込まれているかをチェック
	if board.GetCell(0, 0) != 5 {
		t.Errorf("セル(0,0)は5であるべきですが、%d でした", board.GetCell(0, 0))
	}
}

// TestSaveToString は文字列への保存をテスト
func TestSaveToString(t *testing.T) {
	board := NewBoard()

	// 完成した盤面を設定
	completeGrid := [SIZE][SIZE]int{
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
	board.SetBoard(completeGrid)

	output := board.SaveToString()

	// 出力が適切な形式かチェック
	if !strings.Contains(output, "+-------+-------+-------+") {
		t.Error("出力に罫線が含まれていません")
	}

	lines := strings.Split(output, "\n")
	if len(lines) < 13 { // 9 data lines + 4 border lines
		t.Errorf("出力行数が不足しています。期待値: 13行以上, 実際: %d行", len(lines))
	}

	// 数字が正しく表示されているかチェック
	if !strings.Contains(output, "5 3 4") {
		t.Error("最初の行の数字が正しく表示されていません")
	}
}

// TestSaveToWriter はio.Writerへの書き込みをテスト
func TestSaveToWriter(t *testing.T) {
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

	var buffer bytes.Buffer
	err := board.SaveToWriter(&buffer)
	if err != nil {
		t.Fatalf("Writerへの書き込みに失敗: %v", err)
	}

	output := buffer.String()
	expected := board.SaveToString()

	if output != expected {
		t.Error("WriterとStringの出力が一致しません")
	}
}

// TestLoadFromStringErrorCases は文字列読み込みのエラーケースをテスト
func TestLoadFromStringErrorCases(t *testing.T) {
	board := NewBoard()

	testCases := []struct {
		name  string
		input string
	}{
		{
			name:  "行数不足",
			input: "5 3 . . 7 . . . .\n6 . . 1 9 5 . . .", // 2行のみ
		},
		{
			name:  "行数過多",
			input: strings.Repeat("5 3 . . 7 . . . .\n", 10), // 10行
		},
		{
			name:  "無効な文字",
			input: strings.Repeat("5 3 x . 7 . . . .\n", 9), // 'x'が含まれる
		},
		{
			name:  "列数不足",
			input: strings.Repeat("53..7\n", 9), // 5文字のみ
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := board.LoadFromString(tc.input)
			if err == nil {
				t.Errorf("テストケース '%s' でエラーが発生するべきです", tc.name)
			}
		})
	}
}

// TestFromSimpleStringErrorCases はシンプル文字列読み込みのエラーケースをテスト
func TestFromSimpleStringErrorCases(t *testing.T) {
	board := NewBoard()

	testCases := []struct {
		name  string
		input string
	}{
		{
			name:  "行数不足",
			input: "53..7....\n6..195...", // 2行のみ
		},
		{
			name:  "文字数不足",
			input: strings.Repeat("53..7\n", 9), // 5文字のみ
		},
		{
			name:  "無効な文字",
			input: strings.Repeat("53x.7....\n", 9), // 'x'が含まれる
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := board.FromSimpleString(tc.input)
			if err == nil {
				t.Errorf("テストケース '%s' でエラーが発生するべきです", tc.name)
			}
		})
	}
}

// TestRoundTripConversion は変換の往復テスト
func TestRoundTripConversion(t *testing.T) {
	board := NewBoard()

	originalPuzzle := `5 3 . . 7 . . . .
6 . . 1 9 5 . . .
. 9 8 . . . . 6 .
8 . . . 6 . . . 3
4 . . 8 . 3 . . 1
7 . . . 2 . . . 6
. 6 . . . . 2 8 .
. . . 4 1 9 . . 5
. . . . 8 . . 7 9`

	// 文字列から読み込み
	err := board.LoadFromString(originalPuzzle)
	if err != nil {
		t.Fatalf("元のパズルの読み込みに失敗: %v", err)
	}

	// シンプル文字列に変換
	simpleString := board.ToSimpleString()

	// 新しい盤面に読み込み
	board2 := NewBoard()
	err = board2.FromSimpleString(simpleString)
	if err != nil {
		t.Fatalf("シンプル文字列からの読み込みに失敗: %v", err)
	}

	// 両方の盤面が一致することを確認
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if board.GetCell(i, j) != board2.GetCell(i, j) {
				t.Errorf("セル(%d,%d)が一致しません: %d != %d",
					i, j, board.GetCell(i, j), board2.GetCell(i, j))
			}
		}
	}
}

// TestValidateLine はvalidateLine関数をテスト
func TestValidateLine(t *testing.T) {
	board := NewBoard()

	validCases := []string{
		"53..7....", // 連続文字（9文字）
		"53  7    ", // 混在形式（9文字固定、スペースを空欄として扱う）
	}

	for _, input := range validCases {
		err := board.validateLine(input, 1)
		if err != nil {
			t.Errorf("有効な入力 '%s' でエラー: %v", input, err)
		}
	}

	invalidCases := []string{
		"53x.7....",   // 無効な文字
		"53..7",       // 文字数不足
		"53..7.....1", // 文字数過多
	}

	for _, input := range invalidCases {
		err := board.validateLine(input, 1)
		if err == nil {
			t.Errorf("無効な入力 '%s' でエラーが発生するべきです", input)
		}
	}
}

// TestFileOperations はファイル操作のテスト
func TestFileOperations(t *testing.T) {
	// 一時ファイルを作成
	tmpFile, err := os.CreateTemp("", "sudoku_test_*.txt")
	if err != nil {
		t.Fatalf("一時ファイルの作成に失敗: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// テストデータを書き込み
	testData := `5 3 . . 7 . . . .
6 . . 1 9 5 . . .
. 9 8 . . . . 6 .
8 . . . 6 . . . 3
4 . . 8 . 3 . . 1
7 . . . 2 . . . 6
. 6 . . . . 2 8 .
. . . 4 1 9 . . 5
. . . . 8 . . 7 9`

	_, err = tmpFile.WriteString(testData)
	if err != nil {
		t.Fatalf("テストデータの書き込みに失敗: %v", err)
	}

	// ファイルを閉じて再オープン
	tmpFile.Close()

	file, err := os.Open(tmpFile.Name())
	if err != nil {
		t.Fatalf("ファイルのオープンに失敗: %v", err)
	}
	defer file.Close()

	// 盤面に読み込み
	board := NewBoard()
	err = board.LoadFromReader(file)
	if err != nil {
		t.Fatalf("ファイルからの読み込みに失敗: %v", err)
	}

	// 正しく読み込まれているかをチェック
	if board.GetCell(0, 0) != 5 {
		t.Errorf("セル(0,0)は5であるべきですが、%d でした", board.GetCell(0, 0))
	}

	// 解いて保存をテスト
	if board.Solve() {
		// 新しい一時ファイルに保存
		outFile, err := os.CreateTemp("", "sudoku_output_*.txt")
		if err != nil {
			t.Fatalf("出力ファイルの作成に失敗: %v", err)
		}
		defer os.Remove(outFile.Name())
		defer outFile.Close()

		err = board.SaveToWriter(outFile)
		if err != nil {
			t.Fatalf("ファイルへの書き込みに失敗: %v", err)
		}
	}
}

// TestDifferentInputFormats は様々な入力形式のテスト
func TestDifferentInputFormats(t *testing.T) {
	expectedResult := [SIZE][SIZE]int{
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

	formats := []struct {
		name  string
		input string
	}{
		{
			name: "スペース区切り",
			input: `5 3 . . 7 . . . .
6 . . 1 9 5 . . .
. 9 8 . . . . 6 .
8 . . . 6 . . . 3
4 . . 8 . 3 . . 1
7 . . . 2 . . . 6
. 6 . . . . 2 8 .
. . . 4 1 9 . . 5
. . . . 8 . . 7 9`,
		},
		{
			name: "ゼロ使用",
			input: `5 3 0 0 7 0 0 0 0
6 0 0 1 9 5 0 0 0
0 9 8 0 0 0 0 6 0
8 0 0 0 6 0 0 0 3
4 0 0 8 0 3 0 0 1
7 0 0 0 2 0 0 0 6
0 6 0 0 0 0 2 8 0
0 0 0 4 1 9 0 0 5
0 0 0 0 8 0 0 7 9`,
		},
		{
			name: "混在スペース形式",
			input: `53  7....
6  195...
 98    6.
8   6   3
4  8 3  1
7   2   6
 6    28.
   419  5
    8  79`, // 最後の行を修正
		},
	}

	for _, format := range formats {
		t.Run(format.name, func(t *testing.T) {
			board := NewBoard()
			err := board.LoadFromString(format.input)
			if err != nil {
				t.Fatalf("形式 '%s' の読み込みに失敗: %v", format.name, err)
			}

			// 結果を比較
			for i := 0; i < SIZE; i++ {
				for j := 0; j < SIZE; j++ {
					if board.GetCell(i, j) != expectedResult[i][j] {
						t.Errorf("形式 '%s' でセル(%d,%d)が不一致: 期待値=%d, 実際=%d",
							format.name, i, j, expectedResult[i][j], board.GetCell(i, j))
					}
				}
			}
		})
	}
}

// TestLoadFromStringInvalidPuzzleValidation は無効パズル読み込み時の検証をテスト
func TestLoadFromStringInvalidPuzzleValidation(t *testing.T) {
	testCases := []struct {
		name   string
		puzzle string
	}{
		{
			name: "行に重複する数字",
			puzzle: `1 1 . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .`,
		},
		{
			name: "列に重複する数字",
			puzzle: `1 . . . . . . . .
1 . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .`,
		},
		{
			name: "3x3ボックス内に重複する数字",
			puzzle: `1 . . . . . . . .
. 1 . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			board := NewBoard()
			err := board.LoadFromString(tc.puzzle)
			if err == nil {
				t.Errorf("無効なパズル '%s' でエラーが発生するべきです", tc.name)
			} else {
				t.Logf("期待通りエラーが発生: %v", err)
			}
		})
	}
}

// TestFromSimpleStringInvalidPuzzleValidation はシンプル形式での無効パズル検証をテスト
func TestFromSimpleStringInvalidPuzzleValidation(t *testing.T) {
	board := NewBoard()

	// 行に重複する数字
	invalidPuzzle := `11.......
.........
.........
.........
.........
.........
.........
.........
.........`

	err := board.FromSimpleString(invalidPuzzle)
	if err == nil {
		t.Error("無効なパズルでエラーが発生するべきです")
	} else {
		t.Logf("期待通りエラーが発生: %v", err)
	}
}
