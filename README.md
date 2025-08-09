# ナンバープレース（数独）を CLI 上で解くやつ

ナンバープレース（数独）を解いてくれます。  
あまりにも解けない問題があったときに、「本当に解ける問題？！ 実は誤植ない？！」とムシャクシャして作りました。  
（これに解かせてみたら解けたので、安心して自力で解くことに集中できました）

なお、**「数独」や「SUDOKU」は[株式会社ニコリ](https://www.nikoli.co.jp/ja/)の登録商標**です。


## インストール

### 必要なもの

- Go 1.24 以上

### ビルド方法

```bash
git clone https://github.com/nekonenene/number_place_solver_cli.git
cd number_place_solver_cli

go build -o bin/number_place_solver
```

## 使い方

上の「ビルド方法」の項でビルドした前提にしていますが、  
もちろん、clone したあとに `go run main.go` する方法でも大丈夫です。

### 基本的な使用方法

```bash
# 対話形式で入力（ `Ctrl + C` で途中終了できます）
./bin/number_place_solver

# ファイルを指定して読み込み
./bin/number_place_solver puzzle.txt

# 文字列から直接読み込み
./bin/number_place_solver "53..7...."

# ヘルプを表示
./bin/number_place_solver -h
```

### 使用例

```bash
# サンプルファイルを解く
./bin/number_place_solver number_place/testdata/easy.txt

# 文字列で問題を渡す
./bin/number_place_solver "53..7....
6..195...
.98....6.
8...6...3
4..8.3..1
7...2...6
.6....28.
...419..5
....8..79"

# ファイル内容を引数として渡す
./bin/number_place_solver "$(cat puzzle.txt)"
```


## 開発

### テストの実行

```bash
# 全てのテストを実行
go test ./...

# ベンチマークテストを実行
go test -bench=. ./number_place

# カバレッジを確認
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```
