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

func printmatrix() {
	for i := range M {
		for j := range M[i] {
			if M[i][j] > 9 {
				pf("X")
			} else {
				pf("%d", M[i][j])
			}
		}
		pf("\n")
	}
	pf("\n")
}

func step() {
	for i := range M {
		for j := range M[i] {
			M[i][j]++
		}
	}

	// flashing
	n := numflashing()
	flashed := map[[2]int]bool{}
	for {
		//pf("flashing\n")
		//printmatrix()
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
	//M2 := make([][]int, len(M))
	for i := range M {
		//M2[i] = make([]int, len(M[i])
		for j := range M[i] {
			//M2[i][j] = M[i]
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
	pf("len %d\n", len(lines))

	for i, line := range lines {
		M = append(M, make([]int, len(line)))
		for j := range line {
			M[i][j] = int(line[j] - '0')
		}
	}

	printmatrix()

	for i := 1; i <= 1000; i++ {
		//pf("step %d:\n", i)
		step()
		//printmatrix()
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
