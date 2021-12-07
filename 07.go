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

func main() {
	lines := Input("07.txt", "\n", true)
	pf("len %d\n", len(lines))
	in := Vatoi(Spac(lines[0], ",", -1))
	pf("%d %d\n",Min(in), Max(in))
	pf("example: %d\n", dist2for1(16, 5))
	dist := dist2
	mind := 10000000000
	mintgt := -1
	for tgt := Min(in); tgt <= Max(in); tgt++ {
		d := dist(in, tgt)
		if d < mind {
			mind = d
			mintgt = tgt
		}
	}
	Sol(mintgt, mind)
}
