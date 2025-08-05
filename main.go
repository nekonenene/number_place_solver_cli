package main

import (
	"fmt"
	"os"
	"sudoku_solver/sudoku"
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
	if puzzle.Solve() {
		fmt.Println("\n解答:")
		puzzle.Print()
		fmt.Println("\n数独の解答が完了しました！")
	} else {
		fmt.Println("\n申し訳ございません。この問題は解けませんでした。")
		fmt.Println("入力に誤りがないか確認してください。")
		os.Exit(1)
	}
}
