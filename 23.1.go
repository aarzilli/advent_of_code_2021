package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

/*var energy = map[string]int{
	"A": 1,
	"B": 10,
	"C": 100,
	"D": 1000,
}*/

var energy = [8]int{1, 1, 10, 10, 100, 100, 1000, 1000}

var endstate = `#############
#...........#
###A#B#C#D###
###A#B#C#D###
#############
`

type Point struct {
	i, j int
}

type State struct {
	pos [8]Point
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
				idx := (M[i][j] - 'A') * 2
				if r.pos[idx] != (Point{0, 0}) {
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

// finds the closest node in the fringe, lastmin is an optimization, if we find a node that is at that distance we return it immediately (there can be nothing that's closer)
func minimum(fringe map[State]int, lastmin int) State {
	var mink State
	first := true
	for k, d := range fringe {
		if first {
			mink = k
			first = false
		}
		if d == lastmin {
			return k
		}
		if d < fringe[mink] {
			mink = k
		}
	}
	return mink
}

type fringeNode struct {
	s    State
	dist int
}

/*func heur(s State) int {
	tot := 0
	for idx := range s.pos {
		if !correctlane(idx, s.pos[idx].j) {
			// going up
			tot += energy[idx]*(s.pos[idx].i-2)
			// moving to correct lane
			tot += Abs(s.pos[idx].j - ((idx/2)*2 + 3)) * energy[idx]
			// going down
			tot += energy[idx]
		} else {
			if s.pos[idx].i == 2 {
				tot += energy[idx]
			}
		}

	}
	return tot
}*/

func heur(s State) int {
	return 0
}

func search(start, end State) {
	fringe := map[State]int{start: 0}   // nodes discovered but not visited (start at node 0 with distance 0)
	seen := map[State]bool{start: true} // nodes already visited (we know the minimum distance of those)

	lastmin := 0

	cnt := 0

	for len(fringe) > 0 {
		cur := minimum(fringe, lastmin)

		if cnt%1000 == 0 {
			fmt.Printf("fringe %d (min dist %d)\n", len(fringe), fringe[cur])
		}
		cnt++

		if isend(cur, end) {
			fmt.Printf("FOUND: %v %d\n", cur, fringe[cur])
			return
		}

		distcur := fringe[cur]
		lastmin = distcur
		delete(fringe, cur)
		seen[cur] = true

		maybeadd := func(idx int, np Point, steps int) {
			nb := cur
			nb.pos[idx] = np

			if seen[nb] {
				return
			}

			switch M[np.i][np.j] {
			case '#':
				return
			case '.':
				// ok
			case ',':
				return
			default:
				panic("blah")
			}

			if occupied(cur, np) >= 0 {
				return
			}

			d, ok := fringe[nb]
			if !ok || distcur+(energy[idx]*steps) < d {
				fringe[nb] = distcur + (energy[idx] * steps)
			}
		}

		maybeadd2 := func(idx1 int, np1 Point, idx2 int, np2 Point, steps int) {
			nb := cur
			nb.pos[idx1] = np1
			nb.pos[idx2] = np2
			if seen[nb] {
				return
			}
			d, ok := fringe[nb]
			if !ok || distcur+(energy[idx1]*steps) < d {
				fringe[nb] = distcur + (energy[idx1] * steps)
			}
		}

		// try to add all possible neighbors

		for idx := range cur.pos {
			if cur.pos[idx].i > 1 { // in a room
				// going into hallway
				exitsteps := cur.pos[idx].i - 1
				if exitsteps == 1 || occupied(cur, Point{cur.pos[idx].i - 1, cur.pos[idx].j}) < 0 {
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
				destj := (idx/2)*2 + 3

				first := occupied(cur, Point{2, destj})
				second := occupied(cur, Point{3, destj})
				secondok := (second == -1) || correctlane(second, destj)

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
					if first == -1 && secondok {
						maybeadd(idx, Point{2, destj}, 1+Abs(destj-cur.pos[idx].j))
					} else if (second == -1) && correctlane(first, destj) {
						maybeadd2(idx, Point{2, destj}, first, Point{3, destj}, 2+Abs(destj-cur.pos[idx].j))
					}
				}
			}
		}
	}
}

func occupied(s State, p Point) int {
	for idx := range s.pos {
		if s.pos[idx] == p {
			return idx
		}
	}
	return -1
}

func correctlane(idx, j int) bool {
	if j != (idx/2)*2+3 {
		return false
	}
	return true
}

func isend(s State, end State) bool {
	for i := 0; i < len(s.pos); i += 2 {
		if s.pos[i+0] != end.pos[i+0] && s.pos[i+0] != end.pos[i+1] {
			return false
		}
		if s.pos[i+1] != end.pos[i+0] && s.pos[i+1] != end.pos[i+1] {
			return false
		}
	}
	return true
}

func main() {
	lines := Input("23.txt", "\n", true)
	pf("len %d\n", len(lines))
	var start State
	start, M = convert(lines)
	pf("%v\n", start)
	end, _ := convert(Spac(endstate, "\n", -1))
	pf("%v\n", end)
	pf("%v\n", M)
	search(start, end)
}
