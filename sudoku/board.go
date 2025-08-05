package sudoku

import "fmt"

const SIZE = 9 // 数独盤面のサイズ

// Board は検証と解法機能を持つ数独盤面を表す
type Board struct {
	grid [SIZE][SIZE]int
}

// NewBoard は新しい空の数独盤面を作成する
func NewBoard() *Board {
	return &Board{}
}

// SetBoard は指定された9x9グリッドで盤面を設定する
func (b *Board) SetBoard(grid [SIZE][SIZE]int) {
	b.grid = grid
}

// GetBoard は現在の盤面状態のコピーを返す
func (b *Board) GetBoard() [SIZE][SIZE]int {
	return b.grid
}

// SetCell は指定された位置に値を設定する
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

// GetCell は指定された位置の値を返す
func (b *Board) GetCell(row, col int) int {
	if row < 0 || row >= SIZE || col < 0 || col >= SIZE {
		return -1
	}
	return b.grid[row][col]
}

// IsEmpty はセルが空かどうか（0が入っているか）をチェック
func (b *Board) IsEmpty(row, col int) bool {
	return b.GetCell(row, col) == 0
}

// Print は現在の盤面状態を整形した形で表示する
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

// IsValid は指定された位置に数字を配置できるかどうかをチェック
func (b *Board) IsValid(row, col, num int) bool {
	// 行の制約をチェック
	for j := 0; j < SIZE; j++ {
		if b.grid[row][j] == num {
			return false
		}
	}

	// 列の制約をチェック
	for i := 0; i < SIZE; i++ {
		if b.grid[i][col] == num {
			return false
		}
	}

	// 3x3ボックスの制約をチェック
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

// IsComplete は盤面が完全に埋められているかどうかをチェック
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

// IsSolved は盤面が完成しておりかつ有効かどうかをチェック
func (b *Board) IsSolved() bool {
	if !b.IsComplete() {
		return false
	}

	// 全ての制約をチェック
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			num := b.grid[i][j]
			b.grid[i][j] = 0 // 有効性チェックのため一時的に削除
			if !b.IsValid(i, j, num) {
				b.grid[i][j] = num // もとに戻す
				return false
			}
			b.grid[i][j] = num // もとに戻す
		}
	}
	return true
}
