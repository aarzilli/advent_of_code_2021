package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

var fm [9]int

func mstep() {
	var nfm [9]int
	for k, v := range fm {
		if k == 0 {
			nfm[8] += v
			nfm[6] += v
		} else {
			nfm[k-1] += v
		}
	}
	copy(fm[:], nfm[:])
}

func mcount() int {
	r := 0
	for _, v := range fm {
		r += v
	}
	return r
}

func main() {
	lines := Input("06.txt", "\n", true)
	fishes := Vatoi(Spac(lines[0], ",", -1))
	for _, fish := range fishes {
		fm[fish]++
	}
	for i := 0; i < 256; i++ {
		if i == 80 {
			Sol(mcount())
		}
		mstep()
	}
	Sol(mcount())
}
