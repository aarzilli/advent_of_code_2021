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

var M = map[Point]bool{}

func printmatrix(sp, ep Point) {
	for y := sp.y; y <= ep.y; y++ {
		for x := sp.x; x <= ep.x; x++ {
			if M[Point{x, y}] {
				pf("#")
			} else {
				pf(".")
			}
		}
		pf("\n")
	}
	pf("\n\n")
}

func count() int {
	r := 0
	for range M {
		r++
	}
	return r
}

func main() {
	groups := Input("13.txt", "\n\n", true)
	pf("len %d\n", len(groups))

	lines := Spac(groups[0], "\n", -1)
	pf("%d\n", len(lines))

	for _, line := range lines {
		v := Vatoi(Spac(line, ",", -1))
		M[Point{v[0], v[1]}] = true
	}

	//printmatrix(Point{0,0}, Point{12,14})
	first := true

	for _, line := range Spac(groups[1], "\n", -1) {
		var dir byte
		var n int
		fmt.Sscanf(line, "fold along %c=%d", &dir, &n)
		pf("fold %c %d\n", dir, n)
		switch dir {
		case 'y':
			for p := range M {
				if p.y > n {
					dy := p.y - n
					M[Point{p.x, n - dy}] = true
					delete(M, p)
				} else if p.y == n {
					panic("in overlap")
				}
			}
			//printmatrix(Point{0,0}, Point{12,14})
		case 'x':
			for p := range M {
				if p.x > n {
					dx := p.x - n
					M[Point{n - dx, p.y}] = true
					delete(M, p)
				} else if p.x == n {
					panic("in overlap")
				}
			}
			//printmatrix(Point{0,0}, Point{12,14})
		default:
			panic("blah")
		}
		if first {
			first = false
			Sol(count())
		}
	}

	printmatrix(Point{0, 0}, Point{100, 100})
}
