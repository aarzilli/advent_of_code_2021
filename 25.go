package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

var M [][]byte

func printmatrix() {
	for i := range M {
		for j := range M[i] {
			pf("%c", M[i][j])
		}
		pf("\n")
	}
	pf("\n")
}

type Point struct {
	i, j int
}

func step() bool {
	change := false
	tomove := []Point{}
	for i := range M {
		for j := range M[i] {
			if M[i][j] == '>' {
				nj := (j + 1) % len(M[i])
				if M[i][nj] == '.' {
					tomove = append(tomove, Point{i, j})
					change = true
				}
			}
		}
	}

	for _, p := range tomove {
		nj := (p.j + 1) % len(M[p.i])
		M[p.i][p.j] = '.'
		M[p.i][nj] = '>'
	}

	tomove = tomove[:0]
	for i := range M {
		for j := range M[i] {
			if M[i][j] == 'v' {
				ni := (i + 1) % len(M)
				if M[ni][j] == '.' {
					tomove = append(tomove, Point{i, j})
					change = true
				}
			}
		}
	}

	for _, p := range tomove {
		ni := (p.i + 1) % len(M)
		M[p.i][p.j] = '.'
		M[ni][p.j] = 'v'
	}

	return change
}

func main() {
	lines := Input("25.txt", "\n", true)
	pf("len %d\n", len(lines))

	M = make([][]byte, len(lines))

	for i, _ := range lines {
		M[i] = []byte(lines[i])
	}

	for i := 0; i < 10000; i++ {
		//pf("AFTER %d\n", i)
		//printmatrix()
		change := step()
		if !change {
			Sol(i + 1)
			break
		}
	}
}

// 338
