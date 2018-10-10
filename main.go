package main

import (
	"flag"
	"fmt"
	"strconv"
)

func add(a, b int) int

func sub(a, b int) int

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)

	var a, b, r int
	a, _ = strconv.Atoi(args[0])
	b, _ = strconv.Atoi(args[2])

	switch args[1] {
	case "+":
		r = add(a, b)
	case "-":
		r = sub(a, b)
	}

	fmt.Println(r)
}
