package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx)
		fmt.Println(arg)
		fmt.Println("-----------------------")
	}
}