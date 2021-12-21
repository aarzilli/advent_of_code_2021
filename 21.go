package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

var die = 1
var rolls = 0

func roll() int {
	if die+3 <= 100 {
		r := 3*die + 3
		die += 3
		rolls += 3
		return r
	}
	r := 0
	for i := 0; i < 3; i++ {
		r += die
		die++
		rolls++
		if die == 101 {
			die = 1
		}
	}
	return r
}

type Player struct {
	id    int
	pos   int
	score int
}

func play(p *Player) {
	n := roll()
	p.pos += n
	for p.pos > 10 {
		p.pos -= 10
	}
	p.score += p.pos
	//pf("player %d rolls %d moves to %d score %d\n", p.id, n, p.pos, p.score)
}

var outcomes map[int]int
var totwin [2]int

func play2(univ map[[2]Player]int, i int) map[[2]Player]int {
	pid := i % 2
	r := map[[2]Player]int{}
	for u, cnt := range univ {
		for outcome, count := range outcomes {
			ru := u
			ru[pid].pos += outcome
			for ru[pid].pos > 10 {
				ru[pid].pos -= 10
			}
			ru[pid].score += ru[pid].pos
			if ru[pid].score >= 21 {
				totwin[pid] += cnt * count
			} else {
				r[ru] += cnt * count
			}
		}
	}
	return r
}

func main() {
	lines := Input("21.txt", "\n", true)
	pf("len %d\n", len(lines))
	player := make([]Player, 2)
	player[0].id = 1
	player[0].pos = Atoi(Spac(lines[0], ":", -1)[1])
	player[1].id = 2
	player[1].pos = Atoi(Spac(lines[1], ":", -1)[1])

	for {
		play(&player[0])
		if player[0].score >= 1000 {
			Sol(player[1].score * rolls)
			break
		}
		play(&player[1])
		if player[1].score >= 1000 {
			Sol(player[0].score * rolls)
			break
		}
	}

	outcomes = map[int]int{}

	for _, a := range []int{1, 2, 3} {
		for _, b := range []int{1, 2, 3} {
			for _, c := range []int{1, 2, 3} {
				outcomes[a+b+c]++
			}
		}
	}

	pf("%v\n", outcomes)

	univ := map[[2]Player]int{
		[2]Player{
			Player{1, Atoi(Spac(lines[0], ":", -1)[1]), 0},
			Player{2, Atoi(Spac(lines[1], ":", -1)[1]), 0},
		}: 1,
	}

	pf("%v\n", univ)

	for i := 0; i < 1000; i++ {
		univ = play2(univ, i)
		pf("%d %v\n", len(univ), totwin)
		if len(univ) == 0 {
			pf("DONE\n")
			break
		}
	}

	pf("%v\n", totwin)
	if totwin[0] > totwin[1] {
		Sol(totwin[0])
	} else {
		Sol(totwin[1])
	}
}
