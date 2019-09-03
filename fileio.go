package main

import (
	"bufio"
	"io"
	"os"
)

func ReadText(path string) (res []int, rn int, cn int) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReaderSize(f, 1024)
	res = []int{}
	rn = 0
	cn = 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		rn++
		cn = len(line)
		res = append(res, byte2int(line[:cn])...)
	}
	return res, rn, cn
}

func byte2int(bs []byte) []int {
	bs_len := len(bs)
	is := make([]int, bs_len)
	for i := 0; i < (bs_len); i++ {
		is[i] = int(bs[i]) - 48
	}
	return is
}
