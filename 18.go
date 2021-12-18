package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

type Pair struct {
	a, b *Pair
	val  int
}

func parseOne(line string) (*Pair, string) {
	if line[0] == '[' {
		line = line[1:]
		a, line := parseOne(line)
		if line[0] != ',' {
			panic("syntax error")
		}
		line = line[1:]
		b, line := parseOne(line)
		if line[0] != ']' {
			panic("syntax error")
		}
		line = line[1:]
		return &Pair{a, b, 0}, line
	} else {
		val := int(line[0] - '0')
		return &Pair{nil, nil, val}, line[1:]
	}
}

func printPair(p *Pair) {
	if p.a != nil && p.b != nil {
		pf("[")
		printPair(p.a)
		pf(",")
		printPair(p.b)
		pf("]")
	} else {
		pf("%d", p.val)
	}
}

func add(acc, b *Pair) *Pair {
	acc = &Pair{acc, b, 0}
	acc = reduce(acc)
	return acc
}

func reduce(p *Pair) *Pair {
	for {
		var exploded, splat bool
		p, exploded = explode(p)
		if exploded {
			//pf("exploded\n")
		}
		if !exploded {
			p, splat = split(p)
			if splat {
				//pf("splat\n")
			}
		}
		if !exploded && !splat {
			break
		}
	}
	return p
}

func explode(p *Pair) (*Pair, bool) {
	return explodeRec(p, p, 0)
}

func explodeRec(root *Pair, p *Pair, depth int) (rp *Pair, exploded bool) {
	if depth == 3 {
		if p.a != nil {
			var exploded bool
			if p.a.a != nil {
				exploded = true
				exploding := p.a
				explodeOne(root, exploding)
				p.a = &Pair{nil, nil, 0}
			} else if p.b.a != nil {
				exploded = true
				exploding := p.b
				explodeOne(root, exploding)
				p.b = &Pair{nil, nil, 0}
			}
			return p, exploded
		} else {
			return p, false
		}
	} else {
		if p.a != nil {
			var exploded bool
			p.a, exploded = explodeRec(root, p.a, depth+1)
			if !exploded {
				p.b, exploded = explodeRec(root, p.b, depth+1)
			}
			return p, exploded
		} else {
			return p, false
		}
	}
}

func explodeOne(root *Pair, expl *Pair) {
	if expl.a.a != nil {
		panic("depth error")
	}
	if expl.b.a != nil {
		panic("depth error")
	}
	var lvs []*Pair
	leaves(root, &lvs)

	for i := range lvs {
		if lvs[i] == expl.a {
			if i-1 >= 0 {
				lvs[i-1].val += expl.a.val
			}
			if i+2 < len(lvs) {
				lvs[i+2].val += expl.b.val
			}
			break
		}
	}
}

func leaves(p *Pair, lvs *[]*Pair) {
	if p.a != nil {
		leaves(p.a, lvs)
		leaves(p.b, lvs)
	} else {
		*lvs = append(*lvs, p)
	}
}

func split(p *Pair) (*Pair, bool) {
	if p.a != nil {
		var splat bool
		p.a, splat = split(p.a)
		if !splat {
			p.b, splat = split(p.b)
		}
		return p, splat
	} else {
		if p.val >= 10 {
			return &Pair{&Pair{nil, nil, p.val / 2}, &Pair{nil, nil, (p.val + 1) / 2}, 0}, true
		} else {
			return p, false
		}
	}
}

func mag(p *Pair) int {
	if p.a == nil {
		return p.val
	} else {
		return 3*mag(p.a) + 2*mag(p.b)
	}
}

func main() {
	lines := Input("18.txt", "\n", true)

	var acc *Pair
	for _, line := range lines {
		x, _ := parseOne(line)
		//printPair(x)
		//pf("\n")
		if acc == nil {
			acc = x
		} else {
			acc = add(acc, x)
		}
	}
	printPair(acc)
	pf("\n")
	Sol(mag(acc))

	var maxm = 0
	for i := range lines {
		for j := range lines {
			if i == j {
				continue
			}
			a, _ := parseOne(lines[i])
			b, _ := parseOne(lines[j])

			a = add(a, b)
			m := mag(a)
			if m > maxm {
				maxm = m
			}
		}
	}
	Sol(maxm)
}
