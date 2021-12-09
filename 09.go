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

func neighbors(p0 Point) []Point {
	r := []Point{}
	for _, p := range []Point{Point{p0.i - 1, p0.j}, Point{p0.i, p0.j - 1}, Point{p0.i, p0.j + 1}, Point{p0.i + 1, p0.j}} {
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

func basinsize(p0 Point) int {
	B := map[Point]bool{}
	basinvisit(B, p0)
	return len(B)
}

func basinvisit(B map[Point]bool, p0 Point) {
	if B[p0] {
		return
	}

	B[p0] = true

	for _, p := range neighbors(p0) {
		if M[p.i][p.j] > M[p0.i][p0.j] && M[p.i][p.j] != 9 {
			basinvisit(B, p)
		}
	}
}

func main() {
	lines := Input("09.txt", "\n", true)
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
			for _, p := range neighbors(Point{i, j}) {
				if M[i][j] >= M[p.i][p.j] {
					ok = false
					break
				}
			}
			if ok {
				sz := basinsize(Point{i, j})
				risk += int(M[i][j]) + 1
				basins = append(basins, sz)
			}
		}
	}
	Sol(risk)
	sort.Ints(basins)
	Sol(basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3])
}
