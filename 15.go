package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

var M = [][]int{}

type state struct {
	i, j int
}

// finds the closest node in the fringe, lastmin is an optimization, if we find a node that is at that distance we return it immediately (there can be nothing that's closer)
func minimum(fringe map[state]int, lastmin int) state {
	var mink state
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

func search() {
	fringe := map[state]int{state{0, 0}: 0}   // nodes discovered but not visited (start at node 0 with distance 0)
	seen := map[state]bool{state{0, 0}: true} // nodes already visited (we know the minimum distance of those)

	lastmin := 0

	cnt := 0

	for len(fringe) > 0 {
		cur := minimum(fringe, lastmin)

		if cnt%1000 == 0 {
			fmt.Printf("fringe %d (min dist %d)\n", len(fringe), fringe[cur])
		}
		cnt++

		if cur.i == len(M)-1 && cur.j == len(M[cur.i])-1 {
			fmt.Printf("%v %d\n", cur, fringe[cur])
			return
		}

		distcur := fringe[cur]
		lastmin = distcur
		delete(fringe, cur)
		seen[cur] = true

		maybeadd := func(nb state) {
			if seen[nb] {
				return
			}
			if nb.i < 0 || nb.i >= len(M) || nb.j < 0 || nb.j >= len(M[nb.i]) {
				return
			}
			d, ok := fringe[nb]
			if !ok || distcur+M[nb.i][nb.j] < d {
				fringe[nb] = distcur + M[nb.i][nb.j]
			}
		}

		// try to add all possible neighbors
		maybeadd(state{cur.i - 1, cur.j})
		maybeadd(state{cur.i, cur.j - 1})
		maybeadd(state{cur.i, cur.j + 1})
		maybeadd(state{cur.i + 1, cur.j})
	}
}

func copymap(M2 [][]int, offi, offj, offval int) {
	for i := range M {
		for j := range M[i] {
			M2[i+offi][j+offj] = wraparound(M[i][j] + offval)

		}
	}
}

func wraparound(n int) int {
	for n > 9 {
		n -= 9
	}
	return n
}

func main() {
	lines := Input("15.txt", "\n", true)
	pf("len %d\n", len(lines))
	M = make([][]int, len(lines))
	for i := range lines {
		M[i] = make([]int, len(lines[i]))
		for j := range lines[i] {
			M[i][j] = int(lines[i][j] - '0')
		}
	}
	search()

	M2 := make([][]int, len(M)*5)
	for i := range M2 {
		M2[i] = make([]int, len(M[0])*5)
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			copymap(M2, i*len(M), j*len(M[0]), i+j)
		}
	}

	for i := range M2 {
		for j := range M2[i] {
			pf("%d", M2[i][j])
		}
		pf("\n")
	}

	M = M2
	search()
}
