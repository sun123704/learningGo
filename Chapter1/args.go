package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)
	*/
	/*
		s, sep := "", ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
		fmt.Println(s)
	*/
	// 键盘输入
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
