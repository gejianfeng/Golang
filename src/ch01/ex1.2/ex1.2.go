package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, value := range os.Args[1:] {
		fmt.Printf("Index: %d | Value: %s\n", idx, value)
	}
}