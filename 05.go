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

const dopart2 = true

var M = map[Point]int{}

func main() {
	lines := Input("05.txt", "\n", true)
	pf("len %d\n", len(lines))
	for _, line := range lines {
		v := Spac(line, "->", 2)
		coord0 := Vatoi(Spac(v[0], ",", 2))
		coord1 := Vatoi(Spac(v[1], ",", 2))
		//pf("%v %v\n", coord0, coord1)
		if coord0[0] == coord1[0] {
			x := coord0[0]
			starty := Min([]int{coord0[1], coord1[1]})
			endy := Max([]int{coord0[1], coord1[1]})
			for y := starty; y <= endy; y++ {
				M[Point{x, y}]++
			}
		} else if coord0[1] == coord1[1] {
			y := coord0[1]
			startx := Min([]int{coord0[0], coord1[0]})
			endx := Max([]int{coord0[0], coord1[0]})
			for x := startx; x <= endx; x++ {
				M[Point{x, y}]++
			}
		} else if dopart2 {
			deltax := 1
			if coord0[0] > coord1[0] {
				deltax = -1
			}
			deltay := 1
			if coord0[1] > coord1[1] {
				deltay = -1
			}
			//pf("%v %v\n", coord0, coord1)
			x, y := coord0[0], coord0[1]
			for {
				//pf("\t%d %d\n", x, y)
				M[Point{x, y}]++
				if x == coord1[0] && y == coord1[1] {
					break
				}
				x, y = x+deltax, y+deltay
			}
		}
	}

	/*
		for y := 0; y <= 9; y++ {
			for x := 0; x <= 9; x++ {
				pf("%d", M[Point{x,y}])
			}
			pf("\n")
		}
	*/

	part1 := 0
	for _, v := range M {
		if v > 1 {
			part1++
		}
	}
	Sol(part1)
}

// 21554
