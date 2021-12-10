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
	toopen := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	score := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	part2score := map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}
	stack := []rune{}
	pop := func(ch rune) bool {
		if stack[len(stack)-1] == ch {
			stack = stack[:len(stack)-1]
			return true
		}
		return false
	}
	for _, ch := range line {
		switch ch {
		case '(', '[', '{', '<':
			stack = append(stack, ch)
		case ')', ']', '}', '>':
			if !pop(toopen[ch]) {
				return score[ch], 0
			}
		}
	}
	m := 0
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
		n, m := evaluate(line)
		part1 += n
		if m > 0 {
			part2 = append(part2, m)
		}
	}
	Sol(part1)
	sort.Ints(part2)
	Sol(part2[len(part2)/2])
}
