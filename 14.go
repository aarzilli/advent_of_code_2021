package main

import (
	. "./util"
	"fmt"
)

type Node struct {
	lt   byte
	next *Node
}

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func printlist(n *Node) {
	if n == nil {
		pf("\n")
		return
	}
	pf("%c", n.lt)
	printlist(n.next)
}

func step(s string) string {
	r := []byte{}
	for i := 0; i < len(s)-1; i++ {
		r = append(r, s[i])
		if rules[s[i:i+2]] != "" {
			r = append(r, rules[s[i:i+2]][0])
		}
	}
	r = append(r, s[len(s)-1])
	return string(r)
}

func count(s string) (int, map[byte]int) {
	m := map[byte]int{}
	r := 0
	for i := range s {
		r++
		m[s[i]]++
	}
	return r, m
}

var rules = map[string]string{}

func recexpand(s string, n int, m map[string]int) string {
	if n == 0 {
		m[string(s[0])]++
		m[string(s[1])]++
		return s
	}
	if rules[s] == "" {
		m[string(s[0])]++
		m[string(s[1])]++
		return s
	}
	a := recexpand(string(s[0])+rules[s], n-1, m)
	b := recexpand(rules[s]+string(s[1]), n-1, m)
	_, _ = a, b
	m[rules[s]]--
	return a + b[1:]
}

type memopair struct {
	s string
	n int
}

var memoized = map[memopair]map[string]int{}

func copymap(m map[string]int) map[string]int {
	r := map[string]int{}
	for k := range m {
		r[k] = m[k]
	}
	return r
}

func recexpandmemo(s string, n int) map[string]int {
	if memoized[memopair{s, n}] != nil {
		return copymap(memoized[memopair{s, n}])
	}
	if n == 0 || rules[s] == "" {
		m := map[string]int{}
		m[string(s[0])]++
		m[string(s[1])]++
		memoized[memopair{s, n}] = m
		return m
	}
	a := recexpandmemo(string(s[0])+rules[s], n-1)
	b := recexpandmemo(rules[s]+string(s[1]), n-1)

	m := map[string]int{}

	for k := range a {
		m[k] += a[k]
	}
	for k := range b {
		m[k] += b[k]
	}
	m[rules[s]]--
	memoized[memopair{s, n}] = m
	return m
}

func main() {
	groups := Input("14.txt", "\n\n", true)
	pf("len %d\n", len(groups))

	start := groups[0]

	for _, line := range Spac(groups[1], "\n", -1) {
		v := Spac(line, "->", -1)
		rules[v[0]] = v[1]
	}
	pf("%v\n", rules)

	cur := start
	for i := 0; i < 10; i++ {
		cur = step(cur)
	}

	_, m := count(cur)
	minel, maxel := byte(0), byte(0)
	for el := range m {
		if minel == 0 {
			minel = el
		}
		if maxel == 0 {
			maxel = el
		}
		if m[minel] > m[el] {
			minel = el
		}
		if m[maxel] < m[el] {
			maxel = el
		}
	}

	pf("%c %d %c %d\n", minel, m[minel], maxel, m[maxel])

	Sol(m[maxel] - m[minel])

	{
		N := 40
		m := map[string]int{}
		for i := 0; i < len(start)-1; i++ {
			m2 := recexpandmemo(start[i:i+2], N)
			for k := range m2 {
				m[k] += m2[k]
			}
			m[string(start[i+1])]--
		}
		// TODO: considerare l'ultima lettera dell'input qui
		m[string(start[len(start)-1])]++
		for k := range m {
			pf("%s %d\n", k, m[k])
		}

		minel, maxel := "", ""
		for el := range m {
			if minel == "" {
				minel = el
			}
			if maxel == "" {
				maxel = el
			}
			if m[minel] > m[el] {
				minel = el
			}
			if m[maxel] < m[el] {
				maxel = el
			}
		}

		pf("%s %d %s %d\n", minel, m[minel], maxel, m[maxel])

		Sol(m[maxel] - m[minel])
	}
}

// NBCC NBBB CBHCB

// N = 2
// B = 6
// C = 4
// H = 1
