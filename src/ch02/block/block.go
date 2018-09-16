package main

import "fmt"

func main() {
	x := "hello"

	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}

	fmt.Println()

	if a := f(); a == 0 {
		fmt.Println(a)
	} else if b := g(a); a == b {
		fmt.Println(a, b)
	} else {
		fmt.Println(a, b)
	}

	//fmt.Println(a, b)
}

func f() int {
	return 1
}

func g(a int) int {
	return a
}