package main

import (
	. "./util"
	"fmt"
	"strconv"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func oxygencrit(x string, bit int, m map[byte]int) bool {
	var tgt byte
	if m['0'] > m['1'] {
		tgt = '0'
	} else {
		tgt = '1'
	}
	return x[bit] == tgt
}

func co2scrubrate(x string, bit int, m map[byte]int) bool {
	var tgt byte
	if m['0'] <= m['1'] {
		tgt = '0'
	} else {
		tgt = '1'
	}
	return x[bit] == tgt
}

func selector(in []string, f func(x string, bit int, m map[byte]int) bool) string {
	n := len(in[0])

	for bit := 0; bit < n; bit++ {
		m := map[byte]int{}
		for _, x := range in {
			if x != "" {
				m[x[bit]]++
			}
		}
		if len(m) == 1 {
			break
		}
		for i, x := range in {
			if x != "" {
				if !f(x, bit, m) {
					in[i] = ""
				}
			}
		}
		//pf("%v\n", in)
	}
	for i := range in {
		if in[i] != "" {
			return in[i]
		}
	}
	panic("blah")
}

func copyof(in []string) []string {
	r := make([]string, len(in))
	copy(r, in)
	return r
}

func main() {
	lines := Input("03.txt", "\n", true)
	pf("len %d\n", len(lines))
	gamma := make([]byte, len(lines[0]))
	for i := range gamma {
		m := map[byte]int{}
		for _, line := range lines {
			m[line[i]]++
		}
		if m['0'] > m['1'] {
			gamma[i] = '0'
		} else {
			gamma[i] = '1'
		}
	}

	epsilon := make([]byte, len(gamma))
	for i := range gamma {
		if gamma[i] == '0' {
			epsilon[i] = '1'
		} else {
			epsilon[i] = '0'
		}
	}

	pf("%s %s\n", gamma, epsilon)
	gamman, _ := strconv.ParseInt(string(gamma), 2, 64)
	epsilonn, _ := strconv.ParseInt(string(epsilon), 2, 64)
	Sol(gamman, epsilonn, gamman*epsilonn)

	oc, _ := strconv.ParseInt(selector(copyof(lines), oxygencrit), 2, 64)
	sr, _ := strconv.ParseInt(selector(copyof(lines), co2scrubrate), 2, 64)
	Sol(oc, sr, oc*sr)
}
