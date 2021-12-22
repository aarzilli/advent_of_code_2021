package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

type Point struct {
	x, y, z int
}

var M = map[Point]bool{}

func clamp(x int) int {
	if x < -50 {
		return -50
	}
	if x > 50 {
		return 50
	}
	return x
}

func count() int {
	r := 0
	for _, v := range M {
		if v {
			r++
		}
	}
	return r
}

func outside(a, b int) bool {
	if a < -50 && b < -50 {
		return true
	}
	if a > 50 && b > 50 {
		return true
	}
	return false
}

type Cube struct {
	min, max Point
}

func addcube(c Cube, cubes []Cube) []Cube {
	pf("adding %v (%d)\n", c, len(cubes))
	return append(rmcube(c, cubes), c)
}

func rmcube(c Cube, cubes []Cube) []Cube {
	r := []Cube{}
	for _, c2 := range cubes {
		if overlap(c, c2) {
			r = append(r, difference(c2, c)...)
		} else {
			r = append(r, c2)
		}
	}
	return r
}

func overlap(a, b Cube) bool {
	if a.max.x < b.min.x {
		return false
	}
	if a.min.x > b.max.x {
		return false
	}
	if a.max.y < b.min.y {
		return false
	}
	if a.min.y > b.max.y {
		return false
	}
	if a.max.z < b.min.z {
		return false
	}
	if a.min.z > b.max.z {
		return false
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intersection(a, b Cube) Cube {
	return Cube{
		min: Point{
			x: max(a.min.x, b.min.x),
			y: max(a.min.y, b.min.y),
			z: max(a.min.z, b.min.z),
		},
		max: Point{
			x: min(a.max.x, b.max.x),
			y: min(a.max.y, b.max.y),
			z: min(a.max.z, b.max.z),
		},
	}
}

// difference returns the cubes representing the part of c2 not intersecting c
func difference(c1, c2 Cube) []Cube {
	inters := intersection(c1, c2)
	r := []Cube{}

	// bottom cube
	r = append(r, Cube{
		min: Point{
			x: c1.min.x,
			y: c1.min.y,
			z: c1.min.z,
		},
		max: Point{
			x: c1.max.x,
			y: c1.max.y,
			z: inters.min.z - 1,
		},
	})

	// top cube
	r = append(r, Cube{
		min: Point{
			x: c1.min.x,
			y: c1.min.y,
			z: inters.max.z + 1,
		},
		max: Point{
			x: c1.max.x,
			y: c1.max.y,
			z: c1.max.z,
		},
	})

	// left cube (x)
	r = append(r, Cube{
		min: Point{
			x: c1.min.x,
			y: c1.min.y,
			z: inters.min.z,
		},
		max: Point{
			x: inters.min.x - 1,
			y: c1.max.y,
			z: inters.max.z,
		},
	})

	// right cube (x)
	r = append(r, Cube{
		min: Point{
			x: inters.max.x + 1,
			y: c1.min.y,
			z: inters.min.z,
		},
		max: Point{
			x: c1.max.x,
			y: c1.max.y,
			z: inters.max.z,
		},
	})

	// back cube (y)
	r = append(r, Cube{
		min: Point{
			x: inters.min.x,
			y: c1.min.y,
			z: inters.min.z,
		},
		max: Point{
			x: inters.max.x,
			y: inters.min.y - 1,
			z: inters.max.z,
		},
	})

	// front cube (y)
	r = append(r, Cube{
		min: Point{
			x: inters.min.x,
			y: inters.max.y + 1,
			z: inters.min.z,
		},
		max: Point{
			x: inters.max.x,
			y: c1.max.y,
			z: inters.max.z,
		},
	})

	r2 := []Cube{}

	for _, cc := range r {
		if cc.min.x <= cc.max.x && cc.min.y <= cc.max.y && cc.min.z <= cc.max.z {
			r2 = append(r2, cc)
		}
	}

	return r2
}

func countCubes(cubes []Cube) int {
	r := 0
	for _, c := range cubes {
		r += (c.max.x - c.min.x + 1) * (c.max.y - c.min.y + 1) * (c.max.z - c.min.z + 1)
	}
	return r
}

func main() {
	lines := Input("22.txt", "\n", true)
	pf("len %d\n", len(lines))
	cubes := []Cube{}
	for _, line := range lines {
		fields := Spac(line, " ", 2)
		op := fields[0] == "on"
		v := Vatoi(Getnums(line, true, false))

		cube := Cube{Point{v[0], v[2], v[4]}, Point{v[1], v[3], v[5]}}
		if op {
			cubes = addcube(cube, cubes)
		} else {
			cubes = rmcube(cube, cubes)
		}

		if outside(v[0], v[1]) || outside(v[2], v[3]) || outside(v[4], v[5]) {
			continue
		}
		for i := range v {
			v[i] = clamp(v[i])
		}
		for x := v[0]; x <= v[1]; x++ {
			for y := v[2]; y <= v[3]; y++ {
				for z := v[4]; z <= v[5]; z++ {
					M[Point{x, y, z}] = op
				}
			}
		}
	}
	Sol(count())
	Sol(countCubes(cubes))
}
