package main

import (
	"math/rand"
	"time"
)

type Feilds struct {
	rn    int
	cn    int
	n     int
	Seed  int64
	Cells []int
}

func NewFields(rn, cn, n int) *Feilds {
	s := time.Now().Unix()
	cms := &Feilds{rn: rn, cn: cn, n: n, Seed: s}
	cms.Initialize()
	return cms
}

func (self *Feilds) Initialize() {
	rand.Seed(self.Seed)
	self.Cells = make([]int, self.cn*self.rn*self.n)
	for i := 0; i < self.cn*self.rn; i++ {
		self.Cells[i] = rand.Intn(2)
	}
}
