package main

import (
	"fmt"
	"ch02/tempconv1/tempconv"
)

func main() {
	fmt.Printf("%g\n", tempconv.BoilingC - tempconv.FreezingC)
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF - tempconv.CToF(tempconv.FreezingC))
	//fmt.Printf("%g\n", boilingF - tempconv.FreezingC)

	var c tempconv.Celsius
	var f tempconv.Fahrenheit

	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	//fmt.Println(c == f)
	fmt.Println(c == tempconv.Celsius(f))

	var v = tempconv.FToC(212.0)
	//fmt.Println(v.String())
	fmt.Printf("%v\n", v)
	fmt.Printf("%s\n", v)
	//fmt.Printf(v)
	fmt.Printf("%g\n", v)
	fmt.Println(float64(v))
}