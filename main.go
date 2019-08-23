package main

import (
	"flag"
	"fmt"
	"github.com/rivo/tview"
	"math/rand"
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

	text_list := tview.NewTextView().SetScrollable(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	cms := NewCellmaps(RN, CN, NMAPS)

	go func() {
		game_count := 1
		cycle_count := 0
		temp_list := ""
		pre_cm := make([]int, RCN)
		last_cm := make([]int, RCN)
		new_cm := make([]int, RCN)
		for {
			cycle_count++
			temp := ""

			pre_cm = cms.cells[0:RCN]
			last_cm = cms.cells[RCN:]
			new_cm = next_map(RN, CN, pre_cm)

			if reflect.DeepEqual(new_cm, last_cm) {
				temp_list = temp_list + fmt.Sprintf("game %d: %d cycle\n", game_count, cycle_count)
				text_list.SetText(temp_list)
				game_count++
				cms = NewCellmaps(RN, CN, NMAPS)
				cycle_count = 0
				//break
			} else {
				for rn := 0; rn < RN; rn++ {
					for cn := 0; cn < CN; cn++ {
						if new_cm[rn*CN+cn] == 0 {
							temp = temp + D
						} else {
							temp = temp + L
						}
					}
					temp = temp + "\n"
				}

				cms.cells = append(new_cm, cms.cells[0:RCN]...)

				text_map.SetText(temp).SetTextAlign(tview.AlignCenter)
				text_cycle.SetText(fmt.Sprintf("cycle: %d", cycle_count)).
					SetTextAlign(tview.AlignCenter)
				time.Sleep(DELAY)
			}
		}
	}()

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(text_cycle, 0, 1, true).
			AddItem(text_map, 0, 9, true), 0, 1, true).
		AddItem(text_list, 0, 1, true)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

type Cellmaps struct {
	rn    int
	cn    int
	n     int
	cells []int
}

func NewCellmaps(rn, cn, n int) *Cellmaps {
	cms := &Cellmaps{rn: rn, cn: cn, n: n}
	cms.Initialize()
	return cms
}
func (self *Cellmaps) Initialize() {
	self.cells = make([]int, self.cn*self.rn*self.n)
	for i := 0; i < self.cn*self.rn; i++ {
		self.cells[i] = rand.Intn(2)
	}
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
