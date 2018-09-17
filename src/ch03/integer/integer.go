package main

import (
	"fmt"
)

func main() {
	{
		var u uint8 = 255
		fmt.Println(u, u+1, u*u)

		var i int8 = 127
		fmt.Println(i, i+1, i*i)
	}

	fmt.Println("========================================")

	{
		var x uint8 = 1 << 1 | 1 << 5
		var y uint8 = 1 << 1 | 1 << 2

		fmt.Printf("%08b\n", x)			//00100010
		fmt.Printf("%08b\n", y)			//00000110

		fmt.Printf("%08b\n", x & y)		//00000010
		fmt.Printf("%08b\n", x | y)		//00100110
		fmt.Printf("%08b\n", x ^ y)		//00100100
		fmt.Printf("%08b\n", x & ^y)	//00100000

		for i := uint(0); i < 8; i++ {
			if x & (1 << i) != 0 {
				fmt.Println(i)
			}
		}

		fmt.Printf("%08b\n", x << 1)
		fmt.Printf("%08b\n", x >> 1)
	}

	fmt.Println("========================================")

	{
		medals := []string{"gold", "silver", "bronze"}

		for i := len(medals) - 1; i >= 0; i-- {
			fmt.Println(medals[i])
		}
	}

	fmt.Println("========================================")

	{
		var apples int32 = 1
		var oranges int16 = 2

		//var compote int = apples + oranges
	
		var compote = int(apples) + int(oranges)
		fmt.Println(compote)
	}

	fmt.Println("========================================")

	{
		f := 3.141
		i := int(f)

		fmt.Println(f, i)

		f = 1.99
		fmt.Println(int(f))

		f = 1e100
		i = int(f)
		fmt.Println(i)
	}

	fmt.Println("========================================")

	{
		o := 0666
		fmt.Printf("%d %[1]o %#[1]o\n", o)

		x := int64(0xdeadbeef)
		fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	}

	fmt.Println("========================================")


	{
		ascii := 'a'
		unicode := '国'
		newline := '\n'

		fmt.Printf("%d %[1]c %[1]q\n", ascii)
		fmt.Printf("%d %[1]c %[1]q\n", unicode)
		fmt.Printf("%d %[1]q\n", newline)
	}
}