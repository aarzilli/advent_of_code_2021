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
		if rules[s[i:i+1]] != "" {
			r = append(r, rules[s[i:i+1]])
		}
	}
	r = append(r, s[len(s)-1])
	return r
}

func count(root *Node) (int, map[byte]int) {
	m := map[byte]int{}
	r := 0
	for cur := root; cur != nil; cur = cur.next {
		r++
		m[cur.lt]++
	}
	return r, m
}

var rules = map[string]string{}

func main() {
	groups := Input("14.example.txt", "\n\n", true)
	pf("len %d\n", len(groups))

	var root *Node
	var cur **Node = &root
	for _, ch := range Spac(groups[0], "", -1) {
		*cur = &Node{lt: ch[0]}
		cur = &((*cur).next)
	}
	printlist(root)

	for _, line := range Spac(groups[1], "\n", -1) {
		v := Spac(line, "->", -1)
		rules[v[0]] = v[1]
	}
	pf("%v\n", rules)

	for i := 0; i < 10; i++ {
		pf("%d\n", i)
		step(root)
		//printlist(root)
	}

	_, m := count(root)
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
}
