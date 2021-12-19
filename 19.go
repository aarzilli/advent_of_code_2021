package main

import (
	. "./util"
	"fmt"
	"sort"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

type Point struct {
	x, y, z int
}

type Scanner struct {
	zero Point
	pts  []Point
}

func transposed(s *Scanner, i int, z Point) *Scanner {
	// returns transposed version of s where point s.pts[i] equals z

	dx := z.x - s.pts[i].x
	dy := z.y - s.pts[i].y
	dz := z.z - s.pts[i].z

	rs := &Scanner{}

	for i := range s.pts {
		rp := s.pts[i]
		rp.x = rp.x + dx
		rp.y = rp.y + dy
		rp.z = rp.z + dz
		rs.pts = append(rs.pts, rp)
	}

	rs.zero.x = dx
	rs.zero.y = dy
	rs.zero.z = dz

	return rs
}

func cos(a int) int {
	switch a {
	case 0:
		return 1
	case 90:
		return 0
	case 180:
		return -1
	case 270:
		return 0
	default:
		panic("blah")
	}
}

func sin(a int) int {
	switch a {
	case 0:
		return 0
	case 90:
		return 1
	case 180:
		return 0
	case 270:
		return -1
	default:
		panic("blah")
	}
}

func rotated(s *Scanner, roll, pitch, yaw int) *Scanner {
	cosa := cos(yaw)
	sina := sin(yaw)

	cosb := cos(pitch)
	sinb := sin(pitch)

	cosc := cos(roll)
	sinc := sin(roll)

	var Axx = cosa * cosb
	var Axy = cosa*sinb*sinc - sina*cosc
	var Axz = cosa*sinb*cosc + sina*sinc

	var Ayx = sina * cosb
	var Ayy = sina*sinb*sinc + cosa*cosc
	var Ayz = sina*sinb*cosc - cosa*sinc

	var Azx = -sinb
	var Azy = cosb * sinc
	var Azz = cosb * cosc

	rs := &Scanner{}
	for i := range s.pts {
		px, py, pz := s.pts[i].x, s.pts[i].y, s.pts[i].z
		rs.pts = append(rs.pts, Point{
			x: Axx*px + Axy*py + Axz*pz,
			y: Ayx*px + Ayy*py + Ayz*pz,
			z: Azx*px + Azy*py + Azz*pz,
		})
	}
	return rs
}

func countoverlap(s0, s1 *Scanner) int {
	cnt := 0
	for i := range s0.pts {
		for j := range s1.pts {
			if s0.pts[i] == s1.pts[j] {
				cnt++
			}
		}
	}
	return cnt
}

func clone(s *Scanner) *Scanner {
	return transposed(s, 0, s.pts[0])
}

func align(acc *Scanner, s *Scanner) (bool, Point) {
	for i := 0; i < len(acc.pts); i++ { // acc coord
		for j := 0; j < len(s.pts); j++ { // s coord
			for pitch := 0; pitch < 360; pitch += 90 {
				for roll := 0; roll < 360; roll += 90 {
					for yaw := 0; yaw < 360; yaw += 90 {
						a := rotated(s, pitch, roll, yaw)
						b := transposed(a, j, acc.pts[i])
						d := countoverlap(acc, b)
						if d >= 12 {
							pf("FOUND %v!!!!\n", b.zero)
							acc.pts = append(acc.pts, b.pts...)
							simplify(acc)
							return true, b.zero
						}
					}
				}
			}
		}
	}
	return false, Point{}
}

func simplify(acc *Scanner) {
	sort.Slice(acc.pts, byPoint(acc.pts))
	newpts := []Point{acc.pts[0]}
	for _, pt := range acc.pts {
		if pt != newpts[len(newpts)-1] {
			newpts = append(newpts, pt)
		}
	}
	acc.pts = newpts
}

func byPoint(pts []Point) func(i, j int) bool {
	return func(i, j int) bool {
		pi, pj := pts[i], pts[j]
		return pointLess(pi, pj)
	}
}

func pointLess(pa, pb Point) bool {
	if pa.x == pb.x {
		if pa.y == pb.y {
			return pa.z < pb.z
		}
		return pa.y < pb.y
	}
	return pa.x < pb.x
}

func dist(pa, pb Point) int {
	return Abs(pa.x-pb.x) + Abs(pa.y-pb.y) + Abs(pa.z-pb.z)
}

func main() {
	groups := Input("19.txt", "\n\n", true)
	pf("len %d\n", len(groups))

	var scanners = []*Scanner{}

	for _, group := range groups {
		v := Spac(group, "\n", -1)
		s := &Scanner{}
		for _, ptstr := range v[1:] {
			f := Vatoi(Spac(ptstr, ",", -1))
			s.pts = append(s.pts, Point{f[0], f[1], f[2]})
		}
		scanners = append(scanners, s)
	}

	aligned := make([]bool, len(scanners))
	aligned[0] = true
	acc := clone(scanners[0])

	tsp := make([]Point, len(scanners))

	for {
		for i := range scanners {
			if !aligned[i] {
				pf("aligning %d\n", i)
				ok, zero := align(acc, scanners[i])
				if ok {
					tsp[i] = zero
					aligned[i] = true
				}
			}
		}
		allaligned := true
		for i := range scanners {
			if !aligned[i] {
				allaligned = false
				break
			}
		}
		pf("all aligned: %v\n", allaligned)
		if allaligned {
			break
		}
	}
	Sol(len(acc.pts))

	maxd := 0
	for i := range scanners {
		for j := range scanners {
			if i == j {
				continue
			}
			d := dist(tsp[i], tsp[j])
			if d > maxd {
				maxd = d
			}
		}
	}
	Sol(maxd)
	//pf("%v\n", transposed(scanners[1], 0, Point{ -618,-824,-621 }))

}

// {1 1 1} {2 2 2} {3 3 3} {3 1 2} {-6 -4 -5} {0 7 -8}
