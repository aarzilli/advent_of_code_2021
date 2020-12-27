package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// returns x without the last character
func Nolast(x string) string {
	return x[:len(x)-1]
}

// splits a string, trims spaces on every element
func Spac(in, sep string, n int) []string {
	v := strings.SplitN(in, sep, n)
	for i := range v {
		v[i] = strings.TrimSpace(v[i])
	}
	return v
}

// convert string to integer
func Atoi(in string) int {
	n, err := strconv.Atoi(in)
	Must(err)
	return n
}

// convert vector of strings to integer
func Vatoi(in []string) []int {
	r := make([]int, len(in))
	for i := range in {
		var err error
		r[i], err = strconv.Atoi(in[i])
		Must(err)
	}
	return r
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Exit(n int) {
	os.Exit(n)
}

func Pf(fmtstr string, args ...interface{}) {
	fmt.Printf(fmtstr, args...)
}

func Getints(in string, hasneg bool) []int {
	v := Getnums(in, hasneg, false)
	return Vatoi(v)
}

func Getnums(in string, hasneg, hasdot bool) []string {
	r := []string{}
	start := -1

	flush := func(end int) {
		if start < 0 {
			return
		}
		hasdigit := false
		for i := start; i < end; i++ {
			if in[i] >= '0' && in[i] <= '9' {
				hasdigit = true
				break
			}
		}
		if hasdigit {
			r = append(r, in[start:end])
		}
		start = -1
	}

	for i, ch := range in {
		isnumch := false

		switch {
		case hasneg && (ch == '-'):
			isnumch = true
		case hasdot && (ch == '.'):
			isnumch = true
		case ch >= '0' && ch <= '9':
			isnumch = true
		}

		if start >= 0 {
			if !isnumch {
				flush(i)
			}
		} else {
			if isnumch {
				start = i
			}
		}
	}
	flush(len(in))
	return r
}

// removes empty string elements, modifies v
func Noempty(v []string) []string {
	r := v[:0]
	for _, s := range v {
		if s != "" {
			r = append(r, s)
		}
	}
	return r
}

func Input(path string, sep string, noempty bool) []string {
	buf, err := ioutil.ReadFile(path)
	Must(err)
	lines := Spac(string(buf), sep, -1)
	if noempty {
		lines = Noempty(lines)
	}
	return lines
}

var part int = 1

func Sol(v ...interface{}) {
	fmt.Printf("PART %d: ", part)
	fmt.Println(v...)
	fmt.Printf("copied to clipboard\n")
	cmd := exec.Command("xclip", "-in", "-selection", "-primary")
	cmd.Stdin = bytes.NewReader([]byte(fmt.Sprintf("%v", v[len(v)-1])))
	cmd.Run()
	part++
}
