# golg
Conway's game of life written in golang

## Installation

```bash
go get github.com/t-arae/golg
```

## Usage

```
Usage of golg:
  -c int
    	number of cols, int (default 20)
  -d duration
    	delay time, duration (default 50ms)
  -r int
    	number of rows, int (default 20)
```

## ToDo
- [x] 引数設定でプログラム実行時に変数を設定できるように
- [ ] ファイルからセルをリストア
- [ ] ファイルにセルの状態を保存
- [ ] 任意の周期で収束した時，停止して保存する
