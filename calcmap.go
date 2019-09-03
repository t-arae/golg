package main

func calcNext(me int, others []int) int {
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

func CalcNextField(row_n int, col_n int, mapA []int) []int {
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
			mapB[i*col_n+j] = calcNext(mapA[i*col_n+j], temp_others[:])
		}
	}
	return mapB
}
