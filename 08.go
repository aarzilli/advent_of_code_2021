package main

import (
	. "./util"
	"fmt"
	"sort"
	"strings"
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
	b := Spac(w, "", -1)
	sort.Strings(b)
	return strings.Join(b, "")
}

func findCond(words []string, cond func(string) bool) string {
	for i, w := range words {
		if cond(w) {
			words[i] = ""
			return w
		}
	}
	panic("blah")
}

func deduce(words []string) map[string]int {
	m := map[int][]string{}
	for _, w := range words {
		m[len(w)] = append(m[len(w)], sortsegs(w))
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
	six := findCond(m[6], func(segs string) bool { return !containsAll(segs, one) })
	segments2digit[six] = 6

	// find 9 (only remaining 6 segment that contains all segments from 4)
	nine := findCond(m[6], func(segs string) bool { return containsAll(segs, four) })
	segments2digit[nine] = 9

	// 0 is the only remaining 6 segment
	segments2digit[nonstriken(m[6])] = 0

	// find 3 (only 5 segement that contains both segemnts of 1)
	three := findCond(m[5], func(segs string) bool { return containsAll(segs, one) })
	segments2digit[three] = 3

	// find 5 (only remaining 5 segment that contains all segments of 9 that don't belong to 1)
	fivedetector := removesegs(nine, one)
	five := findCond(m[5], func(segs string) bool { return containsAll(segs, fivedetector) })
	segments2digit[five] = 5

	// find 2 (only remaining 5 segment)
	segments2digit[nonstriken(m[5])] = 2

	return segments2digit
}

func main() {
	lines := Input("08.txt", "\n", true)

	r := 0
	part2 := 0
	for _, line := range lines {
		fields := Spac(line, "|", 2)
		signals := Spac(fields[0], " ", -1)
		segments2digit := deduce(signals)
		output := Spac(fields[1], " ", -1)
		n := 0
		for _, w := range output {
			if len(w) == 2 || len(w) == 4 || len(w) == 3 || len(w) == 7 {
				r++
			}

			d := segments2digit[sortsegs(w)]
			n = n*10 + d
		}
		part2 += n
	}
	Sol(r)     // 330
	Sol(part2) // 1010472
}
