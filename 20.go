package main

import (
	. "./util"
	"fmt"
	"strconv"
	"strings"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

type Point struct {
	i, j int
}

var M = map[Point]bool{}
var Space bool
var alg string

func printmatrix(sp, ep Point) {
	for i := sp.i; i <= ep.i; i++ {
		for j := sp.j; j <= ep.j; j++ {
			if M[Point{i, j}] {
				pf("#")
			} else {
				pf(".")
			}
		}
		pf("\n")
	}
	pf("\n\n")
}

func minmax() (min Point, max Point) {
	first := true
	for p := range M {
		if first {
			min = p
			max = p
			first = false
		}
		if p.i < min.i {
			min.i = p.i
		}
		if p.j < min.j {
			min.j = p.j
		}
		if p.i > max.i {
			max.i = p.i
		}
		if p.j > max.j {
			max.j = p.j
		}
	}
	return
}

const margin = 10

func step() {
	min, max := minmax()

	newM := make(map[Point]bool)
	for i := min.i - margin; i <= max.i+margin; i++ {
		for j := min.j - margin; j <= max.j+margin; j++ {
			np := alg[getgroup(i, j)]
			if np == '#' {
				newM[Point{i, j}] = true
			} else {
				newM[Point{i, j}] = false
			}
		}
	}

	M = newM
	Space = !Space

	for {
		if !removeborder() {
			break
		}
	}

}

func removeborder() bool {
	didsomething := false
	min, max := minmax()

	// remove top
	{
		ok := true
		for j := min.j; j <= max.j; j++ {
			if M[Point{min.i, j}] != Space {
				ok = false
				break
			}
		}
		if ok {
			didsomething = true
			for j := min.j; j <= max.j; j++ {
				delete(M, Point{min.i, j})
			}
		}
	}

	// remove bottom
	{
		ok := true
		for j := min.j; j <= max.j; j++ {
			if M[Point{max.i, j}] != Space {
				ok = false
				break
			}
		}
		if ok {
			didsomething = true
			for j := min.j; j <= max.j; j++ {
				delete(M, Point{max.i, j})
			}
		}
	}

	min, max = minmax()

	// remove left
	{
		ok := true
		for i := min.i; i <= max.i; i++ {
			if M[Point{i, min.j}] != Space {
				ok = false
				break
			}
		}
		if ok {
			didsomething = true
			for i := min.i; i <= max.i; i++ {
				delete(M, Point{i, min.j})
			}
		}
	}

	{
		ok := true
		for i := min.i; i <= max.i; i++ {
			if M[Point{i, max.j}] != Space {
				ok = false
				break
			}
		}
		if ok {
			didsomething = true
			for i := min.i; i <= max.i; i++ {
				delete(M, Point{i, max.j})
			}
		}
	}

	return didsomething
}

func getgroup(ti, tj int) int {
	r := []byte{}
	for i := ti - 1; i <= ti+1; i++ {
		for j := tj - 1; j <= tj+1; j++ {
			v, ok := M[Point{i, j}]
			if ok {
				if v {
					r = append(r, '1')
				} else {
					r = append(r, '0')
				}
			} else {
				if Space {
					r = append(r, '1')
				} else {
					r = append(r, '0')
				}
			}
		}
	}
	n, err := strconv.ParseInt(string(r), 2, 64)
	Must(err)
	return int(n)
}

func count() int {
	r := 0
	for p := range M {
		if M[p] {
			r++
		}
	}
	return r
}

func main() {
	groups := Input("20.txt", "\n\n", true)
	pf("len %d\n", len(groups))

	alg = strings.Replace(groups[0], "\n", "", -1)

	for i, line := range Spac(groups[1], "\n", -1) {
		for j := range line {
			if line[j] == '#' {
				M[Point{i, j}] = true
			}
		}
	}

	printmatrix(minmax())

	//TODO:
	// - make the border removal general, so that it works with both input and example
	// - print both solutions

	for cnt := 0; cnt < 50; cnt++ {
		step()
		//printmatrix(minmax())
	}
	Sol(count())
}
