package main

import (
	. "./util"
	"fmt"
	"sort"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func evaluate(line string) (int, int) {
	score := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	stack := []rune{}
	weird := false
	broken := false
	pop := func(ch rune) {
		if len(stack) == 0 {
			weird = true
		}
		if stack[len(stack)-1] == ch {
			stack = stack[:len(stack)-1]
		} else {
			//pf("broken expected %c got %c\n", ch, stack[len(stack)-1])
			broken = true
		}
	}
	_ = weird
	for _, ch := range line {
		switch ch {
		case '(', '[', '{', '<':
			stack = append(stack, ch)
		case ')':
			pop('(')
		case ']':
			pop('[')
		case '}':
			pop('{')
		case '>':
			pop('<')
		}
		if broken {
			return score[ch], 0
		}
	}
	if weird {
		panic("blah")
	}
	m := 0
	part2score := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	for len(stack) > 0 {
		m = m * 5
		m += part2score[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
	}
	return 0, m
}

func main() {
	lines := Input("10.txt", "\n", true)
	part1 := 0
	part2 := []int{}
	for _, line := range lines {
		pf("%s\n", line)
		n, m := evaluate(line)
		part1 += n
		if m > 0 {
			part2 = append(part2, m)
		}
	}
	Sol(part1)
	sort.Ints(part2)
	pf("%v\n", part2)
	Sol(part2[len(part2)/2])
}
