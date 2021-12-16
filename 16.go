package main

import (
	. "./util"
	"fmt"
	"strconv"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

var part1 int

func parse(rest string) (string, int) {
	consume := func(out string, n int) int {
		x, err := strconv.ParseInt(rest[:n], 2, 64)
		Must(err)
		pf("%s: %s (%d)\n", out, rest[:n], x)
		rest = rest[n:]
		return int(x)
	}
	
	v := consume("Version", 3)
	part1 += v
	typeID := consume("Type ID", 3)
	var ret int
	
	switch typeID { 
	case 4: // literal value
		val := 0
		for {
			flag := consume("Flag", 1)
			chunk := consume("Chunk", 4)
			val <<= 4
			val += chunk
			if flag == 0 {
				break
			}
		}
		pf("Value: %d\n", val)
		ret = val
		
	default: // operator
		ltid := consume("Lenght type ID", 1)
		var args []int
		if ltid == 0 {
			length := consume("Length", 15)
			
			contents := rest[:length]
			rest = rest[length:]
			
			for contents != "" {
				var arg int
				contents, arg = parse(contents)
				args = append(args, arg)
			}
			
		} else {
			numpackets := consume("Num packets", 11)
			
			for i := 0; i < numpackets; i++ {
				var arg int
				rest, arg = parse(rest)
				args = append(args, arg)
			}
		}
		
		switch typeID {
		case 0: // sum
			ret = 0
			for i := range args {
				ret += args[i]
			}
		case 1: // product
			ret = 1
			for i := range args {
				ret *= args[i]
			}
		case 2: // minimum
			ret = args[0]
			for i := range args {
				if args[i] < ret {
					ret = args[i]
				}
			}
		case 3: // maximum
			ret = args[0]
			for i := range args {
				if args[i] > ret {
					ret = args[i]
				}
			}
		case 5: // gt
			if len(args) != 2 {
				panic("wrong args to gt")
			}
			if args[0] > args[1] {
				ret = 1
			} else {
				ret = 0
			}
		case 6: // lt
			if len(args) != 2 {
				panic("wrong args to lt")
			}
			if args[0] < args[1] {
				ret = 1
			} else {
				ret = 0
			}
		case 7: // eq
			if len(args) != 2 {
				panic("wrong args to eq")
			}
			if args[0] == args[1] {
				ret = 1
			} else {
				ret = 0
			}
		default:
			panic("blah")
		}
	}
	
	return rest, ret
}

func main() {
	lines := Input("16.txt", "\n", true)
	pf("len %d\n", len(lines))
	in := lines[0]
	s := ""
	for i := 0; i < len(in); i += 2 {
		n, err := strconv.ParseInt(in[i:i+2], 16, 64)
		Must(err)
		
		s += fmt.Sprintf("%08b", n)
	}
	pf("%s (%d)\n", s, len(s))
	rest := s
	var ret int
	rest, ret = parse(rest)
	pf("%s\n", rest)
	Sol(part1)
	Sol(ret)
	
}
