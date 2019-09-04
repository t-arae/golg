package main

import (
	"flag"
	"fmt"
	"github.com/rivo/tview"
	"reflect"
	"time"
)

const L = "\u25a0"
const D = "  "

func main() {
	var (
		RN     int
		CN     int
		RCN    int
		NMAPS  int
		DELAY  time.Duration
		INF    string
		TEXT   []int
		boards *Boards
	)
	flag.IntVar(&RN, "r", 20, "number of rows, int")
	flag.IntVar(&CN, "c", 20, "number of cols, int")
	flag.DurationVar(&DELAY, "d", 50*time.Millisecond, "delay time, duration")
	flag.StringVar(&INF, "in", "", "input file name, string")
	flag.Parse()

	if INF != "" {
		TEXT, RN, CN = ReadText(INF)
	}

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

	if INF != "" {
		boards = SetNewBoards(TEXT, RN, CN, NMAPS)
	} else {
		boards = NewBoards(RN, CN, NMAPS)
	}
	game_count := 1
	cycle_count := 0
	hist := ""

	go func() {
		last_cm := make([]int, RCN)
		new_cm := make([]int, RCN)
		for {
			cycle_count++
			last_cm = boards.Cells[(len(boards.Cells) - RCN):]
			new_cm = CalcNextField(boards)

			if reflect.DeepEqual(new_cm, last_cm) {
				hist = hist + fmt.Sprintf("game %d (seed %d): %d cycle\n", game_count, boards.Seed, cycle_count)
				hist_list.SetText(hist)
				game_count++
				boards = NewBoards(RN, CN, NMAPS)
				cycle_count = 0
				//break
			} else {
				boards.Cells = append([]int{0}, append(new_cm, boards.Cells[1:RCN+1]...)...)
				text_map.SetText(map2string(new_cm, RN, CN))
				text_cycle.SetText(fmt.Sprintf("cycle: %d", cycle_count))
				time.Sleep(DELAY)
			}
		}
	}()

	menu_list := tview.NewList().
		AddItem("Reset", "press to reset", 'r', func() {
			hist = hist + fmt.Sprintf("game %d (seed %d): %d cycle\n", game_count, boards.Seed, cycle_count)
			hist_list.SetText(hist)
			game_count++
			boards = NewBoards(RN, CN, NMAPS)
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
