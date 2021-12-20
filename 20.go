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

var M = map[Point]byte{}
var Space byte = '.'
var alg string

func printmatrix(sp, ep Point) {
	for i := sp.i; i <= ep.i; i++ {
		for j := sp.j; j <= ep.j; j++ {
			if ch, ok := M[Point{i, j}]; ok {
				pf("%c", ch)
			} else {
				pf("%c", Space)
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

func step() {
	const margin = 1

	min, max := minmax()

	var newSpace byte
	if Space == '#' {
		newSpace = alg[len(alg)-1]
	} else {
		newSpace = alg[0]
	}

	newM := make(map[Point]byte)
	for i := min.i - margin; i <= max.i+margin; i++ {
		for j := min.j - margin; j <= max.j+margin; j++ {
			np := alg[getgroup(i, j)]
			if np != newSpace {
				newM[Point{i, j}] = np
			}
		}
	}

	M = newM
	Space = newSpace
}

func getgroup(ti, tj int) int {
	r := []byte{}
	for i := ti - 1; i <= ti+1; i++ {
		for j := tj - 1; j <= tj+1; j++ {
			v, ok := M[Point{i, j}]
			if ok {
				r = append(r, v)
			} else {
				r = append(r, Space)
			}
		}
	}

	n, err := strconv.ParseInt(strings.Replace(strings.Replace(string(r), "#", "1", -1), ".", "0", -1), 2, 64)
	Must(err)
	return int(n)
}

func count() int {
	if Space == '#' {
		panic("infinite")
	}
	r := 0
	for p := range M {
		if M[p] == '#' {
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
			M[Point{i, j}] = line[j]
		}
	}

	//TODO:
	// - make the border removal general, so that it works with both input and example

	for cnt := 0; cnt < 50; cnt++ {
		step()
		//printmatrix(minmax())
		if cnt == 1 {
			Expect(5259)
			Sol(count())
		}
	}
	Expect(15287)
	Sol(count())
}
