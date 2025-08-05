package sudoku

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// LoadFromString は文字列表現から数独パズルを読み込む
func (b *Board) LoadFromString(input string) error {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) != SIZE {
		return fmt.Errorf("invalid input: expected %d lines, got %d", SIZE, len(lines))
	}

	for i, line := range lines {
		// スペースを削除して数字を解析
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "|", "")
		line = strings.ReplaceAll(line, "+", "")
		line = strings.ReplaceAll(line, "-", "")

		// '.'と'0'以外の非数字文字をフィルタリング
		var digits []rune
		for _, r := range line {
			if (r >= '0' && r <= '9') || r == '.' {
				digits = append(digits, r)
			}
		}

		if len(digits) != SIZE {
			return fmt.Errorf("invalid input at line %d: expected %d digits, got %d", i+1, SIZE, len(digits))
		}

		for j, digit := range digits {
			var value int
			if digit == '.' || digit == '0' {
				value = 0
			} else {
				var err error
				value, err = strconv.Atoi(string(digit))
				if err != nil || value < 1 || value > 9 {
					return fmt.Errorf("invalid digit '%c' at position (%d, %d)", digit, i+1, j+1)
				}
			}
			b.grid[i][j] = value
		}
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

	return nil
}
