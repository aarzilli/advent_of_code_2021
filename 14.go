package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

var rules = map[string]string{}

type memopair struct {
	s string
	n int
}

var memoized = map[memopair]map[string]int{}

func recexpandmemo(s string, n int) map[string]int {
	if memoized[memopair{s, n}] != nil {
		return memoized[memopair{s, n}]
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

func recexpandmemoFull(s string, n int) map[string]int {
	m := map[string]int{}
	for i := 0; i < len(s)-1; i++ {
		m2 := recexpandmemo(s[i:i+2], n)
		for k := range m2 {
			m[k] += m2[k]
		}
		m[string(s[i+1])]--
	}
	m[string(s[len(s)-1])]++
	return m
}

func minmaxmap(m map[string]int) (minel, maxel string) {
	minel, maxel = "", ""
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

	return minel, maxel
}

func main() {
	groups := Input("14.txt", "\n\n", true)

	start := groups[0]

	for _, line := range Spac(groups[1], "\n", -1) {
		v := Spac(line, "->", -1)
		rules[v[0]] = v[1]
	}

	{
		m := recexpandmemoFull(start, 10)
		minel, maxel := minmaxmap(m)

		Sol(m[maxel] - m[minel])
	}

	{
		m := recexpandmemoFull(start, 40)
		minel, maxel := minmaxmap(m)

		Sol(m[maxel] - m[minel])
	}
}
