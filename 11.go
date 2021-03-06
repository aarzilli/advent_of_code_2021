package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

var M [][]int
var totalflashes int
var magicflash bool

func step() {
	for i := range M {
		for j := range M[i] {
			M[i][j]++
		}
	}

	n := numflashing()
	flashed := map[[2]int]bool{}
	for {
		flashonce(flashed)
		n2 := numflashing()
		if n2 == n {
			break
		}
		n = n2
	}

	if numflashing() == len(M)*len(M[0]) {
		magicflash = true
	}
}

func flashonce(flashed map[[2]int]bool) {
	toflash := [][2]int{}
	for i := range M {
		for j := range M[i] {
			if M[i][j] > 9 && !flashed[[2]int{i, j}] {
				toflash = append(toflash, [2]int{i, j})
			}
		}
	}

	for _, p := range toflash {
		flashed[p] = true
		flashone(p[0], p[1])
	}
}

func numflashing() int {
	r := 0
	for i := range M {
		for j := range M[i] {
			if M[i][j] > 9 {
				r++
			}
		}
	}
	return r
}

func flashone(i, j int) {
	totalflashes++
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			if i+di < 0 || i+di >= len(M) {
				continue
			}
			if j+dj < 0 || j+dj >= len(M[i+di]) {
				continue
			}
			M[i+di][j+dj]++
		}
	}
}

func resetflashed() {
	for i := range M {
		for j := range M[i] {
			if M[i][j] > 9 {
				M[i][j] = 0
			}
		}
	}
}

func main() {
	lines := Input("11.txt", "\n", true)

	for i, line := range lines {
		M = append(M, make([]int, len(line)))
		for j := range line {
			M[i][j] = int(line[j] - '0')
		}
	}

	for i := 1; i <= 1000; i++ {
		step()
		resetflashed()
		if i == 100 {
			Sol(totalflashes)
		}
		if magicflash {
			Sol(i)
			break
		}
	}
}
