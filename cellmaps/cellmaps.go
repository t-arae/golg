package cellmaps

import (
	"math/rand"
	"time"
)

type Cellmaps struct {
	rn    int
	cn    int
	n     int
	Seed  int64
	Cells []int
}

func NewCellmaps(rn, cn, n int) *Cellmaps {
	s := time.Now().Unix()
	cms := &Cellmaps{rn: rn, cn: cn, n: n, Seed: s}
	cms.Initialize()
	return cms
}

func (self *Cellmaps) Initialize() {
	rand.Seed(self.Seed)
	self.Cells = make([]int, self.cn*self.rn*self.n)
	for i := 0; i < self.cn*self.rn; i++ {
		self.Cells[i] = rand.Intn(2)
	}
}
