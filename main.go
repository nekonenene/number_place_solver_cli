package main

import (
	"fmt"
	"os"
	"sudoku_solver/sudoku"
	"time"
)

func main() {
	fmt.Println("数独ソルバー")
	fmt.Println("============")
	fmt.Println()

	// 空の数独盤面を作成
	puzzle := sudoku.NewBoard()

	// 対話形式でパズルを入力
	if err := puzzle.LoadFromInteractiveInput(); err != nil {
		fmt.Printf("入力エラー: %v\n", err)
		os.Exit(1)
	}

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
		fmt.Printf("\n数独の解答が完了しました！\n")
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
