package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

type Board struct {
	b       [][]int
	bingoed bool
}

func (b *Board) mark(n int) {
	for i := range b.b {
		for j := range b.b[i] {
			if b.b[i][j] == n {
				b.b[i][j] = -1
			}
		}
	}
}

func (b *Board) bingo() bool {
	if b.bingoed {
		return false
	}
	for i := range b.b {
		ok := true
		for j := range b.b[i] {
			if b.b[i][j] != -1 {
				ok = false
				break
			}
		}
		if ok {
			b.bingoed = true
			return true
		}
	}

	for j := range b.b[0] {
		ok := true
		for i := range b.b {
			if b.b[i][j] != -1 {
				ok = false
				break
			}
		}
		if ok {
			b.bingoed = true
			return true
		}
	}
	return false
}

func (b *Board) sum() int {
	r := 0
	for i := range b.b {
		for j := range b.b[i] {
			if b.b[i][j] != -1 {
				r += b.b[i][j]
			}
		}
	}
	return r
}

func main() {
	groups := Input("04.txt", "\n\n", true)

	in := Vatoi(Spac(groups[0], ",", -1))

	boards := []*Board{}
	first := true
	var last int

	for _, group := range groups[1:] {
		v := Spac(group, "\n", -1)
		board := &Board{}
		board.b = make([][]int, len(v))
		for i := range v {
			board.b[i] = Vatoi(Noempty(Spac(v[i], " ", -1)))
		}
		//pf("%v\n", board)
		boards = append(boards, board)
	}

	for i, n := range in {
		for _, b := range boards {
			b.mark(n)
		}
		for _, b := range boards {
			if b.bingo() {
				pf("bingo %d %d %d\n", i, n, b.sum())
				if first {
					Sol(n, b.sum(), n*b.sum())
					first = false
				}
				last = n * b.sum()

			}
		}
	}
	Sol(last)
}
