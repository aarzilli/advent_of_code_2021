package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

type point struct{ x, y, xspeed, yspeed int }

func intarget(p point, tgt []int) bool {
	return p.y >= tgt[2] && p.y <= tgt[3] && p.x >= tgt[0] && p.x <= tgt[1]
}

func pasttarget(p point, tgt []int) bool {
	if p.xspeed > 0 {
		if p.y < tgt[2] && p.yspeed < 0 {
			//pf("A\n")
			return true
		}
		if p.x > tgt[1] {
			//pf("B\n")
			return true
		}
	}
	if p.xspeed == 0 {
		if p.x < tgt[0] || p.x > tgt[1] {
			//pf("C\n")
			return true
		}
		if p.y < tgt[2] && p.yspeed < 0 {
			//pf("D\n")
			return true
		}
	}
	return false
}

func doesland(tgt []int, xspeed, yspeed int) (ok bool, maxheight int) {
	var p point
	p.xspeed = xspeed
	p.yspeed = yspeed

	for i := 0; i < 10000; i++ {
		//pf("pos=%v\n", p)
		if p.y > maxheight {
			maxheight = p.y
		}
		if intarget(p, tgt) {
			//pf("in target\n")
			return true, maxheight
		}
		if pasttarget(p, tgt) {
			//pf("past target %v\n", p)
			return false, maxheight
		}
		p.x += p.xspeed
		p.y += p.yspeed
		if p.xspeed > 0 {
			p.xspeed--
		}
		p.yspeed--
	}
	panic("blowup")
}

func maxy(yspeed int) int {
	return (yspeed * (yspeed + 1)) / 2
}

func main() {
	lines := Input("17.txt", "\n", true)
	pf("len %d\n", len(lines))
	tgt := Vatoi(Getnums(lines[0], true, false))

	pf("%v\n", tgt)

	var p point
	p.xspeed = 17
	p.yspeed = -4

	//doesland(tgt, 6, -100)

	v := []int{23, -10, 25, -9, 27, -5, 29, -6, 22, -6, 21, -7, 9, 0, 27, -7, 24, -5,
		25, -7, 26, -6, 25, -5, 6, 8, 11, -2, 20, -5, 29, -10, 6, 3, 28, -7,
		8, 0, 30, -6, 29, -8, 20, -10, 6, 7, 6, 4, 6, 1, 14, -4, 21, -6,
		26, -10, 7, -1, 7, 7, 8, -1, 21, -9, 6, 2, 20, -7, 30, -10, 14, -3,
		20, -8, 13, -2, 7, 3, 28, -8, 29, -9, 15, -3, 22, -5, 26, -8, 25, -8,
		25, -6, 15, -4, 9, -2, 15, -2, 12, -2, 28, -9, 12, -3, 24, -6, 23, -7,
		25, -10, 7, 8, 11, -3, 26, -7, 7, 1, 23, -9, 6, 0, 22, -10, 27, -6,
		8, 1, 22, -8, 13, -4, 7, 6, 28, -6, 11, -4, 12, -4, 26, -9, 7, 4,
		24, -10, 23, -8, 30, -8, 7, 0, 9, -1, 10, -1, 26, -5, 22, -9, 6, 5,
		7, 5, 23, -6, 28, -10, 10, -2, 11, -1, 20, -9, 14, -2, 29, -7, 13, -3,
		23, -5, 24, -8, 27, -9, 30, -7, 28, -5, 21, -10, 7, 9, 6, 6, 21, -5,
		27, -10, 7, 2, 30, -9, 21, -8, 22, -7, 24, -9, 20, -6, 6, 9, 29, -5,
		8, -2, 27, -8, 30, -5, 24, -7}
	_ = v

	/*for i := 0; i < len(v); i += 2 {
		ok, _ := doesland(tgt, v[i], v[i+1])
		if !ok {
			pf("error: %d %d\n", v[i], v[i+1])
		}
	}*/

	var my int
	cnt := 0
	for yspeed := -1000; yspeed < 1000; yspeed++ {
		for xspeed := 0; xspeed < 1000; xspeed++ {
			//pf("%d %d\n", xspeed, yspeed)
			ok, curmy := doesland(tgt, xspeed, yspeed)
			if ok && curmy > my {
				my = curmy
			}
			if ok {
				//pf("%d,%d\n", xspeed, yspeed)
				cnt++
			}
		}
	}
	Sol(my)
	Sol(cnt)
}
