package main

import (
	. "./util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func main() {
	lines := Input("01.txt", "\n", true)
	pf("len %d\n", len(lines))
	first := true
	cnt := 0
	prev := 0
	for _, x := range lines {
		n := Atoi(x)
		
		if first {
			first = false
		} else {
			if n > prev {
				cnt++
			}
		}
		prev = n
	}
	Sol(cnt)
	
	v := Vatoi(lines)
	first = true
	cnt2 := 0
	for i := 2; i < len(v); i++ {
		n := v[i] + v[i-1] + v[i-2]
		if first {
			first = false
		} else {
			if n > prev {
				cnt2++
			}
		}
		prev = n
	}
	Sol(cnt2)
}
