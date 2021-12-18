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
	return explodeRec([]*Pair{}, p, 0)
}

func explodeRec(path []*Pair, p *Pair, depth int) (rp *Pair, exploded bool) {
	if depth == 3 {
		if p.a != nil {
			var exploded bool
			if p.a.a != nil {
				exploded = true
				exploding := p.a
				explodeOne(append(path, p), exploding)
				p.a = &Pair{nil, nil, 0}
			} else if p.b.a != nil {
				exploded = true
				exploding := p.b
				explodeOne(append(path, p), exploding)
				p.b = &Pair{nil, nil, 0}
			}
			return p, exploded
		} else {
			return p, false
		}
	} else {
		if p.a != nil {
			var exploded bool
			p.a, exploded = explodeRec(append(path, p), p.a, depth+1)
			if !exploded {
				p.b, exploded = explodeRec(append(path, p), p.b, depth+1)
			}
			return p, exploded
		} else {
			return p, false
		}
	}
}

func explodeOne(path []*Pair, expl *Pair) {
	if expl.a.a != nil {
		panic("depth error")
	}
	if expl.b.a != nil {
		panic("depth error")
	}
	addLeft(path, expl, expl.a.val)
	addRight(path, expl, expl.b.val)
}

func addLeft(path []*Pair, expl *Pair, val int) {
	if len(path) == 0 {
		return // nothing to the left
	}
	parent := path[len(path)-1]
	if parent.a == expl {
		addLeft(path[:len(path)-1], parent, val)
	} else {
		addRightDown(parent.a, val)
	}
}

func addRight(path []*Pair, expl *Pair, val int) {
	if len(path) == 0 {
		return // nothing to the right
	}
	parent := path[len(path)-1]
	if parent.a == expl {
		addLeftDown(parent.b, val)
	} else {
		addRight(path[:len(path)-1], parent, val)
	}
}

func addRightDown(p *Pair, val int) {
	if p.a == nil {
		p.val += val
	} else {
		addRightDown(p.b, val)
	}
}

func addLeftDown(p *Pair, val int) {
	if p.a == nil {
		p.val += val
	} else {
		addLeftDown(p.a, val)
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

func roundcheck(n int) {
	pf("%d %d\n", n/2, (n+1)/2)
}

func main() {
	lines := Input("18.txt", "\n", true)
	pf("len %d\n", len(lines))

	{
		x, _ := parseOne("[[9,1],[1,9]]")
		pf("%d\n", mag(x))
	}

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
