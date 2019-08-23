package main

import (
	"flag"
	"fmt"
	"github.com/rivo/tview"
	"math/rand"
	"time"
)

const L = "\u25a0"
const D = "  "

func main() {
	var (
		RN    int
		CN    int
		DELAY time.Duration
	)
	flag.IntVar(&RN, "r", 20, "number of rows, int")
	flag.IntVar(&CN, "c", 20, "number of cols, int")
	flag.DurationVar(&DELAY, "d", 50*time.Millisecond, "delay time, duration")
	flag.Parse()

	cell_map := create_newmap(RN, CN)

	app := tview.NewApplication()

	textView := tview.NewTextView().SetChangedFunc(func() {
		app.Draw()
	})

	textView2 := tview.NewTextView().SetChangedFunc(func() {
		app.Draw()
	})

	go func() {
		pre_map := cell_map[:]
		count := 0
		for {
			count++
			temp := ""
			new_map := next_map(RN, CN, pre_map)
			for rn := 0; rn < RN; rn++ {
				for cn := 0; cn < CN; cn++ {
					if new_map[rn*CN+cn] == 0 {
						temp = temp + D
					} else {
						temp = temp + L
					}
				}
				temp = temp + "\n"
			}
			//textView.Clear()
			//fmt.Fprintf(textView, "%s\n", temp)
			textView.SetText(temp).SetTextAlign(tview.AlignCenter)
			textView2.SetText(fmt.Sprintf("cycle: %d", count)).
				SetTextAlign(tview.AlignCenter)
			pre_map = new_map[:]
			time.Sleep(DELAY)
		}
	}()

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(textView2, 0, 1, true).
		AddItem(textView, 0, 10, true)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

func create_newmap(row_n int, col_n int) []int {
	rand.Seed(time.Now().Unix())
	new_cell_map := make([]int, row_n*col_n)
	for i := 0; i < row_n*col_n; i++ {
		new_cell_map[i] = rand.Intn(2)
	}
	return new_cell_map
}

func calc_next(me int, others []int) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += others[i]
	}

	if me == 0 && sum == 3 {
		return 1
	}
	if me == 1 && ((sum == 2) || (sum == 3)) {
		return 1
	}
	return 0
}

func next_map(row_n int, col_n int, mapA []int) []int {
	map_len := len(mapA)
	mapB := make([]int, map_len)
	for i := 0; i < row_n; i++ {
		for j := 0; j < col_n; j++ {
			var it, ib, jl, jr int

			switch i {
			case 0:
				it = row_n - 1
			default:
				it = i - 1
			}

			if i == (row_n - 1) {
				ib = 0
			} else {
				ib = i + 1
			}

			switch j {
			case 0:
				jl = col_n - 1
			default:
				jl = j - 1
			}

			if j == (col_n - 1) {
				jr = 0
			} else {
				jr = j + 1
			}

			temp_others := [8]int{
				mapA[it*col_n+jl],
				mapA[it*col_n+j],
				mapA[it*col_n+jr],

				mapA[i*col_n+jl],
				mapA[i*col_n+jr],

				mapA[ib*col_n+jl],
				mapA[ib*col_n+j],
				mapA[ib*col_n+jr],
			}
			mapB[i*col_n+j] = calc_next(mapA[i*col_n+j], temp_others[:])
		}
	}
	return mapB
}
