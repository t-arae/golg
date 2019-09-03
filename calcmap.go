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

func CalcNextField(b *Boards) []int {
	next := make([]int, b.RCN)
	pre := b.Cells[0:b.RCN]
	for i := 0; i < b.RowN; i++ {
		for j := 0; j < b.ColN; j++ {
			temp := i*b.ColN + j
			temp_index := make([]int, 8)
			for k := 0; k < 8; k++ {
				temp_index[k] = pre[b.IndexAround[temp][k]]
			}
			next[i*b.ColN+j] = calcNext(pre[temp], temp_index)
		}
	}
	return next
}
