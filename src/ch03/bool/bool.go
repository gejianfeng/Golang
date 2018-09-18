package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	{
		s := "hello, world"

		fmt.Println(len(s))
		fmt.Println(s[0], s[7])
		fmt.Println(s[0:5])
		fmt.Println(s[:5])
		fmt.Println(s[7:])
		fmt.Println(s[:])
		fmt.Println("goodby" + s[5:])
	}

	fmt.Println("=====================================")

	{
		s := "left foot"
		t := s
		s += ", right foot"

		fmt.Println(s)
		fmt.Println(t)
	}

	fmt.Println("=====================================")

	{
		fmt.Println(HasPrefix("hello", "hel"))
		fmt.Println(HasSuffix("hello", "llo"))
		fmt.Println(Contains("hello", "ll"))
	}

	fmt.Println("=====================================")

	{
		s := "Hello, 世界"
		fmt.Println(s)
		fmt.Println(utf8.RuneCountInString(s))

		for i := 0; i < len(s); {
			r, size := utf8.DecodeRuneInString(s[i:])
			fmt.Printf("%d\t%q\t%c\n", i, r, r)
			i += size
		}
	}

	fmt.Println("=====================================")

	{
		s := "プログラム"
		fmt.Printf("% x\n", s)

		r := []rune(s)
		fmt.Printf("%x\n", r)

		fmt.Println(string(s))
	}

	fmt.Println("=====================================")

	{
		fmt.Println(string(65))
		fmt.Println(string(0x4eac))
	}
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s) - len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}

	return false
}