package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

var link = map[string]map[string]bool{}
var count, count2 int

func allpaths(path []string, seen map[string]bool) {
	if path[len(path)-1] == "end" {
		pf("%v\n", path)
		count++
	}
	for nb := range link[path[len(path)-1]] {
		if issmall(nb) && seen[nb] {
			continue
		}
		seen[nb] = true
		allpaths(append(path, nb), seen)
		delete(seen, nb)
	}
}

func allpaths2(path []string, candouble bool, seen map[string]int) {
	if path[len(path)-1] == "end" {
		pf("%v\n", path)
		count2++
		return
	}
	for nb := range link[path[len(path)-1]] {
		canvisit := false
		newcandouble := candouble
		if issmall(nb) {
			if seen[nb] == 0 {
				canvisit = true
			} else {
				if candouble && nb != "start" {
					canvisit = true
					newcandouble = false
				}
			}
		} else {
			canvisit = true
		}
		if canvisit {
			seen[nb]++
			allpaths2(append(path, nb), newcandouble, seen)
			seen[nb]--
		}
	}
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
	pf("len %d\n", len(lines))
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

	/*seen := map[string]bool{ "start": true }
	allpaths([]string{ "start" }, seen)
	Sol(count)*/

	seen := map[string]int{"start": 1}
	allpaths2([]string{"start"}, true, seen)
	Sol(count2)
}
