package main

import (
	"fmt"
)

func main()  {

	/*const (
		A = iota
		B = iota
		C = iota
	)*/

	const(
		_ = iota
		KB = 1 << (iota * 10)
		MB
		GB
	)

	/*fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)*/
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}