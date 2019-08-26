package main

import (
	"flag"
	"fmt"
	"github.com/rivo/tview"
	"github.com/t-arae/golg/cellmaps"
	"reflect"
	"time"
)

const L = "\u25a0"
const D = "  "

func main() {
	var (
		RN    int
		CN    int
		RCN   int
		NMAPS int
		DELAY time.Duration
	)
	flag.IntVar(&RN, "r", 20, "number of rows, int")
	flag.IntVar(&CN, "c", 20, "number of cols, int")
	flag.DurationVar(&DELAY, "d", 50*time.Millisecond, "delay time, duration")
	flag.Parse()

	RCN = RN * CN
	NMAPS = 2
	app := tview.NewApplication()

	text_map := tview.NewTextView().SetChangedFunc(func() {
		app.Draw()
	})

	text_cycle := tview.NewTextView().SetChangedFunc(func() {
		app.Draw()
	})

	hist_list := tview.NewTextView().SetScrollable(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	cms := cellmaps.NewCellmaps(RN, CN, NMAPS)
	game_count := 1
	cycle_count := 0
	hist := ""

	go func() {
		//game_count := 1
		//cycle_count := 0
		//hist := ""
		pre_cm := make([]int, RCN)
		last_cm := make([]int, RCN)
		new_cm := make([]int, RCN)
		for {
			cycle_count++
			pre_cm = cms.Cells[0:RCN]
			last_cm = cms.Cells[RCN:]
			new_cm = next_map(RN, CN, pre_cm)

			if reflect.DeepEqual(new_cm, last_cm) {
				hist = hist + fmt.Sprintf("game %d (seed %d): %d cycle\n", game_count, cms.Seed, cycle_count)
				hist_list.SetText(hist)
				game_count++
				cms = cellmaps.NewCellmaps(RN, CN, NMAPS)
				cycle_count = 0
				//break
			} else {
				cms.Cells = append(new_cm, cms.Cells[0:RCN]...)
				text_map.SetText(map2string(new_cm, RN, CN))
				text_cycle.SetText(fmt.Sprintf("cycle: %d", cycle_count))
				time.Sleep(DELAY)
			}
		}
	}()

	menu_list := tview.NewList().
		AddItem("Reset", "press to reset", 'r', func() {
			hist = hist + fmt.Sprintf("game %d (seed %d): %d cycle\n", game_count, cms.Seed, cycle_count)
			hist_list.SetText(hist)
			game_count++
			cms = cellmaps.NewCellmaps(RN, CN, NMAPS)
			cycle_count = 0
		}).
		AddItem("Quit", "press to exit", 'q', func() {
			app.Stop()
		})

	flex := tview.NewFlex().
		AddItem(tview.NewTextView(), 3, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(hist_list, 0, 8, true).
			AddItem(menu_list, 0, 2, true), 25, 0, false).
		AddItem(tview.NewTextView(), 3, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(text_cycle, 0, 1, true).
			AddItem(text_map, 0, 9, true), 0, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(menu_list).Run(); err != nil {
		panic(err)
	}
}

func map2string(m []int, RN int, CN int) string {
	s := ""
	for i := 0; i < RN; i++ {
		for j := 0; j < CN; j++ {
			if m[i*CN+j] == 0 {
				s = s + D
			} else {
				s = s + L
			}
		}
		s = s + "\n"
	}
	return s
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
	map_len := col_n * row_n
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
