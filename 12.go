package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

var link = map[string]map[string]bool{}

const showPaths = false

func allpaths(path []string, candouble bool, seen map[string]int) int {
	if path[len(path)-1] == "end" {
		if showPaths {
			pf("%v\n", path)
		}
		return 1
	}
	r := 0
	for nb := range link[path[len(path)-1]] {
		recur := func(newcandouble bool) {
			seen[nb]++
			r += allpaths(append(path, nb), newcandouble, seen)
			seen[nb]--
		}

		if issmall(nb) {
			if seen[nb] == 0 {
				recur(candouble)
			} else {
				if candouble && nb != "start" {
					recur(false)
				}
			}
		} else {
			recur(candouble)
		}
	}
	return r
}

func issmall(n string) bool {
	for i := range n {
		if n[i] < 'a' || n[i] > 'z' {
			return false
		}
	}
	return true
}

func main() {
	lines := Input("12.txt", "\n", true)
	for _, line := range lines {
		v := Spac(line, "-", 2)
		if link[v[0]] == nil {
			link[v[0]] = make(map[string]bool)
		}
		link[v[0]][v[1]] = true
		if link[v[1]] == nil {
			link[v[1]] = make(map[string]bool)
		}
		link[v[1]][v[0]] = true
	}

	Sol(allpaths([]string{"start"}, false, map[string]int{"start": 1}))
	Sol(allpaths([]string{"start"}, true, map[string]int{"start": 1}))
}
