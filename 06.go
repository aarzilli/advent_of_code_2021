package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

var fishes []int
var fm = map[int]int{}

func step() {
	for i := range fishes {
		if fishes[i] == 0 {
			fishes[i] = 6
			fishes = append(fishes, 8)
		} else {
			fishes[i]--
		}
	}
}

func mstep() {
	nfm := map[int]int{}
	for k, v := range fm {
		if k == 0 {
			k = 6
			nfm[8] += v
			nfm[k] += v
		} else {
			nfm[k-1] += v
		}
	}
	fm = nfm
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
	pf("len %d\n", len(lines))
	fishes = Vatoi(Spac(lines[0], ",", -1))
	for _, fish := range fishes {
		fm[fish]++
	}
	for i := 0; i < 256; i++ {
		pf("day %d: %d %d\n", i, len(fishes), mcount())
		/*if len(fishes) != mcount() {
			panic("blah")
		}*/
		//step()
		mstep()
		//pf("after %d: %v\n", i+1, fishes)
	}
	Sol(mcount())
}

// after 2: [1 2 1 6 0 8]
// After  2 days: 1,2,1,6,0,8

// 6 0 6 4 5 6 0 1 1 2 6 0 1 1 1 2 2 3 3 4 6 7 8 8 8 8
// 6,0,6,4,5,6,0,1,1,2,6,0,1,1,1,2,2,3,3,4,6,7,8,8,8,8
