package main

import (
	"fmt"
	"number_place_solver/number_place"
	"os"
	"time"
)

// showUsage は使用方法を表示する
func showUsage() {
	fmt.Printf("ナンバープレースを解きます\n")
	fmt.Printf("============\n\n")
	fmt.Printf("使用方法:\n")
	fmt.Printf("  %s                    # 対話形式で入力\n", os.Args[0])
	fmt.Printf("  %s <file>             # ファイルから読み込み\n", os.Args[0])
	fmt.Printf("  %s \"<puzzle>\"          # 改行区切り文字列から読み込み\n\n", os.Args[0])
	fmt.Printf("例:\n")
	fmt.Printf("  %s puzzle.txt\n", os.Args[0])
	fmt.Printf("  %s \"5 3 . . 7 . . . .\n", os.Args[0])
	fmt.Printf("6 . . 1 9 5 . . .\n")
	fmt.Printf(". 9 8 . . . . 6 .\n")
	fmt.Printf("...\"\n")
	fmt.Printf("  %s \"$(cat puzzle.txt)\" # ファイル内容を引数として渡す\n\n", os.Args[0])
}

// loadPuzzle は引数に応じてパズルを読み込む
func loadPuzzle(puzzle *number_place.Board) error {
	args := os.Args[1:]

	if len(args) == 0 {
		// 引数なし：対話形式
		fmt.Println("ナンバープレースを解きます")
		fmt.Println("============")
		fmt.Println()
		return puzzle.LoadFromInteractiveInput()
	}

	if len(args) == 1 {
		arg := args[0]

		// ファイル存在確認
		if _, err := os.Stat(arg); err == nil {
			// ファイルが存在する場合：ファイルから読み込み
			file, err := os.Open(arg)
			if err != nil {
				return fmt.Errorf("ファイルのオープンに失敗: %v", err)
			}
			defer file.Close()

			fmt.Printf("ファイル '%s' から読み込み中...\n", arg)
			return puzzle.LoadFromReader(file)
		} else {
			// ファイルが存在しない場合：引数を改行区切り文字列として扱う
			fmt.Println("引数から読み込み中...")
			return puzzle.LoadFromString(arg)
		}
	}

	// 引数が2つ以上の場合
	return fmt.Errorf("引数が多すぎます。使用方法を確認してください")
}

// solvePuzzle はパズルを解いて結果を表示する
func solvePuzzle(puzzle *number_place.Board) {
	fmt.Println("\n入力された問題:")
	puzzle.Print()

	// パズルを解く
	fmt.Println("\n解答中...")
	startTime := time.Now()
	solved := puzzle.Solve()
	elapsedTime := time.Since(startTime)

	if solved {
		fmt.Println("\n解答:")
		puzzle.Print()

		// 統計情報を表示
		stats := puzzle.GetStats()
		fmt.Printf("\nナンバープレースの解答が完了しました！\n")
		fmt.Printf("解答時間: %v\n", elapsedTime)
		fmt.Printf("セル設定回数: %d回\n", stats.CellsSet)
		fmt.Printf("バックトラック回数: %d回\n", stats.BacktrackCount)
	} else {
		fmt.Println("\n申し訳ございません。この問題は解けませんでした。")

		// 統計情報を表示
		stats := puzzle.GetStats()
		fmt.Printf("解答時間: %v\n", elapsedTime)
		fmt.Printf("試行回数: セル設定 %d回、バックトラック %d回\n", stats.CellsSet, stats.BacktrackCount)
		fmt.Println("入力に誤りがないか確認してください。")
		os.Exit(1)
	}
}

func main() {
	// ヘルプ表示の確認
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		showUsage()
		return
	}

	// 空のナンバープレースの盤面を作成
	puzzle := number_place.NewBoard()

	// パズルを読み込み
	if err := loadPuzzle(puzzle); err != nil {
		fmt.Printf("入力エラー: %v\n", err)
		fmt.Println()
		showUsage()
		os.Exit(1)
	}

	// パズルを解く
	solvePuzzle(puzzle)
}
