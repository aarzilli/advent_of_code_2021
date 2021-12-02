package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func main() {
	lines := Input("02.txt", "\n", true)
	pf("len %d\n", len(lines))
	i, j := 0, 0
	for _, line := range lines {
		v := Spac(line, " ", 2)
		n := Atoi(v[1])
		switch v[0] {
		case "forward":
			j += n
		case "up":
			i -= n
		case "down":
			i += n
		default:
			panic("blah")
		}
	}
	Sol(i, j, i*j)

	i, j = 0, 0
	aim := 0
	for _, line := range lines {
		v := Spac(line, " ", 2)
		n := Atoi(v[1])
		switch v[0] {
		case "forward":
			j += n
			i += n * aim
		case "up":
			aim -= n
		case "down":
			aim += n
		default:
			panic("blah")
		}
	}
	Sol(i, j, i*j)

}
