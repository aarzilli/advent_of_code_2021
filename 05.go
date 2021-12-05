package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

type Point struct {
	x, y int
}

func sign(x int) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -1
	}
	return 1
}

func trace(M map[Point]int, start, end []int) {
	deltax := sign(end[0] - start[0])
	deltay := sign(end[1] - start[1])
	x, y := start[0], start[1]
	for {
		M[Point{x, y}]++
		if x == end[0] && y == end[1] {
			break
		}
		x, y = x+deltax, y+deltay
	}
}

func overlaps(M map[Point]int) int {
	r := 0
	for _, v := range M {
		if v > 1 {
			r++
		}
	}
	return r
}

const dopart2 = true

func main() {
	lines := Input("05.txt", "\n", true)
	var M1 = map[Point]int{}
	var M2 = map[Point]int{}

	for _, line := range lines {
		v := Spac(line, "->", 2)
		start := Vatoi(Spac(v[0], ",", 2))
		end := Vatoi(Spac(v[1], ",", 2))
		if start[0] == end[0] || start[1] == end[1] {
			trace(M1, start, end)
		}
		trace(M2, start, end)
	}

	Sol(overlaps(M1))
	Sol(overlaps(M2))
}
