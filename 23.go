package main

import (
	. "./util"
	"container/heap"
	"fmt"
	"sort"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

var energy = [4 * perType]int{1, 1, 1, 1, 10, 10, 10, 10, 100, 100, 100, 100, 1000, 1000, 1000, 1000}

const perType = 4

//const perType = 2
//var energy = [8]int{ 1, 1, 10, 10, 100, 100, 1000, 1000 }

type Point struct {
	i, j int
}

type State struct {
	pos [4 * perType]Point
}

type NormState struct {
	pos [4 * perType]Point
}

var M [][]byte

func convert(lines []string) (State, [][]byte) {
	r := State{}
	M := make([][]byte, 0, len(lines))
	for i := range lines {
		M = append(M, []byte(lines[i]))
	}
	for i := range M {
		for j := range M[i] {
			if M[i][j] != '.' && M[i][j] != '#' && M[i][j] != ',' {
				idx := (M[i][j] - 'A') * perType
				for r.pos[idx] != (Point{0, 0}) {
					idx++
				}
				r.pos[idx] = Point{i, j}
				pf("%d %d %c %d\n", i, j, M[i][j], idx)
				M[i][j] = '.'
			}
		}
	}
	return r, M
}

func printState(s State) {
	pf("%v\n", s)
	M2 := make([][]byte, len(M))
	for i := range M {
		M2[i] = append(M2[i], M[i]...)
	}

	for idx := range s.pos {
		M2[s.pos[idx].i][s.pos[idx].j] = byte((idx / perType) + 'A')
	}
	for i := range M2 {
		pf("%s\n", string(M2[i]))
	}
	pf("\n")
}

type fringeNode struct {
	s    NormState
	dist int
}

type fringeHeap struct {
	v []fringeNode
	m map[NormState]int
}

func (h *fringeHeap) Push(x interface{}) {
	el := x.(fringeNode)
	h.v = append(h.v, el)
	h.m[el.s] = el.dist
}

func (h *fringeHeap) Pop() interface{} {
	r := h.v[len(h.v)-1]
	h.v = h.v[:len(h.v)-1]
	delete(h.m, r.s)
	return r
}

func (h *fringeHeap) Len() int {
	return len(h.v)
}

func (h *fringeHeap) Less(i, j int) bool {
	return h.v[i].dist < h.v[j].dist
}

func (h *fringeHeap) Swap(i, j int) {
	h.v[i], h.v[j] = h.v[j], h.v[i]
}

func search(start State) {
	var fringe fringeHeap
	fringe.m = make(map[NormState]int)
	heap.Init(&fringe)
	heap.Push(&fringe, fringeNode{normalize(start), 0})
	seen := map[NormState]bool{} // nodes already visited (we know the minimum distance of those)

	cnt := 0

	for fringe.Len() > 0 {
	popper:
		curnode := heap.Pop(&fringe).(fringeNode)
		cur := curnode.s
		if seen[cur] {
			goto popper
		}

		if cnt%1000 == 0 {
			fmt.Printf("fringe %d (min dist %d)\n", fringe.Len(), curnode.dist)
		}
		cnt++

		if isend(cur) {
			fmt.Printf("FOUND: %v %d\n", cur, curnode.dist)
			return
		}

		//pf("CURRENT STATE:\n")
		//printState(State(cur))

		distcur := curnode.dist
		seen[cur] = true

		maybeadd0 := func(nb State, e int) {
			if seen[normalize(nb)] {
				return
			}
			if !valid(nb) {
				return
			}
			//pf("Neighbour:\n")
			//printState(nb)
			d, ok := fringe.m[normalize(nb)]
			if !ok || distcur+e < d {
				heap.Push(&fringe, fringeNode{normalize(nb), distcur + e})
			}
		}
		_ = maybeadd0

		maybeadd := func(idx int, np Point, steps int) {
			nb := State(cur)
			nb.pos[idx] = np
			maybeadd0(nb, energy[idx]*steps)
		}

		// try to add all possible neighbors

		for idx := range cur.pos {
			/*if idx == 2 || idx == 3 || idx == 6 || idx == 7 || idx == 10 || idx == 11 || idx == 14 || idx == 15 {
				continue
			}*/
			if cur.pos[idx].i > 1 { // in a room
				// going into hallway
				exitsteps := cur.pos[idx].i - 1
				if canmovetoallway(cur, idx) {
					for j := cur.pos[idx].j; j >= 0; j-- {
						if occupied(cur, Point{1, j}) >= 0 {
							break
						}
						maybeadd(idx, Point{1, j}, exitsteps+Abs(j-cur.pos[idx].j))
					}
					for j := cur.pos[idx].j; j < len(M[0]); j++ {
						if occupied(cur, Point{1, j}) >= 0 {
							break
						}
						maybeadd(idx, Point{1, j}, exitsteps+Abs(j-cur.pos[idx].j))
					}
				}
			} else {
				// see if path to room is clear and go there
				destj := (idx/perType)*2 + 3

				roomids := inroom(cur, destj)

				ok := true

				for _, idx2 := range roomids {
					if !correctlane(idx2, destj) {
						ok = false
						break
					}
				}

				if ok {
					dir := 1
					if destj < cur.pos[idx].j {
						dir = -1
					}
					j := cur.pos[idx].j
					ok := true
					for j != destj {
						j += dir
						if occupied(cur, Point{1, j}) >= 0 {
							ok = false
							break
						}
					}
					if ok {
						hcost := Abs(destj - cur.pos[idx].j)
						pushin_nb, vcost := pushin(cur, idx, Point{2, destj})
						_, _ = pushin_nb, vcost
						maybeadd0(pushin_nb, vcost+(hcost*energy[idx]))
					}
				}
			}
		}
	}
}

func canmovetoallway(cur NormState, idx int) bool {
	for i := cur.pos[idx].i - 1; i > 1; i-- {
		if occupied(cur, Point{i, cur.pos[idx].j}) >= 0 {
			return false
		}
	}
	return true
	/*exitsteps := cur.pos[idx].i - 1
	return exitsteps == 1 || occupied(cur, Point{ cur.pos[idx].i-1, cur.pos[idx].j }) < 0*/
}

func inroom(cur NormState, j int) []int {
	r := []int{}
	for i := 2; i <= 5; i++ {
		idx := occupied(cur, Point{i, j})
		if idx >= 0 {
			r = append(r, idx)
		}
	}
	return r
}

func valid(s State) bool {
	for _, p := range s.pos {
		switch M[p.i][p.j] {
		case '#':
			return false
		case '.':
			// ok
		case ',':
			return false
		default:
			panic("blah")
		}
	}
	return true
}

func pushin(cur NormState, idx int, p Point) (State, int) {
	if occupied(cur, p) < 0 {
		nb := State(cur)
		nb.pos[idx] = p
		return nb, energy[idx]
	}

	nb, cost := pushin(cur, occupied(cur, p), Point{p.i + 1, p.j})
	nb.pos[idx] = p
	return nb, cost + energy[idx]
}

func occupied(s NormState, p Point) int {
	for idx := range s.pos {
		if s.pos[idx] == p {
			return idx
		}
	}
	return -1
}

func correctlane(idx, j int) bool {
	if j != (idx/perType)*2+3 {
		return false
	}
	return true
}

func isend(s NormState) bool {
	for idx := range s.pos {
		if s.pos[idx].i <= 1 {
			return false
		}
		if !correctlane(idx, s.pos[idx].j) {
			return false
		}
	}
	return true
}

func ishalfpoint(s NormState) bool {
	for idx := len(s.pos) - 4; idx < len(s.pos); idx++ {
		if s.pos[idx].i <= 1 {
			return false
		}
		if !correctlane(idx, s.pos[idx].j) {
			return false
		}
	}
	return true
}

func pointLess(a, b Point) bool {
	if a.i == b.i {
		return a.j < b.j
	}
	return a.i < b.i
}

func normalize(cur State) NormState {
	r := cur
	for i := 0; i < len(r.pos); i += perType {
		sort.Slice(r.pos[i:i+perType], func(ii, jj int) bool {
			return pointLess(r.pos[i+ii], r.pos[i+jj])
		})
	}
	//if cur != r {
	//pf("%v ->\n%v\n", cur, r)
	//}
	return NormState(r)
}

func main() {
	lines := Input("23.part2.txt", "\n", true)
	pf("len %d\n", len(lines))
	var start State
	start, M = convert(lines)
	pf("%v\n", start)
	pf("%v\n", isend(NormState(start)))
	pf("%v\n", M)
	search(start)
}
