package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func dist1(in []int, tgt int) int {
	r := 0
	for _, x := range in {
		r += Abs(x - tgt)
	}
	return r
}

func dist2for1(x, tgt int) int {
	n := Abs(tgt - x)
	return (n + n*n) / 2
}

func dist2(in []int, tgt int) int {
	r := 0
	for _, x := range in {
		r += dist2for1(x, tgt)
	}
	return r
}

func mindist(in []int, dist func([]int, int) int) (mintgt, mind int) {
	mind = 0
	mintgt = -1
	for tgt := Min(in); tgt <= Max(in); tgt++ {
		d := dist(in, tgt)
		if d < mind || mintgt == -1 {
			mind = d
			mintgt = tgt
		}
	}
	return mintgt, mind
}

func main() {
	lines := Input("07.txt", "\n", true)
	in := Vatoi(Spac(lines[0], ",", -1))
	mintgt, mind := mindist(in, dist1)
	Sol(mintgt, mind)
	mintgt, mind = mindist(in, dist2)
	Sol(mintgt, mind)
}
