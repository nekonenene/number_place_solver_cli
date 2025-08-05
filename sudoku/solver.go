package sudoku

// Solve はバックトラッキングアルゴリズムを使用して数独パズルを解くことを試みる
func (b *Board) Solve() bool {
	// 初期盤面の有効性をチェック
	if !b.IsValidPuzzle() {
		return false
	}

	// 統計情報をリセット
	b.stats = SolveStats{}
	return b.solveRecursive()
}

// solveRecursive はバックトラッキングアルゴリズムを再帰的に実装
func (b *Board) solveRecursive() bool {
	// 最初の空のセルを見つける
	row, col := b.findEmptyCell()
	if row == -1 {
		// 空のセルが見つからない、パズルは解けた
		return true
	}

	// 1-9の数字を試す
	for num := 1; num <= 9; num++ {
		if b.IsValid(row, col, num) {
			// 数字を配置
			b.grid[row][col] = num
			b.stats.CellsSet++

			// 残りを再帰的に解く
			if b.solveRecursive() {
				return true
			}

			// バックトラック：解に繋がらない場合は数字を削除
			b.grid[row][col] = 0
			b.stats.BacktrackCount++
		}
	}

	// このセルに有効な数字が見つからない
	return false
}

// findEmptyCell は盤面内の最初の空のセル（0を含む）を見つける
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

// SolveWithStrategy は最も制約の多い変数ヒューリスティックを使用して解く
func (b *Board) SolveWithStrategy() bool {
	// 初期盤面の有効性をチェック
	if !b.IsValidPuzzle() {
		return false
	}

	return b.solveWithMCV()
}

// solveWithMCV は最も制約の多い変数ヒューリスティックでバックトラッキングを実装
func (b *Board) solveWithMCV() bool {
	// 最も可能性の少ない空のセルを見つける
	row, col := b.findMostConstrainedCell()
	if row == -1 {
		// 空のセルが見つからない、パズルは解けた
		return true
	}

	// 1-9の数字を試す
	for num := 1; num <= 9; num++ {
		if b.IsValid(row, col, num) {
			// 数字を配置
			b.grid[row][col] = num

			// 残りを再帰的に解く
			if b.solveWithMCV() {
				return true
			}

			// バックトラック：解に繋がらない場合は数字を削除
			b.grid[row][col] = 0
		}
	}

	// このセルに有効な数字が見つからない
	return false
}

// findMostConstrainedCell は最も有効な可能性の少ない空のセルを見つける
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

// countPossibilities はセルに配置可能な有効な数字の数を数える
func (b *Board) countPossibilities(row, col int) int {
	count := 0
	for num := 1; num <= 9; num++ {
		if b.IsValid(row, col, num) {
			count++
		}
	}
	return count
}

// HasUniqueSolution はパズルが正確に一つの解を持つかどうかをチェック
func (b *Board) HasUniqueSolution() bool {
	// 最大2つの解まで数える（一意性の判定には十分）
	solutions := b.countSolutions(2)
	return solutions == 1
}

// countSolutions は可能な解の数を数える（maxSolutionsまで）
func (b *Board) countSolutions(maxSolutions int) int {
	// 空のセルが多すぎる場合は計算を避ける
	emptyCells := 0
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if b.grid[i][j] == 0 {
				emptyCells++
			}
		}
	}

	// 空のセルが50個以上の場合は複数解と判定
	if emptyCells > 50 {
		return maxSolutions + 1 // 複数解があることを示す
	}

	// 盤面のコピーを作成
	originalGrid := b.grid
	defer func() { b.grid = originalGrid }()

	return b.countSolutionsRecursive(maxSolutions)
}

// countSolutionsRecursive は再帰的に解の数を数える
func (b *Board) countSolutionsRecursive(maxSolutions int) int {
	// 最初の空のセルを見つける
	row, col := b.findEmptyCell()
	if row == -1 {
		// 空のセルが見つからない、一つの解を発見
		return 1
	}

	solutionCount := 0
	// 1-9の数字を試す
	for num := 1; num <= 9; num++ {
		if b.IsValid(row, col, num) {
			// 数字を配置
			b.grid[row][col] = num

			// 再帰的に解の数を数える
			solutionCount += b.countSolutionsRecursive(maxSolutions)

			// 十分な解が見つかった場合は早期終了
			if maxSolutions > 0 && solutionCount >= maxSolutions {
				b.grid[row][col] = 0
				return solutionCount
			}

			// バックトラック
			b.grid[row][col] = 0
		}
	}

	return solutionCount
}
