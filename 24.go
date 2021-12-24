package main

import (
	. "./util"
	"fmt"
	"strconv"
)

type Instr struct {
	opcode string
	args   []Arg
}

func parseInstr(line string) Instr {
	fields := Spac(line, " ", -1)
	opcode := fields[0]
	args := make([]Arg, len(fields)-1)
	for i, field := range fields[1:] {
		if field[len(field)-1] == ',' {
			field = field[:len(field)-1]
		}
		n, err := strconv.Atoi(field)
		if err == nil {
			args[i] = Arg{val: n}
		} else {
			args[i] = Arg{reg: field}
		}
	}
	return Instr{opcode, args}
}

func (instr Instr) argMustBeReg(argnum int) {
	if instr.args[argnum].reg == "" {
		panic("arg is not register")
	}
}

type Arg struct {
	reg string
	val int
}

func (a Arg) value(regs map[string]int) int {
	if a.reg == "" {
		return a.val
	}
	return regs[a.reg]
}

var text []Instr
var input = []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}

func run(z int) int {
	pc := 0
	regs := map[string]int{}
	regs["z"] = z

	for {
		if pc >= len(text) {
			return regs["z"]
		}
		instr := text[pc]
		//pf("%v %v\n", regs, instr)
		switch instr.opcode {
		case "inp":
			instr.argMustBeReg(0)
			regs[instr.args[0].reg] = input[0]
			input = input[1:]
			pc++
		case "add":
			instr.argMustBeReg(0)
			regs[instr.args[0].reg] = instr.args[0].value(regs) + instr.args[1].value(regs)
			pc++
		case "mul":
			instr.argMustBeReg(0)
			regs[instr.args[0].reg] = instr.args[0].value(regs) * instr.args[1].value(regs)
			pc++
		case "div":
			if instr.args[1].value(regs) == 0 {
				pf("DIVIDE BY 0 at 24.txt:%d\n", pc)
			}
			instr.argMustBeReg(0)
			regs[instr.args[0].reg] = instr.args[0].value(regs) / instr.args[1].value(regs)
			pc++
		case "mod":
			instr.argMustBeReg(0)
			if instr.args[0].value(regs) < 0 || instr.args[1].value(regs) <= 0 {
				pf("BAD MODULO at 24.txt:%d\n", pc)
			}
			regs[instr.args[0].reg] = instr.args[0].value(regs) % instr.args[1].value(regs)
			pc++
		case "eql":
			instr.argMustBeReg(0)
			if instr.args[0].value(regs) == instr.args[1].value(regs) {
				regs[instr.args[0].reg] = 1
			} else {
				regs[instr.args[0].reg] = 0
			}
			pc++
		default:
			panic("blah")
		}
	}
}

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func simchunk(z, in int, var0, var1, var2 int) int {
	x := 0
	if (z%26)+var1 != in {
		x = 1
	}
	z = z / var0
	if x != 0 {
		z = z * 26
		z += in + var2
	}
	return z
}

var chunks = [][3]int{
	{1, 11, 3},
	{1, 14, 7},
	{1, 13, 1},
	{26, -4, 6},
	{1, 11, 14},
	{1, 10, 7},
	{26, -4, 9},
	{26, -12, 9},
	{1, 10, 6},
	{26, -11, 4},
	{1, 12, 0},
	{26, -1, 7},
	{26, 0, 12},
	{26, -11, 1},
}

var maxx = 0
var minx = 0
var first = true

func enum(pass []int, i int, admout []map[int]bool, curz int) {
	if i >= len(pass) {
		x := 0
		for i := range pass {
			x *= 10
			x += pass[i]
			pf("%d", pass[i])
		}
		pf("\n")
		pf("%d\n", x)
		if x > maxx {
			maxx = x
		}
		if first || x < minx {
			minx = x
			first = false
		}
		return
	}
	chunk := chunks[i]
	for in := 9; in > 0; in-- {
		outz := simchunk(curz, in, chunk[0], chunk[1], chunk[2])
		if admout[i][outz] {
			pf("stage %d valid character %d\n", i, in)
			pass[i] = in
			enum(pass, i+1, admout, outz)
		}
	}
}

