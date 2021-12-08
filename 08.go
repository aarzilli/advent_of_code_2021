package main

import (
	. "./util"
	"fmt"
	"sort"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func only(m map[int][]string, n int) string {
	if len(m[n]) != 1 {
		panic("blah")
	}
	r := m[n][0]
	delete(m, n)
	return r
}

func containsAll(segs string, needle string) bool {
	for _, x := range needle {
		found := false
		for _, y := range segs {
			if x == y {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func nonstriken(words []string) string {
	for _, word := range words {
		if word != "" {
			return word
		}
	}
	panic("blah")
}

func removesegs(source, needle string) string {
	r := []rune{}
	for _, sc := range source {
		found := false
		for _, nc := range needle {
			if sc == nc {
				found = true
				break
			}
		}
		if !found {
			r = append(r, sc)
		}
	}
	return string(r)
}

func sortsegs(w string) string {
	b := []byte(w)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}

func deduce(words []string) map[string]int {
	m := map[int][]string{}
	for _, w := range words {
		m[len(w)] = append(m[len(w)], w)
	}
	segments2digit := map[string]int{}

	// find normals
	one := only(m, 2)
	segments2digit[one] = 1
	four := only(m, 4)
	segments2digit[four] = 4
	segments2digit[only(m, 3)] = 7
	segments2digit[only(m, 7)] = 8

	// find 6 (only 6 segment that doesn't contain both segments of 1)
	ok := false
	for i, segs := range m[6] {
		if !containsAll(segs, one) {
			segments2digit[segs] = 6
			m[6][i] = ""
			ok = true
			break
		}
	}
	if !ok {
		panic("could not find 6")
	}

	var nine string

	// find 9 (only remaining 6 segment that contains all segments from 4)
	ok = false
	for i, segs := range m[6] {
		if segs == "" {
			continue
		}
		if containsAll(segs, four) {
			segments2digit[segs] = 9
			nine = segs
			ok = true
			m[6][i] = ""
			break
		}
	}
	if !ok {
		panic("could not find 9")
	}

	// find 3 (only 5 segement that contains both segemnts of 1)
	ok = false
	for i, segs := range m[5] {
		if containsAll(segs, one) {
			segments2digit[segs] = 3
			ok = true
			m[5][i] = ""
			break
		}
	}
	if !ok {
		panic("could not find 3")
	}

	// find 5 (only remaining 5 segment that contains all segments of 9 that don't belong to 1)
	fivedetector := removesegs(nine, one)
	ok = false
	for i, segs := range m[5] {
		if containsAll(segs, fivedetector) {
			segments2digit[segs] = 5
			ok = true
			m[5][i] = ""
			break
		}
	}
	if !ok {
		panic("could not find 5")
	}

	// find 2 (only remaining 5 segment)
	segments2digit[nonstriken(m[5])] = 2

	// 0 is the only remaining 6 segment
	segments2digit[nonstriken(m[6])] = 0

	//pf("segments2digit: %v\n", segments2digit)
	//pf("%v\n", m)

	r := map[string]int{}
	for k, v := range segments2digit {
		r[sortsegs(k)] = v
	}
	return r
}

func main() {
	lines := Input("08.txt", "\n", true)
	pf("len %d\n", len(lines))

	r := 0
	part2 := 0
	for _, line := range lines {
		fields := Spac(line, "|", 2)
		//pf("%q\n", fields)
		signals := Spac(fields[0], " ", -1)
		segments2digit := deduce(signals)
		pf("segment map: %v\n", segments2digit)
		output := Spac(fields[1], " ", -1)
		n := 0
		for _, w := range output {
			if len(w) == 2 || len(w) == 4 || len(w) == 3 || len(w) == 7 {
				r++
			}

			d := segments2digit[sortsegs(w)]
			n = n*10 + d
		}
		pf("%q %d\n", output, n)
		part2 += n
	}
	Sol(r)
	Sol(part2)
}
