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

type Scanner struct {
	id   int
	zero Point
	pts  []Point
}

func transpose(s *Scanner, i int, z Point) {
	// transposes s to where point s.pts[i] equals z

	dx := z.x - s.pts[i].x
	dy := z.y - s.pts[i].y
	dz := z.z - s.pts[i].z

	for i := range s.pts {
		rp := s.pts[i]
		rp.x = rp.x + dx
		rp.y = rp.y + dy
		rp.z = rp.z + dz
		s.pts[i] = rp
	}

	s.zero.x += dx
	s.zero.y += dy
	s.zero.z += dz
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

func rotated(s *Scanner, pitch, roll, yaw int) *Scanner {
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

func overlapmore(s0, s1 *Scanner, n int) bool {
	cnt := 0

	Sort(s1.pts, pointLess)

	i, j := 0, 0

	for i < len(s0.pts) && j < len(s1.pts) {
		p0, p1 := s0.pts[i], s1.pts[j]

		switch {
		case p0 == p1:
			cnt++
			if cnt >= n {
				return true
			}
			i++
			j++
		case pointLess(p0, p1):
			i++
		default:
			j++
		}
	}

	return false
}

var allrots = [][3]int{
	{0, 0, 0},
	{0, 0, 180},
	{0, 0, 270},
	{0, 0, 90},
	{0, 180, 0},
	{0, 180, 180},
	{0, 180, 270},
	{0, 180, 90},
	{0, 270, 0},
	{0, 270, 180},
	{0, 270, 270},
	{0, 270, 90},
	{0, 90, 0},
	{0, 90, 180},
	{0, 90, 270},
	{0, 90, 90},
	{270, 0, 0},
	{270, 0, 180},
	{270, 0, 270},
	{270, 0, 90},
	{90, 0, 0},
	{90, 0, 180},
	{90, 0, 270},
	{90, 0, 90},
}

func align(acc *Scanner, s *Scanner) (bool, Point) {
	for _, rot := range allrots {
		rotS := rotated(s, rot[0], rot[1], rot[2])
		for i := 0; i < len(acc.pts); i++ { // acc coord
			for j := 0; j < len(s.pts); j++ { // s coord
				transpose(rotS, j, acc.pts[i])
				if overlapmore(acc, rotS, 12) {
					pf("FOUND %v!!!!\n", rotS.zero)
					acc.pts = append(acc.pts, rotS.pts...)
					simplify(acc)
					return true, rotS.zero
				}
			}
		}
	}
	return false, Point{}
}

func simplify(acc *Scanner) {
	Sort(acc.pts, pointLess)
	newpts := []Point{acc.pts[0]}
	for _, pt := range acc.pts {
		if pt != newpts[len(newpts)-1] {
			newpts = append(newpts, pt)
		}
	}
	acc.pts = newpts
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
	for i := range scanners {
		scanners[i].id = i
	}

	aligned := make([]bool, len(scanners))
	aligned[0] = true
	acc := rotated(scanners[0], 0, 0, 0)
	Sort(acc.pts, pointLess)

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
	Expect(315)
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
	Expect(13192)
	Sol(maxd)

}