func main() {
	if len(input) != 14 {
		panic("blah")
	}
	lines := Input("24.simple.txt", "\n", true)
	pf("len %d\n", len(lines))
	for _, line := range lines {
		text = append(text, parseInstr(line))
	}
	/*	for i := 0; i < len(text); i += 18 {
			var0 := text[i+4].args[1].val
			var1 := text[i+5].args[1].val
			var2 := text[i+15].args[1].val

			pf("{ %d, %d, %d }\n", var0, var1, var2)


	// 		pf("INPUT %d (%d %d %d)\n", i/18, var0, var1, var2)
	// 		pf("x = 1 if (z %% 26) %+d != INPUTCHAR else x = 0\n", var1)
	// 		pf("z = (z / %d)\n", var0)
	// 		pf("if x != 0 {\n")
	// 		pf("\tz = z * 26\n")
	// 		pf("\tz = z + INPUTCHAR + 1\n")
	// 		pf("}\n\n")


		}*/

	//input[0] = 9
	//run()
	/*for in := 0; in <= 9; in++ {
		for z := 0; z <= 25; z++ {
			input = []int{ in }
			out := run(z)
			pf("in=%d z=%d = %d %d\n", in, z, out, simchunk(z, in, 26, 0, 12))
			if out != simchunk(z, in, 26, 0, 12) {
				panic("blah")
			}
		}
	}*/

	/*for z := 0; z <= 50; z++ {
		zz := simchunk(z, 8, 26, 0, 12)
		zz = simchunk(zz, 9, 26, -11, 1)
		pf("%d -> %d\n", z, zz)
	}*/
	/*for z := 0; z <= 50; z++ {
		zz := simchunk(z, 9, 26, -11, 1)
		pf("%d -> %d\n", z%26, zz)
	}*/

	/*for in := 0; in <= 9; in++ {
		chunk := chunks[4]
		for z := 0; z <= 200; z++ {
			input = []int{ in }
			out := run(z)
			pf("z=%d in=%d -> %d %d\n", z, in, simchunk(z, in, chunk[0], chunk[1], chunk[2]), out)
		}
	}
	*/

	admissibleOut := map[int]bool{0: true}
	admout := make([]map[int]bool, len(chunks))

	for i := len(chunks) - 1; i >= 0; i-- {
		admout[i] = admissibleOut
		chunk := chunks[i]
		pf("admissible z outputs at stage %d: %v\n", i, admissibleOut)
		admissibleIn := map[int]bool{}
		for inz := 0; inz <= 2000000; inz++ {
			for in := 1; in <= 9; in++ {
				outz := simchunk(inz, in, chunk[0], chunk[1], chunk[2])
				if admissibleOut[outz] {
					admissibleIn[inz] = true
				}
			}
		}
		admissibleOut = admissibleIn
	}

	pf("admissible input values of z: %v\n", admissibleOut)

	pass := make([]int, len(chunks))
	enum(pass, 0, admout, 0)
	Sol(maxx)
	Sol(minx)

	/*

		curz := 0
		for i := 0; i < len(chunks); i++ {
			chunk := chunks[i]
			for in := 0; in <= 9; in++ {
				outz := simchunk(curz, in, chunk[0], chunk[1], chunk[2])
				pf("%d -> %d\n", in, outz)
				if admout[i][outz] {
					pf("stage %d valid character %d\n", i, in)
					if in > pass[i] {
						pass[i] = in
					}
				}
			}

			outz := simchunk(curz, pass[i], chunk[0], chunk[1], chunk[2])
			curz = outz
		}
		for i := range pass {
			pf("%d", pass[i])
		}
		pf("\n")*/
}
