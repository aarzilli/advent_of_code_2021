package main

import (
	. "./util"
	"fmt"
	"sort"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

var M [][]byte

type Point struct {
	i, j int
}

func neighbors(i, j int) []Point {
	r := []Point{}
	for _, p := range []Point{Point{i - 1, j}, Point{i, j - 1}, Point{i, j + 1}, Point{i + 1, j}} {
		if p.i < 0 || p.i >= len(M) {
			continue
		}
		if p.j < 0 || p.j >= len(M[p.i]) {
			continue
		}
		r = append(r, p)
	}
	return r
}

func basinsize(i, j int) int {
	BM := make([][]byte, len(M))
	for i := range M {
		BM[i] = make([]byte, len(M[i]))
	}

	basinvisit(BM, i, j)

	r := 0
	for i := range BM {
		for j := range BM[i] {
			//pf("%d", BM[i][j])
			if BM[i][j] != 0 {
				r++
			}
		}
		//pf("\n")
	}
	return r
}

func basinvisit(BM [][]byte, i, j int) {
	if BM[i][j] != 0 {
		return
	}

	BM[i][j] = 1

	for _, p := range neighbors(i, j) {
		if M[p.i][p.j] > M[i][j] && M[p.i][p.j] != 9 {
			basinvisit(BM, p.i, p.j)
		}
	}
}

func main() {
	lines := Input("09.txt", "\n", true)
	pf("len %d\n", len(lines))
	M = make([][]byte, 0, len(lines))
	for i := range lines {
		M = append(M, []byte(lines[i]))
		for j := range M[i] {
			M[i][j] = M[i][j] - '0'
		}
	}
	var risk int
	basins := []int{}
	for i := range M {
		for j := range M[i] {
			ok := true
			for _, p := range neighbors(i, j) {
				if M[i][j] >= M[p.i][p.j] {
					ok = false
					break
				}
			}
			if ok {
				sz := basinsize(i, j)
				pf("low point %d %d (%d %d)\n", i, j, int(M[i][j])+1, sz)
				risk += int(M[i][j]) + 1
				basins = append(basins, sz)
			}
		}
	}
	Sol(risk)

	sort.Ints(basins)

	pf("%d %d %d\n", basins[len(basins)-1], basins[len(basins)-2], basins[len(basins)-3])
	Sol(basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3])
}
