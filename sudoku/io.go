package sudoku

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// LoadFromString は文字列表現から数独パズルを読み込む
func (b *Board) LoadFromString(input string) error {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) != SIZE {
		return fmt.Errorf("入力エラー: %d行必要ですが、%d行入力されました", SIZE, len(lines))
	}

	for i, line := range lines {
		// 装飾文字を削除
		line = strings.ReplaceAll(line, "|", "")
		line = strings.ReplaceAll(line, "+", "")
		line = strings.ReplaceAll(line, "-", "")

		// 各文字を順次処理して9個の値を抽出
		var digits []rune
		if strings.Contains(line, " ") {
			// スペースが含まれる場合は、スペース区切りまたは混在形式として処理
			fields := strings.Fields(line)
			if len(fields) == SIZE {
				// 完全にスペース区切りの場合
				for _, field := range fields {
					if field == "." || field == "0" {
						digits = append(digits, '.')
					} else if len(field) == 1 && field[0] >= '1' && field[0] <= '9' {
						digits = append(digits, rune(field[0]))
					} else {
						return fmt.Errorf("%d行目のエラー: 無効な値 '%s' があります", i+1, field)
					}
				}
			} else {
				// 混在形式の場合：文字ごとに処理し、スペースは空欄とする
				for _, r := range line {
					if r >= '1' && r <= '9' {
						digits = append(digits, r)
					} else if r == '.' || r == '0' || r == ' ' {
						digits = append(digits, '.')
					}
					// 他の文字は無視
				}
				if len(digits) != SIZE {
					return fmt.Errorf("%d行目のエラー: %d文字必要ですが、%d文字入力されました", i+1, SIZE, len(digits))
				}
			}
		} else {
			// 連続した文字の場合
			for _, r := range line {
				if (r >= '0' && r <= '9') || r == '.' {
					digits = append(digits, r)
				}
			}
			if len(digits) != SIZE {
				return fmt.Errorf("%d行目のエラー: %d文字必要ですが、%d文字入力されました", i+1, SIZE, len(digits))
			}
		}

		for j, digit := range digits {
			var value int
			if digit == '.' || digit == '0' {
				value = 0
			} else {
				var err error
				value, err = strconv.Atoi(string(digit))
				if err != nil || value < 1 || value > 9 {
					return fmt.Errorf("無効な文字 '%c' が %d行目の%d列目にあります", digit, i+1, j+1)
				}
			}
			b.grid[i][j] = value
		}
	}

	// 読み込み後に盤面の有効性をチェック
	if !b.IsValidPuzzle() {
		return fmt.Errorf("無効な数独パズル: 制約違反があります（重複する数字など）")
	}

	return nil
}

// LoadFromReader はio.Readerから数独パズルを読み込む
func (b *Board) LoadFromReader(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	var lines []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") { // 空行とコメントをスキップ
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	input := strings.Join(lines, "\n")
	return b.LoadFromString(input)
}

// SaveToString は盤面を文字列表現に変換する
func (b *Board) SaveToString() string {
	var builder strings.Builder

	for i := 0; i < SIZE; i++ {
		if i%3 == 0 {
			builder.WriteString("+-------+-------+-------+\n")
		}
		for j := 0; j < SIZE; j++ {
			if j%3 == 0 {
				builder.WriteString("| ")
			}
			if b.grid[i][j] == 0 {
				builder.WriteString(". ")
			} else {
				builder.WriteString(fmt.Sprintf("%d ", b.grid[i][j]))
			}
		}
		builder.WriteString("|\n")
	}
	builder.WriteString("+-------+-------+-------+")

	return builder.String()
}

// SaveToWriter は盤面をio.Writerに書き込む
func (b *Board) SaveToWriter(writer io.Writer) error {
	_, err := writer.Write([]byte(b.SaveToString()))
	return err
}

// ToSimpleString は盤面をシンプルな9行形式に変換する
func (b *Board) ToSimpleString() string {
	var builder strings.Builder

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if b.grid[i][j] == 0 {
				builder.WriteString(".")
			} else {
				builder.WriteString(strconv.Itoa(b.grid[i][j]))
			}
		}
		if i < SIZE-1 {
			builder.WriteString("\n")
		}
	}

	return builder.String()
}

// FromSimpleString はシンプルな9行形式から読み込む
func (b *Board) FromSimpleString(input string) error {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) != SIZE {
		return fmt.Errorf("invalid input: expected %d lines, got %d", SIZE, len(lines))
	}

	for i, line := range lines {
		if len(line) != SIZE {
			return fmt.Errorf("invalid input at line %d: expected %d characters, got %d", i+1, SIZE, len(line))
		}

		for j, char := range line {
			var value int
			if char == '.' || char == '0' {
				value = 0
			} else if char >= '1' && char <= '9' {
				value = int(char - '0')
			} else {
				return fmt.Errorf("invalid character '%c' at position (%d, %d)", char, i+1, j+1)
			}
			b.grid[i][j] = value
		}
	}

	// 読み込み後に盤面の有効性をチェック
	if !b.IsValidPuzzle() {
		return fmt.Errorf("無効な数独パズル: 制約違反があります（重複する数字など）")
	}

	return nil
}

// LoadFromInteractiveInput は対話形式で数独パズルを入力する
func (b *Board) LoadFromInteractiveInput() error {
	fmt.Println("数独の問題を入力してください（9行×9列）")
	fmt.Println("空欄は半角スペース、ドット(.)、または0で入力できます")
	fmt.Println("例: 53..7....")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for i := 0; i < SIZE; i++ {
		for {
			fmt.Printf("%d行目を入力してください: ", i+1)
			if !scanner.Scan() {
				return fmt.Errorf("入力の読み取りに失敗しました")
			}

			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				fmt.Println("空行は無効です。もう一度入力してください。")
				continue
			}

			// 入力の検証
			if err := b.validateLine(line, i+1); err != nil {
				fmt.Printf("エラー: %v\n", err)
				fmt.Println("もう一度入力してください。")
				continue
			}

			lines = append(lines, line)
			break
		}
	}

	// 全ての行をまとめて処理
	input := strings.Join(lines, "\n")
	return b.LoadFromString(input)
}

// validateLine は入力行の妥当性をチェック
func (b *Board) validateLine(line string, lineNum int) error {
	// 装飾文字を削除
	line = strings.ReplaceAll(line, "|", "")
	line = strings.ReplaceAll(line, "+", "")
	line = strings.ReplaceAll(line, "-", "")

	// 数字とドットのみを抽出（スペースは除去）
	var digits []rune
	for _, r := range line {
		if (r >= '0' && r <= '9') || r == '.' {
			digits = append(digits, r)
		} else if r == ' ' {
			// スペースは空欄として扱う
			digits = append(digits, '.')
		}
	}

	if len(digits) != SIZE {
		return fmt.Errorf("%d文字必要ですが、%d文字入力されました", SIZE, len(digits))
	}

	// 各文字の検証
	for j, digit := range digits {
		if digit != '.' && digit != '0' && (digit < '1' || digit > '9') {
			return fmt.Errorf("無効な文字 '%c' が%d列目にあります", digit, j+1)
		}
	}

	return nil
}
