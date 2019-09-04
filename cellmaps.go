package main

import (
	"math/rand"
	"time"
)

type Boards struct {
	RowN   int
	ColN   int
	BoardN int
	Seed   int64
	// Set by Intialize()
	RCN         int
	RCFN        int
	IndexAround [][]int
	Cells       []int
}

func NewBoards(rn, cn, n int) *Boards {
	s := time.Now().Unix()
	b := &Boards{RowN: rn, ColN: cn, BoardN: n, Seed: s}
	b.Initialize()
	return b
}

func (self *Boards) Initialize() {
	rand.Seed(self.Seed)
	self.RCN = self.RowN * self.ColN
	self.RCFN = self.RCN * self.BoardN
	self.IndexAround = link(self)
	self.Cells = make([]int, self.RCFN+1)
	for i := 1; i < (self.RCN + 1); i++ {
		self.Cells[i] = rand.Intn(2)
	}
}

func SetNewBoards(b []int, rn, cn, bn int) *Boards {
	new_b := NewBoards(rn, cn, bn)
	for i := 1; i < (rn*cn + 1); i++ {
		if i < (rn*cn + 1) {
			new_b.Cells[i] = b[i-1]
		} else {
			new_b.Cells[i] = 0
		}
	}
	return new_b
}

func link(b *Boards) [][]int {
	IndexAround := make([][]int, b.RCN)
	for i := 0; i < b.RowN; i++ {
		for j := 0; j < b.ColN; j++ {
			var it, ib, jl, jr int

			switch i {
			case 0:
				it = b.RowN - 1
			default:
				it = i - 1
			}

			if i == (b.RowN - 1) {
				ib = 0
			} else {
				ib = i + 1
			}

			switch j {
			case 0:
				jl = b.ColN - 1
			default:
				jl = j - 1
			}

			if j == (b.ColN - 1) {
				jr = 0
			} else {
				jr = j + 1
			}

			IndexAround[i*b.ColN+j] = []int{
				it*b.ColN + jl,
				it*b.ColN + j,
				it*b.ColN + jr,
				i*b.ColN + jl,
				i*b.ColN + jr,
				ib*b.ColN + jl,
				ib*b.ColN + j,
				ib*b.ColN + jr,
			}
		}
	}
	return IndexAround
}
