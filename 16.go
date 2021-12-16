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

const debugParse = false

type Node struct {
	Ver, TypeID int
	Val         int
	Args        []*Node
}

func parse(rest string) (*Node, string) {
	consume := func(out string, n int) int {
		x, err := strconv.ParseInt(rest[:n], 2, 64)
		Must(err)
		if debugParse {
			pf("%s: %s (%d)\n", out, rest[:n], x)
		}
		rest = rest[n:]
		return int(x)
	}

	v := consume("Version", 3)
	part1 += v
	typeID := consume("Type ID", 3)

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
		if debugParse {
			pf("Value: %d\n", val)
		}
		return &Node{Ver: v, TypeID: typeID, Val: val}, rest

	default: // operator
		ltid := consume("Lenght type ID", 1)
		var args []*Node
		if ltid == 0 {
			length := consume("Length", 15)

			contents := rest[:length]
			rest = rest[length:]

			for contents != "" {
				var arg *Node
				arg, contents = parse(contents)
				args = append(args, arg)
			}

		} else {
			numpackets := consume("Num packets", 11)

			for i := 0; i < numpackets; i++ {
				var arg *Node
				arg, rest = parse(rest)
				args = append(args, arg)
			}
		}

		return &Node{Ver: v, TypeID: typeID, Args: args}, rest

		/*
			switch typeID {
			}*/
	}
}

func opcode(typeID int) string {
	switch typeID {
	case 0:
		return "sum"
	case 1:
		return "mul"
	case 2:
		return "min"
	case 3:
		return "max"
	case 5:
		return "gt"
	case 6:
		return "lt"
	case 7:
		return "eq"
	default:
		return fmt.Sprintf("%d", typeID)
	}
}

func printNode(n *Node) {
	if n.TypeID == 4 {
		pf("%d", n.Val)
		return
	}
	pf("(%s", opcode(n.TypeID))
	for i := range n.Args {
		pf(" ")
		printNode(n.Args[i])
	}
	pf(")")
}

func eval(n *Node) int {
	var ret int
	args := n.Args
	switch n.TypeID {
	case 4: // literal value
		return n.Val
	case 0: // sum
		ret = 0
		for i := range args {
			ret += eval(args[i])
		}
	case 1: // product
		ret = 1
		for i := range args {
			ret *= eval(args[i])
		}
	case 2: // minimum
		ret = eval(args[0])
		for i := 1; i < len(args); i++ {
			if a := eval(args[i]); a < ret {
				ret = a
			}
		}
	case 3: // maximum
		ret = eval(args[0])
		for i := 1; i < len(args); i++ {
			if a := eval(args[i]); a > ret {
				ret = a
			}
		}
	case 5: // gt
		if len(args) != 2 {
			panic("wrong args to gt")
		}
		if eval(args[0]) > eval(args[1]) {
			ret = 1
		} else {
			ret = 0
		}
	case 6: // lt
		if len(args) != 2 {
			panic("wrong args to lt")
		}
		if eval(args[0]) < eval(args[1]) {
			ret = 1
		} else {
			ret = 0
		}
	case 7: // eq
		if len(args) != 2 {
			panic("wrong args to eq")
		}
		if eval(args[0]) == eval(args[1]) {
			ret = 1
		} else {
			ret = 0
		}
	default:
		panic("blah")
	}
	return ret
}

func main() {
	lines := Input("16.txt", "\n", true)
	in := lines[0]
	s := ""
	for i := 0; i < len(in); i += 2 {
		n, err := strconv.ParseInt(in[i:i+2], 16, 64)
		Must(err)

		s += fmt.Sprintf("%08b", n)
	}
	node, rest := parse(s)
	pf("%s\n", rest)
	Sol(part1)
	printNode(node)
	pf("\n")
	Sol(eval(node))
}
